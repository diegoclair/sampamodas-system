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

//Controller holds ping handler functions
type Controller struct {
}

//NewController to handle requests
func NewController() *Controller {
	once.Do(func() {
		instance = &Controller{}
	})
	return instance
}

// handlePing - handle a Ping request
func (s *Controller) handlePing(c echo.Context) error {
	return c.JSON(200, gin.H{
		"message": "pong",
	})
}
