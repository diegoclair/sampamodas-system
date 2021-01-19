// Package errors has utilities to create and extend errors, easing the error handling process
package errors

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type errorString string

func (err errorString) Error() string {
	return string(err)
}

// ErrNotFound represents a record not found error
const ErrNotFound = errorString("Item não encontrado")

// errorWrapper defines the interface for an error wrapper that extends an error with additional information
type errorWrapper interface {
	Error() string
	GetOriginalError() error
}

// WrappedError holds an error wrapped with a context message
type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

// Error returns the string representation of the error
func (err WrappedError) Error() string {
	if len(err.messages) > 0 {
		retVal := fmt.Sprintf("%s: ", err.path)

		for _, message := range err.messages {
			retVal += message + "; "
		}

		return fmt.Sprintf("%s => %v", retVal, err.originalError)
	}

	return fmt.Sprintf("%s => %v", err.path, err.originalError)
}

// GetOriginalError returns the original error
func (err WrappedError) GetOriginalError() error {
	if err.originalError != nil {
		originalError, ok := (err.originalError).(errorWrapper)
		if ok {
			return originalError.GetOriginalError()
		}
	}

	return err.originalError
}

// Wrap wraps an error with a context message and adds execution path
func Wrap(err error, messages ...string) error {
	if err != nil {
		// get caller function path
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		funcRef := runtime.FuncForPC(pc[0])

		pathArr := strings.Split(funcRef.Name(), "/")

		path := pathArr[len(pathArr)-1]

		return &WrappedError{
			originalError: err,
			path:          path,
			messages:      messages,
		}
	}

	return nil
}

// GetOriginalError returns the original error if the provided error is a WrappedError.
// Returns the provided error otherwise
func GetOriginalError(err error) error {
	wrappedErr, ok := err.(errorWrapper)
	if ok {
		return wrappedErr.GetOriginalError()
	}

	return err
}

// NewValidationError returns a ValidationError instance with the provided parameters
func NewValidationError(field string, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// NewNullArgumentError returns a preformatted error for null arguments
func NewNullArgumentError(argumentName string) *NullArgumentError {
	return &NullArgumentError{argumentName}
}

// NewApplicationError returns a ApplicationError instance
func NewApplicationError(path string, message string) *ApplicationError {
	return &ApplicationError{
		Message: message,
		Path:    path,
	}
}

// NewNotAuthorizedError returns a NotAuthorizedError instance
func NewNotAuthorizedError(message string) *NotAuthorizedError {
	return &NotAuthorizedError{
		Message: message,
	}
}

// NewForbiddenError returns a ForbiddenError instance
func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{
		Message: message,
	}
}

// NewConflictError returns a ConflictError instance
func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}

// NewHTTPError returns a HTTPError instance
func NewHTTPError(httpStatus int, message string) *HTTPError {
	return &HTTPError{
		Status:  httpStatus,
		Message: message,
	}
}

// NewNotFoundError returns a Not Found HTTPError instance
func NewNotFoundError() *HTTPError {
	return NewHTTPError(http.StatusNotFound, "Not Found")
}

// HTTPError represents a generic HTTP error response
type HTTPError struct {
	Status  int
	Message string
}

// Error returns the string representation of the error
func (e *HTTPError) Error() string {
	return fmt.Sprintf("%v - %s", e.Status, e.Message)
}

// NullArgumentError represents a error that is used to sinalize that a provided argument is null
type NullArgumentError struct {
	ArgumentName string
}

// Error returns the string representation of the error
func (e *NullArgumentError) Error() string {
	return fmt.Sprintf("Parameter %s can't be null", e.ArgumentName)
}

// ValidationError represents an input validation error
type ValidationError struct {
	Field   string            `json:"field_name,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

func (e *ValidationError) Error() string {
	var errorList []string

	for _, err := range e.Errors {
		errorList = append(errorList, err.Error())
	}

	output := e.Message

	if len(errorList) > 0 {
		output += fmt.Sprintf("\n - %v", strings.Join(errorList, ";\n - "))
	}

	return fmt.Sprintf("%v: %v", e.Field, output)
}

// AddError adds a new validation error to the chain
func (e *ValidationError) AddError(field string, message string) {
	e.Errors = append(e.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// ApplicationError represents a common applicatino error structure
type ApplicationError struct {
	Message string
	Path    string
}

// Error returns the string representation of the error
func (e *ApplicationError) Error() string {
	return e.Message
}

// NotAuthorizedError represents an access restriction error structure
type NotAuthorizedError struct {
	Message string
}

// Error returns the string representation of the error
func (e *NotAuthorizedError) Error() string {
	return e.Message
}

// ForbiddenError represents an forbidden error structure
type ForbiddenError struct {
	Message string
}

// Error returns the string representation of the error
func (e *ForbiddenError) Error() string {
	return e.Message
}

// ConflictError represents an conflict error structure
type ConflictError struct {
	Message string
}

// Error returns the string representation of the error
func (e *ConflictError) Error() string {
	return e.Message
}
