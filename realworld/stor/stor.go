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
	if err := s.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
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
