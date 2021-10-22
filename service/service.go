package service


type IFlashcardRepo interface {
	Create(interface{}) error
}

type FlashcardService struct {
	Repo IFlashcardRepo
}

func NewFlashcardService(f IFlashcardRepo) FlashcardService {
	return FlashcardService{
		Repo : f, 
	}
}



