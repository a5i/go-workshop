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
