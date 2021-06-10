package routeutils

import (
	"net/http"

	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/labstack/echo/v4"
)

const ErrorMessageServiceUnavailable = "Serviço Indisponível"

// ResponseNoContent returns a standard API success with no content response
func ResponseNoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// ResponseCreated returns a standard API successful as a result, a resource has been created
func ResponseCreated(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

// ResponseAPIOK returns a standard API success response
func ResponseAPIOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

// ResponseAPINotFoundError returns a standard API not found error to the response
func ResponseAPINotFoundError(c echo.Context) error {
	return ResponseAPIError(c, http.StatusNotFound, "Not Found", "", nil)
}

// ResponseAPIError returns a standard API error to the response
func ResponseAPIError(c echo.Context, status int, message string, err string, causes interface{}) error {
	returnValue := resterrors.NewRestError(message, status, err, causes)
	return c.JSON(status, returnValue)
}

func HandleAPIError(c echo.Context, errorToHandle error) (err error) {
	statusCode := http.StatusServiceUnavailable
	errorMessage := ErrorMessageServiceUnavailable

	if errorToHandle != nil {
		logger.Error("HandleAPIError: ", errorToHandle)

		errorString := errorToHandle.Error()

		restErr, ok := errorToHandle.(resterrors.RestErr)
		if !ok {
			return ResponseAPIError(c, statusCode, errorMessage, errorString, nil)
		}

		return c.JSON(restErr.StatusCode(), restErr)

	}

	return ResponseAPIError(c, statusCode, errorMessage, "", nil)
}
