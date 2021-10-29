package handler

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(fc FlashcardHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcard", fc.CreateCard).Methods("POST")
	r.HandleFunc("/flashcard", fc.GetAll).Methods("GET")
	r.HandleFunc("/flashcard/{Id}", fc.GetById).Methods("GET")
	r.HandleFunc("/flashcard/{Id}", fc.UpdateById).Methods("PUT")

	return r

}
