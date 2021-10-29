package handler

import (
	"net/http"

	"github.com/gpark1005/FlashCardsTeamTwo/Repo"
	"github.com/gpark1005/FlashCardsTeamTwo/service"
)

func errorHandlers(w http.ResponseWriter, err error) {
	switch err {
	case service.AnswerNotAccepted:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.BadRequest:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case service.InvalidId:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case service.InvalidCategory:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.InvalidType:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.InvalidAnswer:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.InvalidQuestion:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.InvalidData:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.QuestionsInaccurate:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.FaultyAnswers:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.ErroneousOptions:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.UnacceptableFormat:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.AnswerNotFound:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case service.MisMatch:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	case Repo.ServerError:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case Repo.NotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	}
}
