package services

import (
	"github.com/stretchr/testify/mock"
)

// FileRepositoryMock type
type FileRepositoryMock struct {
	mock.Mock
}

func (m *FileRepositoryMock) Write(message []byte, directory string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *FileRepositoryMock) Exist(path string) (bool, error) {
	args := m.Called()
	return true, args.Error(1)
}
