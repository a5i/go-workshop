# Step 9. Add auth middleware

Update **httpservice/service.go**

```go
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

```

Update **httpservice/articles.go**

```go
package httpservice

import (
	"net/http"

	"example.com/realworld/stor"
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo"
)

type article struct {
	stor.Article
	Author struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"author"`
}

type articleListResponse struct {
	Articles      []article `json:"articles"`
	ArticlesCount int       `json:"articlesCount"`
}

func (s *Service) ArticleList(c echo.Context) error {
	var resp articleListResponse
	articles, total, err := s.Stor.ArticleList(stor.ArticleListParams{})
	if err != nil {
		return err
	}
	resp.ArticlesCount = total
	for _, a := range articles {
		art := article{
			Article: a,
		}
		resp.Articles = append(resp.Articles, art)
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Service) ArticleFeed(c echo.Context) error {
	var resp articleListResponse
	for i := 0; i < gofakeit.Number(10, 20); i++ {
		a := article{
			Article: stor.Article{
				Slug:        gofakeit.BeerName(),
				Title:       gofakeit.Sentence(4),
				Description: gofakeit.Sentence(5),
				Body:        gofakeit.Paragraph(1, 7, 5, "\n"),
			},
		}
		resp.Articles = append(resp.Articles, a)
	}
	return c.JSON(http.StatusOK, resp)
}
```

## Run

```shell script
go run cmd/main.go
```

Open a browser http://127.0.0.1:3333