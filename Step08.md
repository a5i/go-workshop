# Step 8. Add users

## Update the storage

Create **stor/users.go**

```go
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
	if err!=nil {
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
```

Update **stor/storage.go**

```go
func (s *Storage) Migrate() error {
	if err := s.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	// ...
	return nil
}
```

Update **stor/seed.go**

```go
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
// ...
}
```

## Update the HTTP service

Create **httpservice/users.go**

```go
package httpservice

import (
	"net/http"
	"time"

	"example.com/realworld/stor"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

type userCreateRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

const jwtSecret = "gjsdt67588ids"

func getTokenForUser(u *stor.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString([]byte(jwtSecret))
}

func (s *Service) UserCreate(c echo.Context) error {
	var body userCreateRequest
	if err := c.Bind(&body); err != nil {
		return err
	}
	user := stor.User{
		Username: body.User.Username,
		Email:    body.User.Email,
	}
	if err := user.SetPassword(body.User.Password); err != nil {
		return err
	}
	if err := s.Stor.UserCreate(&user); err != nil {
		return err
	}
	var resp struct {
		User struct {
			stor.User
			Token string `json:"token"`
		} `json:"user"`
	}
	resp.User.User = user
	var err error
	resp.User.Token, err = getTokenForUser(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (s *Service) UserLogin(c echo.Context) error {
	var body userLoginRequest
	if err := c.Bind(&body); err != nil {
		return err
	}
	user, err := s.Stor.UserGetByEmail(body.User.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "user does not exists or password incorrect")
	}
	if !user.CheckPassword(body.User.Password) {
		return echo.NewHTTPError(http.StatusForbidden, "user does not exists or password incorrect")
	}
	var resp struct {
		User struct {
			stor.User
			Token string `json:"token"`
		} `json:"user"`
	}
	resp.User.User = user
	resp.User.Token, err = getTokenForUser(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

```


Update **httpservice/service.go**

```go
// SetupAPI initializes the HTTP endpoints
func (s *Service) SetupAPI(e *echo.Echo) error {
	// ...
	api.POST("/users", s.UserCreate)
	api.POST("/users/login", s.UserLogin)
	return nil
}
```

## Run

```shell script
go run cmd/main.go
```

Open a browser http://127.0.0.1:3333