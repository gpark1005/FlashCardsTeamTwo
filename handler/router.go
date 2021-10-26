package handler

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(fc FlashcardHandler) *mux.Router{
  r := mux.NewRouter()

  r.HandleFunc("/flashcard/golang/{Type}", fc.CreateCard).Methods("POST")
  r.HandleFunc("/flashcard/golang", fc.GetAll).Methods("GET")
  


  return r 

}
