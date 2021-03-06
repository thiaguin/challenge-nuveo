package services

/// FileServiceInterface interface
type FileServiceInterface interface {
	Write(message []byte, directory string) error
	Exist(path string) (bool, error)
}
