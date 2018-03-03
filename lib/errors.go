package serverless

// ServiceError struct
type ServiceError struct {
	msg  string // error description
	Code int    // where the error occurred
}

func (e *ServiceError) Error() string {
	return e.msg
}

// Error func to init the ServiceError struct
func Error(text string, code int) *ServiceError {
	return &ServiceError{text, code}
}
