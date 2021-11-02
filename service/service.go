package service

import (
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashCardRepo interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQnA(entities.QnA) error
	GetAll() ([]entities.FlashCardStruct, error)
	GetById(string) (interface{}, error)
	UpdateMatchingById(string,entities.Matching)error
	UpdateMultipleById(string, entities.Multiple) error
	UpdateTrueFalseById(string, entities.TrueFalse) error
	UpdateInfoById(string, entities.Info) error
	UpdateQnAById(string, entities.QnA) error
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

func (f FlashCardService) CreateQnA(card entities.QnA) error {
	card.SetQnAId()
	err := ValidateQnA(card)
	if err != nil {
		return err
	}
	return f.Repo.CreateQnA(card)

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

func (f FlashCardService) UpdateQnAById(id string, card entities.QnA) error {
	card.SetQnAId()
	err := ValidateQnA(card)
	if err != nil {
		return err
	}
	return f.Repo.UpdateQnAById(id, card)

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


