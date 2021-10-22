package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
	"log"
	"net/http"
)

type IFlashcardService interface {
	Create(interface{}) error
	GetAll() (entities.FlashCardStruct, error)
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
		err = f.serv.Create(matchcard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "Multiple":
		multiplecard := entities.Multiple{}
		err := json.NewDecoder(r.Body).Decode(&multiplecard)
		if err != nil {
			log.Fatalln(err)
		}
		err = f.serv.Create(multiplecard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "TF":
		tfcard := entities.TrueFalse{}
		err := json.NewDecoder(r.Body).Decode(&tfcard)
		if err != nil {
			log.Fatalln(err)

		}
		err = f.serv.Create(tfcard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "Info":
		infocard := entities.Info{}
		err := json.NewDecoder(r.Body).Decode(&infocard)
		if err != nil {
			log.Fatalln(err)
		}

		err = f.serv.Create(infocard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}


	 w.WriteHeader(http.StatusCreated)
	 w.Header().Set("Content-Type", "application/json")


}
