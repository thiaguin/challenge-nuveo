package services

import (
	"github.com/stretchr/testify/mock"
)

// MessageRepositoryMock type
type MessageRepositoryMock struct {
	mock.Mock
}

func (m *MessageRepositoryMock) Enqueue(message []byte) error {
	args := m.Called()
	return args.Error(0)
}
