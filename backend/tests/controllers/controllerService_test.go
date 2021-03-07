package tests

import (
	clientController "backend/controllers/client"
	"backend/models"
	clientService "backend/tests/services/client"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	service := clientService.ClientServiceMock{}
	service.On("GetById").Return(&models.Client{}, nil)
	controller := clientController.NewClientController(&service)

	req, err := http.NewRequest("GET", "/client/id", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetById)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "GetById")

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK, "Should return status code ok")
}

func TestGetByIdError(t *testing.T) {
	service := clientService.ClientServiceMock{}
	notFoundErrorStatus := 404
	service.On("GetById").Return(nil, notFoundErrorStatus)
	controller := clientController.NewClientController(&service)

	req, err := http.NewRequest("GET", "/client/id", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetById)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "GetById")

	assert.Equal(t, rr.Result().StatusCode, notFoundErrorStatus, "Should return service error status code")
}

func TestGetAll(t *testing.T) {
	service := clientService.ClientServiceMock{}
	service.On("GetAll").Return([]models.Client{}, nil)
	controller := clientController.NewClientController(&service)

	req, err := http.NewRequest("GET", "/client", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAll)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "GetAll")

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK, "Should return status code ok")
}

func TestGetAllError(t *testing.T) {
	service := clientService.ClientServiceMock{}
	internalErrorStatus := 500
	service.On("GetAll").Return(nil, internalErrorStatus)
	controller := clientController.NewClientController(&service)

	req, err := http.NewRequest("GET", "/client", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAll)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "GetAll")

	assert.Equal(t, rr.Result().StatusCode, internalErrorStatus, "Should return service error status code")
}

func TestCreate(t *testing.T) {
	service := clientService.ClientServiceMock{}
	service.On("Create").Return(&models.Client{}, nil)
	controller := clientController.NewClientController(&service)
	body := []byte(`{"name": "Nome", "address": "address"}`)

	req, err := http.NewRequest("POST", "/client", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Create)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "Create")

	assert.Equal(t, rr.Result().StatusCode, http.StatusCreated, "Should return status code ok")
}

func TestCreateDecodeBodyError(t *testing.T) {
	service := clientService.ClientServiceMock{}
	badRequestStatus := 400
	controller := clientController.NewClientController(&service)
	body := []byte(``)

	req, err := http.NewRequest("POST", "/client", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Create)

	handler.ServeHTTP(rr, req)

	service.AssertNotCalled(t, "Create")

	assert.Equal(t, rr.Result().StatusCode, badRequestStatus, "Should return service error status code")
}

func TestCreateErrorFromService(t *testing.T) {
	service := clientService.ClientServiceMock{}
	badRequestStatus := 400
	service.On("Create").Return(nil, badRequestStatus)
	controller := clientController.NewClientController(&service)
	body := []byte(`{"name": "Nome", "address": "address"}`)

	req, err := http.NewRequest("POST", "/client", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Create)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "Create")

	assert.Equal(t, rr.Result().StatusCode, badRequestStatus, "Should return service error status code")
}

func TestUpdate(t *testing.T) {
	service := clientService.ClientServiceMock{}
	service.On("Update").Return(&models.Client{}, nil)
	controller := clientController.NewClientController(&service)
	body := []byte(`{"name": "Upddate Nome", "address": "Update Address"}`)

	req, err := http.NewRequest("UPDATE", "/client/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Update)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "Update")

	assert.Equal(t, rr.Result().StatusCode, http.StatusNoContent, "Should return status code ok")
}

func TestUpdateDecodeBodyError(t *testing.T) {
	service := clientService.ClientServiceMock{}
	controller := clientController.NewClientController(&service)
	body := []byte(``)

	req, err := http.NewRequest("UPDATE", "/client/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Update)

	handler.ServeHTTP(rr, req)

	service.AssertNotCalled(t, "Update")

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest, "Should return status code ok")
}

func TestUpdateServiceError(t *testing.T) {
	service := clientService.ClientServiceMock{}
	serviceStatusError := 500
	service.On("Update").Return(nil, serviceStatusError)
	controller := clientController.NewClientController(&service)
	body := []byte(`{"name": "Upddate Nome", "address": "Update Address"}`)

	req, err := http.NewRequest("UPDATE", "/client/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Update)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "Update")

	assert.Equal(t, rr.Result().StatusCode, serviceStatusError, "Should return status code ok")
}

func TestDelete(t *testing.T) {
	service := clientService.ClientServiceMock{}
	service.On("Delete").Return(nil)
	controller := clientController.NewClientController(&service)

	req, err := http.NewRequest("DELETE", "/client/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Delete)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "Delete")

	assert.Equal(t, rr.Result().StatusCode, http.StatusNoContent, "Should return status code ok")
}

func TestDeleteError(t *testing.T) {
	service := clientService.ClientServiceMock{}
	serviceStatusError := 500
	service.On("Delete").Return(serviceStatusError)
	controller := clientController.NewClientController(&service)

	req, err := http.NewRequest("DELETE", "/client/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Delete)

	handler.ServeHTTP(rr, req)

	service.AssertCalled(t, "Delete")

	assert.Equal(t, rr.Result().StatusCode, serviceStatusError, "Should return status code ok")
}
