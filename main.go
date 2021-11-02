package main

import (
	"fmt"
	"net/http"

	"github.com/gpark1005/FlashCardsTeamTwo/Repo"
	"github.com/gpark1005/FlashCardsTeamTwo/handler"
	"github.com/gpark1005/FlashCardsTeamTwo/service"

	//"net/http"
	"log"
	"path/filepath"
	//"fmt"
	//"github.com/gpark1005/FlashCardsTeamTwo/entities"
	//"github.com/gpark1005/FlashCardsTeamTwo/handler"
)

func main() {

	filename := "structFlashCard.json"

	ext := filepath.Ext(filename)

	if ext != ".json" {
		log.Fatalln("Invalid file")
	}

	Repo := Repo.NewFlashcardRepo(filename)
	serv := service.NewFlashcardService(Repo)
	handle := handler.NewFlashCardHandler(serv)

	router := handler.ConfigureRouter(handle)

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	fmt.Println("Successfully running server 8080")

	log.Fatal(server.ListenAndServe())

}
