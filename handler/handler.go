package handler

import (
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
	"net/http"
	"encoding/json"
	"log"


)

type IFlashcardService interface {
	Create(e entities.FlashCardStruct) error
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

	 postCard := entities.FlashCardStruct{}

	 err := json.NewDecoder(r.Body).Decode(&postCard)
	 if err != nil {
		 log.Fatalln(err)
	 }

	 err = f.serv.Create(postCard)
	 if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	 }

	 w.WriteHeader(http.StatusCreated)
	 w.Header().Set("Content-Type", "application/json")

	 

	 
}
