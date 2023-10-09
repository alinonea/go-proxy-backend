package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/alinonea/main/controller"
	domain "github.com/alinonea/main/domain"
	repo "github.com/alinonea/main/repo"
)

var customTransport = http.DefaultTransport
var env = domain.NewEnvironmentVariables()

func main() {

	db, err := repo.CreateConnection(*env)
	if err != nil {
		log.Fatalf("Error at connecting to db: %v", err)
	}

	err = db.CreateDB()
	if err != nil {
		log.Fatalf("Error at creating the db: %v", err)
	}

	requestsRepo := repo.NewRequestRepository(db.Db)
	handler := controller.NewHandler(requestsRepo)

	// Create a new HTTP server with the handleRequest function as the handler
	// and having a middleware attached to it for rate limiting per ip
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", env.PORT),
		Handler: limitMiddleware(handler.HandleRequest),
	}

	// Start the server and log any errors
	log.Println(fmt.Sprintf("Starting proxy server on : %v", env.PORT))
	log.Printf("test")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server: ", err)
	}
}

func limitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	var limiter = NewIPRateLimiter(rate.Every(time.Duration(env.NUMBER_OF_SECONDS)*time.Second), env.NUMBER_OF_REQUESTS)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
