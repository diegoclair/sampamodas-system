package authroute

import (
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routeutils"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/viewmodel"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/domain/service"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

type Controller struct {
	authService service.AuthService
	mapper      mapper.Mapper
}

func NewController(authService service.AuthService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			authService: authService,
			mapper:      mapper,
		}
	})
	return instance
}

func (s *Controller) handleSignup(c echo.Context) error {

	input := viewmodel.NewUser{}

	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	userSignup := entity.User{}

	err = s.mapper.From(input).To(&userSignup)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = s.authService.Signup(userSignup)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}
