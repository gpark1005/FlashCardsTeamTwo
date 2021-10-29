package service

import (
	"errors"
)

var (
	BadRequest error = errors.New("bad request")
	InvalidId  error = errors.New("invalid id")
	InvalidCategory error =  errors.New("category not found")
	InvalidType error = errors.New("type not found")
	InvalidQuestion error = errors.New("the question field must cannot be empty, please enter a question")
	InvalidAnswer error = errors.New("the answer field cannot be empty, please enter an answer")
	AnswerNotAccepted error = errors.New("the answer must be either true or false. please try again")
	InvalidData error = errors.New("this field cannot be empty, please enter the details of this card")
	QuestionsInaccurate error = errors.New("requires more than 1 question to be asked. please try again")
	FaultyAnswers error = errors.New("requires more than 1 answer to be given. please try again")
	ErroneousOptions error = errors.New("requires more than 1 option, please try again")
	UnacceptableFormat error = errors.New("all the options, question, and answer field must contain the same amount info. please review and correct")
	AnswerNotFound error = errors.New("answer is not available in the choices, please review and correct")
	MisMatch error = errors.New("the keys and values do not add up. please review and try again")

)
