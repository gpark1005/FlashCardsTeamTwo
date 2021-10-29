package handler

import (
	"encoding/json"
	"io/ioutil"
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
	GetById(string) (interface{}, error)
	UpdateMatchingById(string, entities.Matching) error
	UpdateMultipleById(string, entities.Multiple) error
	UpdateTrueFalseById(string, entities.TrueFalse) error
	UpdateInfoById(string, entities.Info) error
	UpdateQandAById(string, entities.QandA) error
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
	var c map[string]interface{}
	file, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(file, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if cardtype, ok := c["Type"]; ok {
		switch cardtype {
		case "Matching":
			matchcard := entities.Matching{}
			err := json.Unmarshal(file, &matchcard)
			if err != nil {
				log.Fatalln(err)
			}
			err = f.serv.CreateMatching(matchcard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "Multiple":
			multiplecard := entities.Multiple{}
			err := json.Unmarshal(file, &multiplecard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateMultiple(multiplecard)
			if err != nil {
				errorHandlers(w, err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")
		case "TrueFalse":
			tfcard := entities.TrueFalse{}
			err := json.Unmarshal(file, &tfcard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateTrueFalse(tfcard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")

		case "Info":
			infocard := entities.Info{}
			err := json.Unmarshal(file, &infocard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateInfo(infocard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card created"))
			w.Header().Set("Content-Type", "application/json")

		case "QandA":
			QandAcard := entities.QandA{}
			err := json.Unmarshal(file, &QandAcard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.CreateQandA(QandAcard)
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

func (f FlashcardHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	allCards, err := f.serv.GetAll()
	if err != nil {
		errorHandlers(w, err)
		return

	}

	allCardsBytes, err := json.MarshalIndent(allCards, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	w.Write(allCardsBytes)

	w.Header().Set("Content-Type", "application/json")

}

func (f FlashcardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"]

	card, err := f.serv.GetById(cardId)
	if err != nil {
		errorHandlers(w, err)
		return
	}

	cardBytes, err := json.MarshalIndent(card, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cardBytes)

}

func (f FlashcardHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["Id"] // get Id by endpoints

	var c map[string]interface{}

	file, err := ioutil.ReadAll(r.Body) // unmarshal to get the type
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(file, &c) //unmarshal into map[string]interface{}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if cardtype, ok := c["Type"]; ok { //check type in map "if type exist" and get the value which is in cardtype
		switch cardtype { // switching through cardtype and if matches then we unmarshal pass in id and updated info
		case "Matching":
			matchcard := entities.Matching{}
			err := json.Unmarshal(file, &matchcard)
			if err != nil {
				log.Fatalln(err)
			}
			err = f.serv.UpdateMatchingById(cardId, matchcard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")
		case "Multiple":
			multiplecard := entities.Multiple{}
			err := json.Unmarshal(file, &multiplecard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateMultipleById(cardId, multiplecard)
			if err != nil {
				errorHandlers(w, err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")
		case "TrueFalse":
			tfcard := entities.TrueFalse{}
			err := json.Unmarshal(file, &tfcard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateTrueFalseById(cardId, tfcard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")

		case "Info":
			infocard := entities.Info{}
			err := json.Unmarshal(file, &infocard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateInfoById(cardId, infocard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")

		case "QandA":
			QandAcard := entities.QandA{}
			err := json.Unmarshal(file, &QandAcard)
			if err != nil {
				log.Fatalln(err)
			}

			err = f.serv.UpdateQandAById(cardId, QandAcard)
			if err != nil {
				errorHandlers(w, err)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Card updated"))
			w.Header().Set("Content-Type", "application/json")
		default:
			http.Error(w, "invalid type", http.StatusUnprocessableEntity)
			return
		}

	}

}
