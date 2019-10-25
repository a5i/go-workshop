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
