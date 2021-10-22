package service

import (
	"reflect"

	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashcardRepo interface {
	Create(interface{}) error
}

type FlashcardService struct {
	Repo IFlashcardRepo
}

func NewFlashcardService(f IFlashcardRepo) FlashcardService {
	return FlashcardService{
		Repo: f,
	}
}

func (f FlashcardService) Create(card interface{}) error {

	switch reflect.TypeOf(card) {
	case reflect.TypeOf(entities.Matching):
	}
}