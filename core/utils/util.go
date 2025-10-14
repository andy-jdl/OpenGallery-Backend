package core

import (
	"fmt"
)

type InternalServerError struct {
	Code    int
	Message string
}

type NotFoundError struct {
	Code    int
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Error: %d: %s", e.Code, e.Message)
}
