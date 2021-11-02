package Repo

import "github.com/gpark1005/FlashCardsTeamTwo/entities"

func IdCheck(id string, deck entities.FlashCardStruct) (interface{}, error) {
	var cardFound interface{}
	if len(deck.Matching) > 0 {
		for _, v := range deck.Matching {
			if v.Id == id {
				cardFound  = v
				return cardFound , nil
			}

		}
	}

	if len(deck.TrueFalse) > 0 {
		for _, v := range deck.TrueFalse {
			if v.Id == id {
				cardFound  = v
				return cardFound , nil
			}

		}
	}

	if len(deck.Info) > 0 {
		for _, v := range deck.Info {
			if v.Id == id {
				cardFound = v
				return cardFound , nil
			}

		}
	}

	if len(deck.Multiple) > 0 {
		for _, v := range deck.Multiple {
			if v.Id == id {
				cardFound  = v
				return cardFound , nil
			}

		}
	}

	if len(deck.QandA) > 0 {
		for _, v := range deck.QandA {
			if v.Id == id {
				cardFound  = v
				return cardFound , nil
			}

		}
	}

	return cardFound , NotFound

}

func DeleteCheck (id string, deck entities.FlashCardStruct) (int, string, error) {

	

	if len(deck.Matching) > 0 {
		for i, v := range deck.Matching {
			if v.Id == id {
				
				return i, v.Type, nil
			}

		}
	}

	if len(deck.TrueFalse) > 0 {
		for i, v := range deck.TrueFalse {
			if v.Id == id {
				
				return i, v.Type, nil
			}

		}
	}

	if len(deck.Info) > 0 {
		for i, v := range deck.Info {
			if v.Id == id {
				return i, v.Type, nil
			}

		}
	}

	if len(deck.Multiple) > 0 {
		for i, v := range deck.Multiple {
			if v.Id == id {
				return i, v.Type, nil
			}

		}
	}

	if len(deck.QandA) > 0 {
		for i, v := range deck.QandA {
			if v.Id == id {
				return i, v.Type, nil
			}

		}
	}



	return -1, "Not Found", NotFound

}

