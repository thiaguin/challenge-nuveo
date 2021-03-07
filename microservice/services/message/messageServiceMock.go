package services

import (
	"github.com/stretchr/testify/mock"
)

// MessageRepositoryMock type
type MessageRepositoryMock struct {
	mock.Mock
}

func (m *MessageRepositoryMock) Dequeue() ([]byte, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return []byte{}, args.Error(1)
}
