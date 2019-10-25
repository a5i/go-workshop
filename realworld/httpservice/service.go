package httpservice

import (
	"example.com/realworld/stor"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Service represents the HTTP service
type Service struct {
	Stor *stor.Storage
}

// SetupAPI initializes the HTTP endpoints
func (s *Service) SetupAPI(e *echo.Echo) error {
	api := e.Group("/api")
	api.GET("/articles", s.ArticleList)
	api.GET("/articles/feed", s.ArticleFeed, s.restricted)

	api.Use()

	api.POST("/users", s.UserCreate)
	api.POST("/users/login", s.UserLogin)
	return nil
}

func (s *Service) restricted(next echo.HandlerFunc) echo.HandlerFunc {
	jwtM := middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "user",
		Claims:        jwt.MapClaims{},
		SigningKey:    []byte(jwtSecret),
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Token",
	})
	return jwtM(func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		c.Set("userID", uint(id))
		return next(c)
	})
}
