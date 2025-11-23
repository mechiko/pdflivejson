package usecase

import (
	"pdflive/domain"
)

const modError = "usecase"

type usecase struct {
	domain.Apper
}

func New(app domain.Apper) *usecase {
	return &usecase{
		Apper: app,
	}
}
