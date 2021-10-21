package handler

import (
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashcardService interface {
	Create(e entities.FlashCardStruct) error
	GetAll() error
	GetById(string)
}

type FlashcardHandler struct {
	serv IFlashcardService
}

func NewFlashcardHandler(f IFlashcardService) FlashcardHandler {
	return FlashcardHandler{
		serv: f,
	}
}
