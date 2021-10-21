package entities

type FlashCardStruct struct {
	Id            string
	Category      string
	Type          string
	Info          string
	MatchQuestion CardQuestion
	Question      string
	MatchAnswer   CardAnswer
	Answer        string
	Options       CardOption
}

type CardQuestion struct {
	A string
	B string
	C string
	D string
}

type CardAnswer struct {
	A string
	B string
	C string
	D string
}

type CardOption struct {
	One   string
	Two   string
	Three string
	Four  string
}
