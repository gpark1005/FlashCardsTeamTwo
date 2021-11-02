package service

import (
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashCardRepo interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQandA(entities.QandA) error
	GetAll() ([]entities.FlashCardStruct, error)
	GetById(string) (interface{}, error)
	UpdateMatchingById(string,entities.Matching)error
	UpdateMultipleById(string, entities.Multiple) error
	UpdateTrueFalseById(string, entities.TrueFalse) error
	UpdateInfoById(string, entities.Info) error
	UpdateQandAById(string, entities.QandA) error
	DeleteById(id string)error
}

type FlashCardService struct {
	Repo IFlashCardRepo
}

func NewFlashcardService(f IFlashCardRepo) FlashCardService {
	return FlashCardService{
		Repo: f,
	}
}

func (f FlashCardService) CreateMatching(card entities.Matching) error {
	card.SetMatchingId()
	err := ValidateMatching(card)
	if err != nil {
		return err
	}
	return f.Repo.CreateMatching(card)

}

func (f FlashCardService) CreateTrueFalse(card entities.TrueFalse) error {
	card.SetTrueFalseId()
	err := ValidateTrueFalse(card)
	if err != nil {
		return err
	}
	return f.Repo.CreateTrueFalse(card)

}

func (f FlashCardService) CreateInfo(card entities.Info) error {
	card.SetInfoId()
	err := ValidateInfo(card)
	if err != nil {
		return err
	}
	return f.Repo.CreateInfo(card)

}

func (f FlashCardService) CreateMultiple(card entities.Multiple) error {
	card.SetMultipleId()
	err := ValidateMultiple(card)
	if err != nil {
		return err
	}
	return f.Repo.CreateMultiple(card)

}

func (f FlashCardService) CreateQandA(card entities.QandA) error {
	card.SetQandAId()
	err := ValidateQandA(card)
	if err != nil {
		return err
	}
	return f.Repo.CreateQandA(card)

}

func (f FlashCardService) GetAll() ([]entities.FlashCardStruct, error) {
	allCards, err := f.Repo.GetAll()
	if err != nil {
		return allCards, BadRequest
	}

	return allCards, nil

}

func (f FlashCardService) GetById(id string) (interface{}, error) {
	deck := entities.FlashCardStruct{}
	if len(id) == 0 {
		return deck, InvalidId
	}
	return f.Repo.GetById(id)
}

func (f FlashCardService) UpdateMatchingById(id string,card entities.Matching)error {
	card.SetMatchingId()
	err := ValidateMatching(card)
	if err != nil {
		return err
	}
	return f.Repo.UpdateMatchingById(id, card)

}

func (f FlashCardService) UpdateTrueFalseById(id string,card entities.TrueFalse)error {
	card.SetTrueFalseId()
	err := ValidateTrueFalse(card)
	if err != nil {
		return err
	}
	return f.Repo.UpdateTrueFalseById(id, card)

}

func (f FlashCardService) UpdateInfoById(id string, card entities.Info) error {
	card.SetInfoId()
	err := ValidateInfo(card)
	if err != nil {
		return err
	}
	return f.Repo.UpdateInfoById(id,card)

}

func (f FlashCardService) UpdateMultipleById(id string, card entities.Multiple) error {
	card.SetMultipleId()
	err := ValidateMultiple(card)
	if err != nil {
		return err
	}
	return f.Repo.UpdateMultipleById(id, card)

}

func (f FlashCardService) UpdateQandAById(id string, card entities.QandA) error {
	card.SetQandAId()
	err := ValidateQandA(card)
	if err != nil {
		return err
	}
	return f.Repo.UpdateQandAById(id, card)

}

func (f FlashCardService)DeleteById(id string) error {
	if len(id) == 0 {
		return InvalidId
	}

	err := f.Repo.DeleteById(id)

	if err != nil {
		return err
	}
	return nil
}


