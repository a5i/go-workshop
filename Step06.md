# Step 6. Add list of Articles and Tags

**httpservice/arcticles.go**

```go
package httpservice

import (
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo"
)

type article struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         struct {
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
	for i := 0; i < gofakeit.Number(10, 20); i++ {
		a := article{
			Slug:        gofakeit.BeerName(),
			Title:       gofakeit.Sentence(4),
			Description: gofakeit.Sentence(5),
			Body:        gofakeit.Paragraph(1, 7, 5, "\n"),
			CreatedAt:   gofakeit.Date(),
			UpdatedAt:   gofakeit.Date(),
		}
		resp.Articles = append(resp.Articles, a)
	}
	return c.JSON(http.StatusOK, resp)
}

```

**httpservice/service.go**

```go
package httpservice

import (
	"github.com/labstack/echo"
)

// Service represents the HTTP service
type Service struct {
}

// SetupAPI initializes the HTTP endpoints
func (s *Service) SetupAPI(e *echo.Echo) error {
	api := e.Group("/api")
	api.GET("/articles", s.ArticleList)
	return nil
}
```

## Run

```shell script
go run cmd/main.go
```

Open a browser http://127.0.0.1:3333