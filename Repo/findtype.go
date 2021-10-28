package Repo


import 
"github.com/gpark1005/FlashCardsTeamTwo/entities"


func IdCheck(id string, deck entities.FlashCardStruct) (interface{}, error) {
	


	if len(deck.Matching) > 0 {
		for _, v := range deck.Matching {
			if v.Id == id {
				returnDeck = v
				return returnDeck, nil
			}

		}
	}


}