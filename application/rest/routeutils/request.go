package routeutils

import (
	"fmt"
	"strconv"

	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/labstack/echo/v4"
)

func GetAndValidateInt64Param(c echo.Context, paramName string, isRequired bool) (param int64, err error) {

	stringParam := c.Param(paramName)
	if isRequired && stringParam == "" {
		return param, resterrors.NewBadRequestError(fmt.Sprintf("The route parameter '%s' is required", paramName), nil)
	}

	param, err = strconv.ParseInt(stringParam, 10, 64)
	if err != nil {
		return param, resterrors.NewBadRequestError(fmt.Sprintf("The route parameter '%s' should be a number", paramName), err)
	}

	return param, err
}

func GetAndValidateStringParam(c echo.Context, paramName string, isRequired bool) (param string, err error) {

	param = c.Param(paramName)
	if isRequired && param == "" {
		return param, resterrors.NewBadRequestError(fmt.Sprintf("The route parameter '%s' is required", paramName), err)
	}

	return param, err
}
