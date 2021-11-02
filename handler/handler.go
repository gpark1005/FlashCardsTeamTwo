package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gpark1005/FlashCardsTeamTwo/entities"
)

type IFlashCardService interface {
	CreateMatching(entities.Matching) error
	CreateTrueFalse(entities.TrueFalse) error
	CreateMultiple(entities.Multiple) error
	CreateInfo(entities.Info) error
	CreateQnA(entities.QnA) error
	GetAll() ([]entities.FlashCardStruct, error)
	GetById(string) (interface{}, error)
	UpdateMatchingById(string, entities.Matching) error
	UpdateMultipleById(string, entities.Multiple) error
	UpdateTrueFalseById(string, entities.TrueFalse) error
	UpdateInfoById(string, entities.Info) error
	UpdateQnAById(string, entities.QnA) error
	DeleteById(id string) error
}

type FlashCardHandler struct {
	serv IFlashCardService
}

func NewFlashcardHandler(f IFlashCardService) FlashCardHandler {
	return FlashCardHandler{
		serv: f,
	}
}

func (f FlashCardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var cardMap map[string]interface{}
	file, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(file, &cardMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if cardType, ok := cardMap["Type"]; ok {
		switch cardType {
		case "Matching":
			matchCard := entities.Matching{}
			err := json.Unmarshal(file, &matchCard)
			if err != nil {
				log.Fatalln(err)
			}
			err = f.serv.CreateMatching(matchCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "Multiple":
			multipleCard := entities.Multiple{}
			err := json.Unmarshal(file, &multipleCard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateMultiple(multipleCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "TrueFalse":
			tfCard := entities.TrueFalse{}
			err := json.Unmarshal(file, &tfCard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateTrueFalse(tfCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")

		case "Info":
			infoCard := entities.Info{}
			err := json.Unmarshal(file, &infoCard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateInfo(infoCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")

		case "QnA":
			QnACard := entities.QnA{}
			err := json.Unmarshal(file, &QnACard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateQnA(QnACard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		default:
			http.Error(w, "invalid type", http.StatusUnprocessableEntity)
			return

		}

	}
}

func (f FlashCardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	allCards, err := f.serv.GetAll()
	if err != nil {
		errorHandlers(w, err)
		return

	}

	allCardsBytes, err := json.MarshalIndent(allCards, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Write(allCardsBytes)

	w.Header().Set("Content-Type", "application/json")

}

func (f FlashCardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"]

	card, err := f.serv.GetById(cardId)
	if err != nil {
		errorHandlers(w, err)
		return
	}

	cardBytes, err := json.MarshalIndent(card, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cardBytes)

}

func (f FlashCardHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"] // get Id by endpoints

	var cardMap map[string]interface{}

	file, err := ioutil.ReadAll(r.Body) // unmarshal to get the type
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(file, &cardMap) //unmarshal into map[string]interface{}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if cardType, ok := cardMap["Type"]; ok { //check type in map "if type exist" and get the value which is in cardtype
		switch cardType { // switching through cardtype and if matches then we unmarshal pass in id and updated info
		case "Matching":
			matchCard := entities.Matching{}
			err := json.Unmarshal(file, &matchCard)
			if err != nil {
				log.Fatalln(err)
			}
			err = f.serv.UpdateMatchingById(cardId, matchCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")
		case "Multiple":
			multipleCard := entities.Multiple{}
			err := json.Unmarshal(file, &multipleCard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateMultipleById(cardId, multipleCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")
		case "TrueFalse":
			tfCard := entities.TrueFalse{}
			err := json.Unmarshal(file, &tfCard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateTrueFalseById(cardId, tfCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")

		case "Info":
			infoCard := entities.Info{}
			err := json.Unmarshal(file, &infoCard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateInfoById(cardId, infoCard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")

		case "QnA":
			QnACard := entities.QnA{}
			err := json.Unmarshal(file, &QnACard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateQnAById(cardId, QnACard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")
		default:
			http.Error(w, "invalid type", http.StatusBadRequest)
			return
		}

	}

}

func (f FlashCardHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := f.serv.DeleteById(id)

	if err != nil {
		errorHandlers(w, err) 
		return
	}


	w.Write([]byte("card deleted"))
	w.Header().Set("Content-Type", "application/json")
	

}

