# Step 7. Add a database

## Create a storage implementation

```shell script
mkdir stor
```

**stor/stor.go**

```go
package stor

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Storage implements a SQLite storage
type Storage struct {
	db *gorm.DB
}

func (s *Storage) Migrate() error {
	if err := s.db.AutoMigrate(&Article{}).Error; err != nil {
		return err
	}
	return nil
}

func New() (*Storage, error) {
	db, err := gorm.Open("sqlite3", "realworld.db")
	if err != nil {
		return nil, err
	}
	s := Storage{db: db}
	return &s, nil
}

```


**stor/articles.go**

```go
package stor

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Article describes an article
type Article struct {
	gorm.Model
	Slug           string `gorm:"unique_index"`
	Title          string `json:"slug"`
	Description    string `json:"description"`
	Body           string `json:"body"`
	Favorited      bool   `json:"favorited"`
	FavoritesCount int    `json:"favoritesCount"`
	AuthorID       uint   `json:"-"`
}

func (s *Storage) ArticleCreate(a *Article) error {
	if a.Slug == "" {
		a.Slug = uuid.NewV1().String()
	}
	return s.db.Create(a).Error
}

type ArticleListParams struct {
}

func (s *Storage) ArticleList(params ArticleListParams) (articles []Article, total int, err error) {
	if err := s.db.Model(&Article{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = s.db.Find(&articles).Error
	return articles, total, err
}
```

**stor/seed.go**

```go
package stor

import "github.com/brianvoe/gofakeit"

func (s *Storage) Seed() error {
	if err := s.db.Delete(&Article{}).Error; err != nil {
		return err
	}
	for i := 0; i < gofakeit.Number(10, 20); i++ {
		a := Article{
			Title:       gofakeit.Sentence(4),
			Description: gofakeit.Sentence(5),
			Body:        gofakeit.Paragraph(1, 7, 5, "\n"),
		}
		if err := s.ArticleCreate(&a); err != nil {
			return err
		}
	}
	return nil
}

```

## create a seed tool

```shell script
mkdir seed
```

**seed/main.go**

```go
package main

import (
	"log"

	"example.com/realworld/stor"
	"github.com/brianvoe/gofakeit"
)

func main() {
	gofakeit.Seed(98)
	s, err := stor.New()
	if err != nil {
		log.Panic(err)
	}
	if err := s.Migrate(); err != nil {
		log.Panic(err)
	}
	if err := s.Seed(); err != nil {
		log.Panic(err)
	}
}

```

**run it**


```shell script
go run seed/main.go
```

## Update other files

**cmd/main.go**

```go
package main

import (
	"log"

	"example.com/realworld/httpservice"
	"example.com/realworld/stor"
	"github.com/labstack/echo"
)

func mainImpl() error {
	e := echo.New()

	st, err := stor.New()
	if err != nil {
		log.Panic(err)
	}
	if err := st.Migrate(); err != nil {
		log.Panic(err)
	}
	
	s := httpservice.Service{Stor:st}
	if err := s.SetupAPI(e); err != nil {
		return err
	}
	e.Static("/", "static")
	return e.Start(":3333")
}

func main() {
	if err := mainImpl(); err != nil {
		log.Fatal(err)
	}
}

```

**httpservice/service.go**

```go
//...
type Service struct {
	Stor *stor.Storage
}
//...
``` 

**httpservice/articles.go**

```go
package httpservice

import (
	"net/http"

	"example.com/realworld/stor"
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

```