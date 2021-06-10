package pingroute

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

type Controller struct {
}

func NewController() *Controller {
	once.Do(func() {
		instance = &Controller{}
	})
	return instance
}

func (s *Controller) handlePing(c echo.Context) error {
	return c.JSON(200, gin.H{
		"message": "pong",
	})
}
