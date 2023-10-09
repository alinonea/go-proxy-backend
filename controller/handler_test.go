package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alinonea/main/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleRequest(t *testing.T) {
	requestsRepo := &mocks.RequestRepositoryInterface{}
	handler := NewHandler(requestsRepo)

	t.Run("should return internal server error when saving the request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		res := httptest.NewRecorder()

		requestsRepo.On("SaveRequest", mock.Anything).Return(errors.New("db err")).Once()

		handler.HandleRequest(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		res := httptest.NewRecorder()

		requestsRepo.On("SaveRequest", mock.Anything).Return(nil).Once()

		handler.HandleRequest(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
	})

}
