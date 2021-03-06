package services

/// MessageServiceInterface interface
type MessageServiceInterface interface {
	Dequeue() ([]byte, error)
}
