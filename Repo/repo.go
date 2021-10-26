package Repo

import "github.com/gpark1005/FlashCardsTeamTwo/entities"



type flashcardRepo struct {
	filename string 
}

func NewFlashcardRepo(filename string) flashcardRepo {
	return flashcardRepo{
		filename: filename,
	}
}

func (r flashcardRepo) CreatMatching (card entities.Matching) error {
	
}


