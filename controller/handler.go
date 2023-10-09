package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alinonea/main/domain"
	repo "github.com/alinonea/main/repo"
)

type Handler struct {
	repo repo.RequestRepositoryInterface
}

func NewHandler(repo repo.RequestRepositoryInterface) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (handler *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Create a new HTTP request with the same method, URL, and body as the original request
	targetURL := r.URL
	baseUrl := "https://jsonplaceholder.typicode.com"

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	fmt.Println(targetURL.String())

	resp, err := c.Get(baseUrl + targetURL.String())
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()

	// // Copy the headers from the proxy response to the original response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	byteBody, err := io.ReadAll(resp.Body)

	var structBody map[string]interface{} // declaring a map for key names as string and values as interface

	err = json.Unmarshal(byteBody, &structBody)
	if err != nil {
		fmt.Printf("Unmarshal error: %v", err)
	}

	structBody["foo"] = "bar"

	responseBody, _ := json.Marshal(structBody)
	requestToSave := domain.Request{
		RequestBody: (*json.RawMessage)(&responseBody),
		RemoteAddr:  r.RemoteAddr,
	}
	err = handler.repo.SaveRequest(requestToSave)
	if err != nil {
		errorStruct := make(map[string]string)
		errorStruct["message"] = err.Error()
		errorBody, _ := json.Marshal(errorStruct)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorBody)
		return
	}
	// // Set the status code of the original response to the status code of the proxy response
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}
