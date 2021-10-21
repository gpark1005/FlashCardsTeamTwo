package service


import (
	"github.com/gpark1005/FlashCardsTeamTwo/entities"


)

type IFlashcardRepo interface {
	Create(e entities.FlashCardStruct) error 
}

type FlashcardService struct {
	Repo IFlashcardRepo
}

func NewFlashcardService(f IFlashcardRepo) FlashcardService {
	return FlashcardService{
		Repo : f, 
	}
}

