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
	resp.Articles = make([]article, 0, len(articles))
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
