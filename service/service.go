package service

import (
	"fmt"
	"strings"

	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashcardRepo interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQandA(entities.QandA) error
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
	cardType := strings.ToLower(card.Type)         //change to lowercase for comparison
	cardCategory := strings.ToLower(card.Category) //change to lowercase for comparison
	cardQuestion := card.Question
	cardOption := card.Options
	cardAnswer := card.Answer
	if cardType != "matching" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	} else {
		for _, v := range cardQuestion { //loop through the map to chack every key has a value
			if v == nil {
				return BadRequest //bad request

			}
		}
	}
	if len(cardOption) != len(cardQuestion) {
		return BadRequest //bad request
	} else {
		for _, v := range cardOption {
			if v == nil {
				return BadRequest //bad request
			}
		}
	}
	if len(cardAnswer) != len(cardQuestion) {
		return BadRequest //bad request
	} else {
		for k, v := range cardAnswer {
			if v == nil {
				return BadRequest //bad request
			} else { //where my issue is with the interface to string
				valStr := fmt.Sprint(v) //write the interface as a string
				upperCaseVal := strings.ToUpper(valStr)
				lowerCaseVal := strings.ToLower(valStr)
				upperCaseKey := strings.ToUpper(k)
				lowerCaseKey := strings.ToLower(k)
				found := false
				//The main issue
				//compare the question with the key of the cardAnswer
				if _, ok := cardQuestion[upperCaseKey]; ok {
					found = true
				}
				if _, ok := cardQuestion[lowerCaseKey]; ok {
					found = true
				}
				if !found {
					return BadRequest //bad request
				}

				//compare the option with the value of the cardAnswer
				found = false
				if _, ok := cardOption[upperCaseVal]; ok {
					found = true
				}
				if _, ok := cardOption[lowerCaseVal]; ok {
					found = true
				}
				if !found {
					return BadRequest //bad request
				}
			}

		}
	}
	return f.CreateMatching(card)

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
