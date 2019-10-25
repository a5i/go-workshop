package stor

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username" json:"username"`
	Email        string  `gorm:"column:email;unique_index" json:"email"`
	Bio          string  `gorm:"column:bio;size:1024" json:"bio"`
	Image        *string `gorm:"column:image" json:"image"`
	PasswordHash string  `gorm:"column:password;not null" json:"-"`
}

func (u *User) SetPassword(password string) error {
	pHash, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	if err != nil {
		return err
	}
	u.PasswordHash = base64.StdEncoding.EncodeToString(pHash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	pHash, err := base64.StdEncoding.DecodeString(u.PasswordHash)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword(pHash, []byte(password)) == nil
}

func (s *Storage) UserCreate(u *User) error {
	return s.db.Create(u).Error
}

func (s *Storage) UserGetByEmail(email string) (user User, err error) {
	err = s.db.Where("email = ?", email).Take(&user).Error
	return
}
