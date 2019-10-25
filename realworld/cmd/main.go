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

	s := httpservice.Service{Stor: st}
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
