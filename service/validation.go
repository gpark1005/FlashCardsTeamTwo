package service

import (
	"fmt"
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
	"strings"
)



func ValidateMatching(m entities.Matching)error {
	cardType := strings.ToLower(m.Type)         //change to lowercase for comparison
	cardCategory := strings.ToLower(m.Category) //change to lowercase for comparison
	cardQuestion := m.Question
	cardOption := m.Options
	cardAnswer := m.Answer
	if cardType != "matching" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	} else {
		for _, v := range cardQuestion { //loop through the map to check every key has a value
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
	return nil
}

func ValidateTrueFalse(tf entities.TrueFalse)error {
	cardType := strings.ToLower(tf.Type)
	cardCategory := strings.ToLower(tf.Category)
	cardQuestion := tf.Question
	cardAnswer := strings.ToLower(tf.Answer)

	if cardType != "truefalse" {
		return BadRequest
	}

	if cardCategory != "golang" {
		return BadRequest
	}

	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	}
	if cardAnswer != "true" && cardAnswer != "false" {
		return BadRequest

	}
	return nil
}

func ValidateInfo(i entities.Info) error {
	cardType := strings.ToLower(i.Type)
	cardCategory := strings.ToLower(i.Category)
	cardDetails := i.Details

	if cardType != "info" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardDetails) == 0 {
		return BadRequest
	}
	return nil
}

func ValidateMultiple(mu entities.Multiple) error {
	cardType := strings.ToLower(mu.Type)
	cardCategory := strings.ToLower(mu.Category)
	cardQuestion := mu.Question
	cardOption := mu.Options
	cardAnswer := mu.Answer

	if cardType != "multiple" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0  {
		return BadRequest //bad request
	}
	if len(cardOption) == 0 {
		return BadRequest //bad request
	} else {
		for _, v := range cardOption {
			if v == nil {
				return BadRequest //bad request
			}
		}
	}
	if len(cardAnswer) == 0 {
		return BadRequest
	} else if _, ok := cardOption[cardAnswer]; !ok {
		cardType := strings.ToLower(mu.Type)
		cardCategory := strings.ToLower(mu.Category)
		cardQuestion := mu.Question
		cardOption :=mu.Options
		cardAnswer := mu.Answer

		if cardType != "multiple" {
			return BadRequest //bad request
		}
		if cardCategory != "golang" {
			return BadRequest //bad request
		}
		if len(cardQuestion) == 0 {
			return BadRequest //bad request
		}
		if len(cardOption) == 0 {
			return BadRequest //bad request
		} else {
			for _, v := range cardOption {
				if v == nil {
					return BadRequest //bad request
				}
			}
		}
		if len(cardAnswer) == 0 {
			return BadRequest
		} else if _, ok := cardOption[cardAnswer]; !ok {
			return BadRequest
		}
	}
	return nil
}

func ValidateQandA(q entities.QandA) error {
	cardType := strings.ToLower(q.Type)
	cardCategory := strings.ToLower(q.Category)
	cardQuestion := q.Question
	cardAnswer := q.Answer

	if cardType != "qanda" {
		return BadRequest //bad request
	}
	if cardCategory != "golang" {
		return BadRequest //bad request
	}
	if len(cardQuestion) == 0 {
		return BadRequest //bad request
	}
	if len(cardAnswer) == 0  {
		return BadRequest //bad request
	}
	return nil
}