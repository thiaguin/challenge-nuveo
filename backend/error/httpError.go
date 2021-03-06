package errors

// HTTPError
type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}

	return e.Detail + " : " + e.Cause.Error()
}

func NewHTTPError(err error, status int, detail string) *HTTPError {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}
