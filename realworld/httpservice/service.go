package httpservice

import (
	"example.com/realworld/stor"
	"github.com/labstack/echo"
)

// Service represents the HTTP service
type Service struct {
	Stor *stor.Storage
}

// SetupAPI initializes the HTTP endpoints
func (s *Service) SetupAPI(e *echo.Echo) error {
	api := e.Group("/api")
	api.GET("/articles", s.ArticleList)

	api.POST("/users", s.UserCreate)
	api.POST("/users/login", s.UserLogin)
	return nil
}
