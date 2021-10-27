package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashcardService interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQandA(entities.QandA) error
	GetAll() ([]entities.FlashCardStruct, error)
	GetById(string) ([]entities.FlashCardStruct, error)
}

type FlashcardHandler struct {
	serv IFlashcardService
}

func NewFlashcardHandler(f IFlashcardService) FlashcardHandler {
	return FlashcardHandler{
		serv: f,
	}
}

func (f FlashcardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardtype := vars["Type"]

	switch cardtype {
	case "Matching":
		matchcard := entities.Matching{}
		err := json.NewDecoder(r.Body).Decode(&matchcard)
		if err != nil {
			log.Fatalln(err)
		}
		err = f.serv.CreateMatching(matchcard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "Multiple":
		multiplecard := entities.Multiple{}
		err := json.NewDecoder(r.Body).Decode(&multiplecard)
		if err != nil {
			log.Fatalln(err)
		}
		err = f.serv.CreateMultiple(multiplecard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "TrueFalse":
		tfcard := entities.TrueFalse{}
		err := json.NewDecoder(r.Body).Decode(&tfcard)
		if err != nil {
			log.Fatalln(err)

		}
		err = f.serv.CreateTrueFalse(tfcard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "Info":
		infocard := entities.Info{}
		err := json.NewDecoder(r.Body).Decode(&infocard)
		if err != nil {
			log.Fatalln(err)
		}

		err = f.serv.CreateInfo(infocard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "QandA":
		QandAcard := entities.QandA{}
		err := json.NewDecoder(r.Body).Decode(&QandAcard)
		if err != nil {
			log.Fatalln(err)
		}

		err = f.serv.CreateQandA(QandAcard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}

func (f FlashcardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	allCards, err := f.serv.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	allCardsBytes, err := json.MarshalIndent(allCards, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	w.Write(allCardsBytes)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}

func (f FlashcardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"]

	f.serv.GetById(cardId)

}
