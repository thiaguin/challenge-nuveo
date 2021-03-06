package services

// MessageServiceInterface interface
type MessageServiceInterface interface {
	Enqueue(message []byte) error
}
