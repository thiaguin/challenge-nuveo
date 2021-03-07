package controllers

import (
	"errors"
	fileService "microservice/services/file"
	messageService "microservice/services/message"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	fileService := fileService.FileRepositoryMock{}
	messageService := messageService.MessageRepositoryMock{}
	fileService.On("Exist").Return(true, nil)
	fileService.On("Write").Return(nil)
	messageService.On("Dequeue").Return([]byte{}, nil)
	controller := NewFileController(&messageService, &fileService)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Write)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK, "Should return status code ok")

	fileService.AssertCalled(t, "Exist")
	fileService.AssertCalled(t, "Write")
	messageService.AssertCalled(t, "Dequeue")
}

func TestWriteFilePathError(t *testing.T) {
	fileService := fileService.FileRepositoryMock{}
	messageService := messageService.MessageRepositoryMock{}
	fileService.On("Exist").Return(true, errors.New(""))
	fileService.On("Write").Return(nil)
	messageService.On("Dequeue").Return([]byte{}, nil)
	controller := NewFileController(&messageService, &fileService)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Write)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest, "Should return status code ok")

	fileService.AssertCalled(t, "Exist")
	fileService.AssertNotCalled(t, "Write")
	messageService.AssertNotCalled(t, "Dequeue")
}

func TestWriteFileMessageError(t *testing.T) {
	fileService := fileService.FileRepositoryMock{}
	messageService := messageService.MessageRepositoryMock{}
	fileService.On("Exist").Return(true, nil)
	fileService.On("Write").Return(nil)
	messageService.On("Dequeue").Return([]byte{}, errors.New(""))
	controller := NewFileController(&messageService, &fileService)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Write)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError, "Should return status code ok")

	fileService.AssertCalled(t, "Exist")
	messageService.AssertCalled(t, "Dequeue")
	fileService.AssertNotCalled(t, "Write")
}

func TestWriteFileNoMessage(t *testing.T) {
	fileService := fileService.FileRepositoryMock{}
	messageService := messageService.MessageRepositoryMock{}
	fileService.On("Exist").Return(true, nil)
	fileService.On("Write").Return(nil)
	messageService.On("Dequeue").Return(nil, nil)
	controller := NewFileController(&messageService, &fileService)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Write)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusNotFound, "Should return status code ok")

	fileService.AssertCalled(t, "Exist")
	messageService.AssertCalled(t, "Dequeue")
	fileService.AssertNotCalled(t, "Write")
}

func TestWriteFileOnWriteError(t *testing.T) {
	fileService := fileService.FileRepositoryMock{}
	messageService := messageService.MessageRepositoryMock{}
	fileService.On("Exist").Return(true, nil)
	fileService.On("Write").Return(errors.New(""))
	messageService.On("Dequeue").Return([]byte{}, nil)
	controller := NewFileController(&messageService, &fileService)

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Write)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError, "Should return status code ok")

	fileService.AssertCalled(t, "Exist")
	messageService.AssertCalled(t, "Dequeue")
	fileService.AssertCalled(t, "Write")
}
