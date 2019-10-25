package httpservice

import (
	"net/http"
	"time"

	"example.com/realworld/stor"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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
