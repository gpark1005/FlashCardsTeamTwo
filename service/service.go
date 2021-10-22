package service

import (
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

func (f FlashcardService) CreateMatching(card entities.Matching) error {
	card.SetMatchingId()
	return nil //placeholder

}

func (f FlashcardService) CreateTrueFalse(card entities.TrueFalse) error {
	card.SetTrueFalseId()
	return nil
}

func (f FlashcardService) CreateInfo(card entities.Info) error {
	card.SetInfoId()
	return nil

}

func (f FlashcardService) CreateMultiple(card entities.Multiple) error {
	card.SetMultipleId()
	return nil

}
