package Repo

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type flashCardRepo struct {
	fileName string
}

func NewFlashcardRepo(filename string) flashCardRepo {
	return flashCardRepo{
		fileName: filename,
	}
}

func (r flashCardRepo) CreateMatching(card entities.Matching) error {
	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	dbFlashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &dbFlashCardStruct)
	if err != nil {
		return ConversionError
	}

	dbFlashCardStruct.Matching = append(dbFlashCardStruct.Matching, card)

	Marshaled, err := json.MarshalIndent(&dbFlashCardStruct, "", " ")
	if err != nil {
		return ConversionError
	}
	err = ioutil.WriteFile(r.fileName, Marshaled, 0644)
	if err != nil {
		return StorageError
	}
	return nil
}

func (r flashCardRepo) CreateTrueFalse(card entities.TrueFalse) error {
	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	dbFlashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &dbFlashCardStruct)
	if err != nil {
		return ConversionError
	}

	dbFlashCardStruct.TrueFalse = append(dbFlashCardStruct.TrueFalse, card)

	Marshaled, err := json.MarshalIndent(&dbFlashCardStruct, "", " ")
	if err != nil {
		return ConversionError
	}
	err = ioutil.WriteFile(r.fileName, Marshaled, 0644)
	if err != nil{
		return StorageError
	}
	return nil
}

func (r flashCardRepo) CreateMultiple(card entities.Multiple) error {
	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	dbFlashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &dbFlashCardStruct)
	if err != nil {
		return ConversionError
	}

	dbFlashCardStruct.Multiple = append(dbFlashCardStruct.Multiple, card)

	Marshaled, err := json.MarshalIndent(&dbFlashCardStruct, "", " ")
	if err != nil {
		return ConversionError
	}
	err = ioutil.WriteFile(r.fileName, Marshaled, 0644)
	if err != nil{
		return StorageError
	}
	return nil
}

func (r flashCardRepo) CreateInfo(card entities.Info) error {

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	dbFlashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &dbFlashCardStruct)
	if err != nil {
		return ConversionError
	}

	dbFlashCardStruct.Info = append(dbFlashCardStruct.Info, card)

	Marshaled, err := json.MarshalIndent(&dbFlashCardStruct, "", " ")
	if err != nil {
		return ConversionError
	}
	err = ioutil.WriteFile(r.fileName, Marshaled, 0644)
	if err != nil{
		return StorageError
	}
	return nil
}

func (r flashCardRepo) CreateQandA(card entities.QandA) error {
	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	dbFlashCardStruct := entities.FlashCardStruct{}
	err = json.Unmarshal(file, &dbFlashCardStruct)
	if err != nil {
		return ConversionError
	}

	dbFlashCardStruct.QandA = append(dbFlashCardStruct.QandA, card)

	Marshaled, err := json.MarshalIndent(&dbFlashCardStruct, "", " ")
	if err != nil {
		return ConversionError
	}
	err = ioutil.WriteFile(r.fileName, Marshaled, 0644)
	if err != nil{
		return StorageError
	}
	return nil
}

func (r flashCardRepo) GetAll() ([]entities.FlashCardStruct, error) {
	deck := entities.FlashCardStruct{}
	allCards := []entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return allCards, FileError
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return allCards, ConversionError
	}

	allCards = append(allCards, deck)

	return allCards, nil

}

func (r flashCardRepo) GetById(id string) (interface{}, error) {
	deck := entities.FlashCardStruct{}
	var cardFound interface{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return cardFound, FileError
	}

	err = json.Unmarshal(file, &deck)
	if err != nil {
		return cardFound, ConversionError
	}

	cardFound, err = IdCheck(id, deck)
	if err != nil {
		return cardFound, NotFound

	}
	return cardFound, nil
}

func (r flashCardRepo) UpdateMatchingById(id string, card entities.Matching) error {
	updatedCard := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	err = json.Unmarshal(file, &updatedCard)
	if err != nil {
		return ConversionError
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
				return ConversionError
			}

			err = ioutil.WriteFile(r.fileName, result, 0644)
			if err != nil {
				return StorageError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashCardRepo) UpdateMultipleById(id string, card entities.Multiple) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ConversionError
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
				return ConversionError
			}

			err = ioutil.WriteFile(r.fileName, result, 0644)
			if err != nil {
				return StorageError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashCardRepo) UpdateInfoById(id string, card entities.Info) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ConversionError
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
				return ConversionError
			}

			err = ioutil.WriteFile(r.fileName, result, 0644)
			if err != nil {
				return StorageError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashCardRepo) UpdateQandAById(id string, card entities.QandA) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ConversionError
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
				return ConversionError
			}

			err = ioutil.WriteFile(r.fileName, result, 0644)
			if err != nil {
				return StorageError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashCardRepo) UpdateTrueFalseById(id string, card entities.TrueFalse) error {
	fc := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return FileError
	}

	err = json.Unmarshal(file, &fc)
	if err != nil {
		return ConversionError
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
				return ConversionError
			}

			err = ioutil.WriteFile(r.fileName, result, 0644)
			if err != nil {
				return StorageError
			}
			return nil

		}
	}

	return NotFound

}

func (r flashCardRepo)DeleteById(id string) error {
	deck := entities.FlashCardStruct{}

	file, err := ioutil.ReadFile(r.fileName) // reads flashcard database
	if err != nil {
		return FileError                        // checks for errors
	}

	err = json.Unmarshal(file, &deck)    // convert to Go
	if err != nil {
		return ConversionError
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
			return ConversionError
		}

	err = ioutil.WriteFile(r.fileName, result, 0644)
		if err != nil {
			return StorageError
		}

   return nil 

}
