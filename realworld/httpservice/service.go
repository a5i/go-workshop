package httpservice

import (
	"net/http"

	"github.com/labstack/echo"
)

// Service represents the HTTP service
type Service struct {
}

// SetupAPI initializes the HTTP endpoints
func (s *Service) SetupAPI(e *echo.Echo) error {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "RealWorld!")
	})
	return nil
}
