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
