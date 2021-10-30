package Repo

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type flashcardRepo struct {
	filename string
}

func NewFlashcardRepo(filename string) flashcardRepo {
	return flashcardRepo{
		filename: filename,
	}
}

func (r flashcardRepo) CreateMatching(card entities.Matching) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.Matching = append(DbflashCardStruct.Matching, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateTrueFalse(card entities.TrueFalse) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.TrueFalse = append(DbflashCardStruct.TrueFalse, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateMultiple(card entities.Multiple) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.Multiple = append(DbflashCardStruct.Multiple, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateInfo(card entities.Info) error {

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.Info = append(DbflashCardStruct.Info, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) CreateQandA(card entities.QandA) error {
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}
	DbflashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &DbflashCardStruct)
	if err != nil {
		return ServerError
	}

	DbflashCardStruct.QandA = append(DbflashCardStruct.QandA, card)

	Marshaled, err := json.MarshalIndent(&DbflashCardStruct, "", " ")
	if err != nil {
		return ServerError
	}
	ioutil.WriteFile(r.filename, Marshaled, 0644)

	return nil
}

func (r flashcardRepo) GetAll() ([]entities.FlashCardStruct, error) {
	deck := entities.FlashCardStruct{}
	allCards := []entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return allCards, ServerError
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return allCards, ServerError
	}

	allCards = append(allCards, deck)

	return allCards, nil

}

func (r flashcardRepo) GetById(id string) (interface{}, error) {
	deck := entities.FlashCardStruct{}
	var returnDeck interface{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return returnDeck, ServerError
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return returnDeck, ServerError
	}

	returnDeck, err = IdCheck(id, deck)
	if err != nil {
		return returnDeck, NotFound

	}
	return returnDeck, nil
}

func (r flashcardRepo) UpdateMatchingById(id string, card entities.Matching) error {
	updatedCard := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &updatedCard)
	if err != nil {
		return ServerError
	}

	for i, v := range updatedCard.Matching {
		if v.Id == id {
			updatedCard.Matching = append(updatedCard.Matching[:i], updatedCard.Matching[i+1:]...)
			v.Type = card.Type
			v.Category = card.Category
			v.Question = card.Question
			v.Options = card.Options
			v.Answer = card.Answer
			updatedCard.Matching = append(updatedCard.Matching, v)

			result, err := json.MarshalIndent(&updatedCard, "", " ")
			if err != nil {
				return ServerError
			}

			err = ioutil.WriteFile(r.filename, result, 0644)
			if err != nil {
				return ServerError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashcardRepo) UpdateMultipleById(id string, card entities.Multiple) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ServerError
	}

	for i, v := range fc.Multiple {
		if v.Id == id {
			fc.Multiple = append(fc.Multiple[:i], fc.Multiple[i+1:]...)
			v.Type = card.Type
			v.Category = card.Category
			v.Question = card.Question
			v.Options = card.Options
			v.Answer = card.Answer
			fc.Multiple = append(fc.Multiple, v)

			result, err := json.MarshalIndent(&fc, "", " ")
			if err != nil {
				return ServerError
			}

			err = ioutil.WriteFile(r.filename, result, 0644)
			if err != nil {
				return ServerError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashcardRepo) UpdateInfoById(id string, card entities.Info) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ServerError
	}

	for i, v := range fc.Info {
		if v.Id == id {
			fc.Info = append(fc.Info[:i], fc.Info[i+1:]...)
			v.Type = card.Type
			v.Category = card.Category
			v.Details = card.Details
			fc.Info = append(fc.Info, v)

			result, err := json.MarshalIndent(&fc, "", " ")
			if err != nil {
				return ServerError
			}

			err = ioutil.WriteFile(r.filename, result, 0644)
			if err != nil {
				return ServerError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashcardRepo) UpdateQandAById(id string, card entities.QandA) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ServerError
	}

	for i, v := range fc.QandA {
		if v.Id == id {
			fc.QandA = append(fc.QandA[:i], fc.QandA[i+1:]...)
			v.Type = card.Type
			v.Category = card.Category
			v.Question = card.Question
			v.Answer = card.Answer
			fc.QandA = append(fc.QandA, v)

			result, err := json.MarshalIndent(&fc, "", " ")
			if err != nil {
				return ServerError
			}

			err = ioutil.WriteFile(r.filename, result, 0644)
			if err != nil {
				return ServerError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashcardRepo) UpdateTrueFalseById(id string, card entities.TrueFalse) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return ServerError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ServerError
	}

	for i, v := range fc.TrueFalse {
		if v.Id == id {
			fc.TrueFalse = append(fc.TrueFalse[:i], fc.TrueFalse[i+1:]...)
			v.Type = card.Type
			v.Category = card.Category
			v.Question = card.Question
			v.Answer = card.Answer
			fc.TrueFalse = append(fc.TrueFalse, v)

			result, err := json.MarshalIndent(&fc, "", " ")
			if err != nil {
				return ServerError
			}

			err = ioutil.WriteFile(r.filename, result, 0644)
			if err != nil {
				return ServerError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashcardRepo)DeleteById(id string) error {
	deck := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.filename) // reads flashcard database
	if err != nil {
		return err                        // checks for errors
	}

	err = json.Unmarshal(file, &deck)    // convert to Go
	if err != nil {
		return err
	}

	index, cardType, err := DeleteCheck(id, deck)
	if err != nil {
		return err
	}

	switch cardType{

	case "Matching" : 
		deck.Matching = append(deck.Matching[:index], deck.Matching[index+1:]...)
	
	case "TrueFalse" : 
		deck.TrueFalse = append(deck.TrueFalse[:index], deck.TrueFalse[index+1:]...)

	case "Info" : 
		deck.Info = append(deck.Info[:index], deck.Info[index+1:]...)

	case "Multiple" : 
		deck.Multiple = append(deck.Multiple[:index], deck.Multiple[index+1:]...)	

	case "QandA" : 
		deck.QandA = append(deck.QandA[:index], deck.QandA[index+1:]...)	
		
	}



	result, err := json.MarshalIndent(&deck, "", " ")
		if err != nil {
			return ServerError
		}

	err = ioutil.WriteFile(r.filename, result, 0644)
		if err != nil {
			return ServerError
		}

   return nil 

}
