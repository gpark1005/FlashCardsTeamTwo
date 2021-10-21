package main

import (
	//"net/http"
	"path/filepath"
	"log"
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

	


	//router := handler.NewRouter(handler)



	/*server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}
	


	fmt.Println("Successfully running server 8080")

	log.Fatal(server.ListenAndServe())

	*/
}