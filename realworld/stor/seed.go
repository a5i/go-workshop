package stor

import "github.com/brianvoe/gofakeit"

func (s *Storage) Seed() error {
	if err := s.db.Delete(&User{}).Error; err != nil {
		return err
	}
	user := User{
		Username: "test",
		Email:    "test@example.com",
		Bio:      gofakeit.Sentence(45),
		Image:    nil,
	}
	if err := user.SetPassword("test"); err != nil {
		return err
	}
	if err := s.UserCreate(&user); err != nil {
		return err
	}

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
