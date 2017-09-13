package util

import "fmt"

type PrivIdExUnknownMethodError struct {
	MethodName string
}

func (err PrivIdExUnknownMethodError) Error() string {
	return fmt.Sprintf("Unknown method invoked: %v", err.MethodName)
}