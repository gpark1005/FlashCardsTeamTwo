package service

import (
	"fmt"
	"strings"

	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

func ValidateMatching(m entities.Matching) error {
	cardType := strings.ToLower(m.Type)         //change to lowercase for comparison
	cardCategory := strings.ToLower(m.Category) //change to lowercase for comparison
	cardQuestion := m.Question
	cardOption := m.Options
	cardAnswer := m.Answer
	if cardType != "matching" {
		return InvalidType
	}
	if cardCategory != "golang" {
		return InvalidCategory
	}
	if len(cardQuestion) == 0 {
		return QuestionsInaccurate
	} else {
		for _, v := range cardQuestion { //loop through the map to check every key has a value
			if len(v.(string)) == 0 {
				return InvalidQuestion

			}
		}
	}
	if len(cardOption) != len(cardQuestion) {
		return UnacceptableFormat
	} else {
		for _, v := range cardOption {
			if len(v.(string)) == 0 {
				return ErroneousOptions
			}
		}
	}
	if len(cardAnswer) != len(cardQuestion) {
		return UnacceptableFormat
	} else {
		for k, v := range cardAnswer {
			if len(v.(string)) == 0 {
				return FaultyAnswers
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
					return MisMatch
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
					return MisMatch
				}
			}

		}
	}
	return nil
}

func ValidateTrueFalse(tf entities.TrueFalse) error {
	cardType := strings.ToLower(tf.Type)
	cardCategory := strings.ToLower(tf.Category)
	cardQuestion := tf.Question
	cardAnswer := strings.ToLower(tf.Answer)

	if cardType != "truefalse" {
		return InvalidType
	}

	if cardCategory != "golang" {
		return InvalidCategory
	}

	if len(cardQuestion) == 0 {
		return InvalidQuestion
	}
	if cardAnswer != "true" && cardAnswer != "false" {
		return AnswerNotAccepted

	}
	return nil
}

func ValidateInfo(i entities.Info) error {
	cardType := strings.ToLower(i.Type)
	cardCategory := strings.ToLower(i.Category)
	cardDetails := i.Details

	if cardType != "info" {
		return InvalidType
	}
	if cardCategory != "golang" {
		return InvalidCategory
	}
	if len(cardDetails) == 0 {
		return InvalidData
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
		return InvalidType
	}
	if cardCategory != "golang" {
		return InvalidCategory
	}
	if len(cardQuestion) == 0 {
		return InvalidQuestion
	}
	if len(cardOption) == 0 {
		return ErroneousOptions
	} else {
		for _, v := range cardOption {
			if len(v.(string)) == 0 {
				return ErroneousOptions
			}
		}
	}
	if len(cardAnswer) == 0 {
		return InvalidAnswer
	} else if _, ok := cardOption[cardAnswer]; !ok {
		cardType := strings.ToLower(mu.Type)
		cardCategory := strings.ToLower(mu.Category)
		cardQuestion := mu.Question
		cardOption := mu.Options
		cardAnswer := mu.Answer

		if cardType != "multiple" {
			return InvalidType
		}
		if cardCategory != "golang" {
			return InvalidCategory
		}
		if len(cardQuestion) == 0 {
			return InvalidQuestion
		}
		if len(cardOption) == 0 {
			return ErroneousOptions
		} else {
			for _, v := range cardOption {
				if len(v.(string)) == 0 {
					return ErroneousOptions
				}
			}
		}
		if len(cardAnswer) == 0 {
			return InvalidAnswer
		} else if _, ok := cardOption[cardAnswer]; !ok {
			return AnswerNotFound
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
		return InvalidType
	}
	if cardCategory != "golang" {
		return InvalidCategory
	}
	if len(cardQuestion) == 0 {
		return InvalidQuestion
	}
	if len(cardAnswer) == 0 {
		return InvalidAnswer
	}
	return nil
}
