package services

import (
	customError "backend/error"
	"backend/models"
	dto "backend/models/dto"
	clientRepository "backend/repositories/client"
	messageService "backend/services/message"
	"errors"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(&models.Client{}, nil)
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	client, clientErr := service.GetById("id")

	repository.AssertCalled(t, "GetById")

	assert.Nil(t, clientErr)
	assert.NotNil(t, client)
	assert.Equal(t, reflect.TypeOf(client), reflect.TypeOf(&models.Client{}), "Should be equal")
}

func TestGetByIdNotFound(t *testing.T) {
	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(nil, errors.New(""))
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	client, clientErr := service.GetById("id")

	repository.AssertCalled(t, "GetById")

	assert.NotNil(t, clientErr, "Should client error not be nil")
	assert.Nil(t, client, "Should client be nil")
	assert.Equal(t, reflect.TypeOf(clientErr), reflect.TypeOf(&customError.HTTPError{}), "Should be return http custom error")
	assert.Equal(t, clientErr.Status, 404, "Should return 404 error status code")
}

func TestGetAll(t *testing.T) {
	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetAll").Return([]models.Client{}, nil)
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	clients, clientsErr := service.GetAll()

	repository.AssertCalled(t, "GetAll")

	assert.NotNil(t, clients, "Should clients be not nil")
	assert.Nil(t, clientsErr, "Shoul client error be nil")
	assert.Equal(t, reflect.TypeOf(clients), reflect.TypeOf([]models.Client{}), "Should be equal")
}

func TestGetAllError(t *testing.T) {
	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetAll").Return(nil, errors.New(""))
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	clients, clientsErr := service.GetAll()

	repository.AssertCalled(t, "GetAll")

	assert.Nil(t, clients, "Should clients be nil")
	assert.NotNil(t, clientsErr, "Shoul client error be not nil")
	assert.Equal(t, reflect.TypeOf(clientsErr), reflect.TypeOf(&customError.HTTPError{}), "Should be return http custom error")
	assert.Equal(t, clientsErr.Status, 500, "Should return 500 error status code")
}

func TestCreate(t *testing.T) {
	client := dto.CreateClientDTO{Name: "Client Name", Address: "Client Address"}

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("Create").Return(&models.Client{}, nil)
	messageServiceMock := messageService.MessageRepositoryMock{}
	messageServiceMock.On("Enqueue").Return(nil)
	service := NewClientService(&repository, &messageServiceMock)

	newClient, newClientErr := service.Create(client)

	repository.AssertCalled(t, "Create")
	messageServiceMock.AssertCalled(t, "Enqueue")

	assert.Nil(t, newClientErr, "Should new client error be nil")
	assert.NotNil(t, newClient, "Should new client not be nil")
	assert.Equal(t, reflect.TypeOf(newClient), reflect.TypeOf(&models.Client{}), "Should client created be model client type")
}

func TestCreateEmptyName(t *testing.T) {
	client := dto.CreateClientDTO{Name: "   ", Address: "Client Address"}

	repository := clientRepository.ClientRepositoryMock{}
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	newClient, newClientErr := service.Create(client)

	repository.AssertNotCalled(t, "Create")
	messageServiceMock.AssertNotCalled(t, "Enqueue")

	assert.NotNil(t, newClientErr, "Should new client error not be nil")
	assert.Nil(t, newClient, "Should new client be nil")
	assert.Equal(t, reflect.TypeOf(newClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client client error be custom http status error")
	assert.Equal(t, newClientErr.Status, 400, "Should return status code error 400")
}

func TestCreateEmptyAddress(t *testing.T) {
	client := dto.CreateClientDTO{Name: "Client Name", Address: "   "}

	repository := clientRepository.ClientRepositoryMock{}
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	newClient, newClientErr := service.Create(client)

	repository.AssertNotCalled(t, "Create")
	messageServiceMock.AssertNotCalled(t, "Enqueue")

	assert.NotNil(t, newClientErr, "Should new client error not be nil")
	assert.Nil(t, newClient, "Should new client be nil")
	assert.Equal(t, reflect.TypeOf(newClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client client error be custom http status error")
	assert.Equal(t, newClientErr.Status, 400, "Should return status code error 400")
}

func TestCreateRepositoryError(t *testing.T) {
	client := dto.CreateClientDTO{Name: "Client Name", Address: "Client Address"}

	repository := clientRepository.ClientRepositoryMock{}
	messageServiceMock := messageService.MessageRepositoryMock{}
	repository.On("Create").Return(nil, errors.New(""))
	service := NewClientService(&repository, &messageServiceMock)

	newClient, newClientErr := service.Create(client)

	repository.AssertCalled(t, "Create")
	messageServiceMock.AssertNotCalled(t, "Enqueue")

	assert.NotNil(t, newClientErr, "Should new client error not be nil")
	assert.Nil(t, newClient, "Should new client be nil")
	assert.Equal(t, reflect.TypeOf(newClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client client error be custom http status error")
	assert.Equal(t, newClientErr.Status, 400, "Should return status code error 400")
	assert.Equal(t, newClientErr.Detail, "Database", "Should client error detail be database")
}

func TestCreateMessageServiceError(t *testing.T) {
	client := dto.CreateClientDTO{Name: "Client Name", Address: "Client Address"}

	repository := clientRepository.ClientRepositoryMock{}
	messageServiceMock := messageService.MessageRepositoryMock{}
	repository.On("Create").Return(&models.Client{}, nil)
	messageServiceMock.On("Enqueue").Return(errors.New(""))
	service := NewClientService(&repository, &messageServiceMock)

	newClient, newClientErr := service.Create(client)

	repository.AssertCalled(t, "Create")
	messageServiceMock.AssertCalled(t, "Enqueue")

	assert.NotNil(t, newClientErr, "Should new client error not be nil")
	assert.Nil(t, newClient, "Should new client be nil")
	assert.Equal(t, reflect.TypeOf(newClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client client error be custom http status error")
	assert.Equal(t, newClientErr.Status, 500, "Should return status code error 500")
	assert.Equal(t, newClientErr.Detail, "Message Service", "Should client error detail be database")
}

func TestUpdate(t *testing.T) {
	clientId := "clientId"
	data := map[string]interface{}{"name": "Client Name Update", "address": "Client Address Updated"}

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(&models.Client{}, nil)
	repository.On("Update").Return(&models.Client{}, nil)
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	updatedClient, updatedClientErr := service.Update(clientId, data)

	repository.AssertCalled(t, "GetById")
	repository.AssertCalled(t, "Update")

	assert.Nil(t, updatedClientErr, "Should updated client error be nil")
	assert.NotNil(t, updatedClient, "Should updated client not be nil")
	assert.Equal(t, reflect.TypeOf(updatedClient), reflect.TypeOf(&models.Client{}), "Should client updated be model client type")
}

func TestUpdateNotFound(t *testing.T) {
	clientId := "clientId"
	data := map[string]interface{}{"name": "Client Name Update", "address": "Client Address Updated"}

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(nil, errors.New(""))
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	updatedClient, updatedClientErr := service.Update(clientId, data)

	repository.AssertCalled(t, "GetById")
	repository.AssertNotCalled(t, "Update")

	assert.NotNil(t, updatedClientErr, "Should updated client error not be nil")
	assert.Nil(t, updatedClient, "Should updated client be nil")
	assert.Equal(t, reflect.TypeOf(updatedClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client updated error be custom http error type")
	assert.Equal(t, updatedClientErr.Status, 404, "Should return status code error 404")
}

func TestUpdateError(t *testing.T) {
	clientId := "clientId"
	data := map[string]interface{}{"name": "Client Name Update", "address": "Client Address Updated"}

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(&models.Client{}, nil)
	repository.On("Update").Return(nil, errors.New(""))
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	updatedClient, updatedClientErr := service.Update(clientId, data)

	repository.AssertCalled(t, "GetById")
	repository.AssertCalled(t, "Update")

	assert.NotNil(t, updatedClientErr, "Should updated client error not be nil")
	assert.Nil(t, updatedClient, "Should updated client be nil")
	assert.Equal(t, reflect.TypeOf(updatedClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client updated error be custom http error type")
	assert.Equal(t, updatedClientErr.Status, 500, "Should return status code error 500")
	assert.Equal(t, updatedClientErr.Detail, "Database", "Should custom error detail database")
}

func TestDelete(t *testing.T) {
	clientId := "clientId"

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(&models.Client{}, nil)
	repository.On("Delete").Return(nil)
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	deletedClientErr := service.Delete(clientId)

	repository.AssertCalled(t, "GetById")
	repository.AssertCalled(t, "Delete")

	assert.Nil(t, deletedClientErr, "Should updated client error be nil")
}

func TestDeleteNotFound(t *testing.T) {
	clientId := "clientId"

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(nil, errors.New(""))
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	deletedClientErr := service.Delete(clientId)

	repository.AssertCalled(t, "GetById")
	repository.AssertNotCalled(t, "Delete")

	assert.NotNil(t, deletedClientErr, "Should updated client error not be nil")
	assert.Equal(t, reflect.TypeOf(deletedClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client updated error be custom http error type")
	assert.Equal(t, deletedClientErr.Status, 404, "Should return status code error 404")
}

func TestDeleteError(t *testing.T) {
	clientId := "clientId"

	repository := clientRepository.ClientRepositoryMock{}
	repository.On("GetById").Return(&models.Client{}, nil)
	repository.On("Delete").Return(errors.New(""))
	messageServiceMock := messageService.MessageRepositoryMock{}
	service := NewClientService(&repository, &messageServiceMock)

	deletedClientErr := service.Delete(clientId)

	repository.AssertCalled(t, "GetById")
	repository.AssertCalled(t, "Delete")

	assert.NotNil(t, deletedClientErr, "Should updated client error not be nil")
	assert.Equal(t, reflect.TypeOf(deletedClientErr), reflect.TypeOf(&customError.HTTPError{}), "Should client updated error be custom http error type")
	assert.Equal(t, deletedClientErr.Status, 500, "Should return status code error 500")
	assert.Equal(t, deletedClientErr.Detail, "Database", "Should custom error detail database")
}
