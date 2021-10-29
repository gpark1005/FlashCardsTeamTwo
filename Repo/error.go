package Repo

import (
	"errors"
)

var (
	BadRequest error = errors.New("bad request")
	InvalidId  error = errors.New("invalid id")
	NotFound   error = errors.New("card not found")
	ServerError error = errors.New("server error")
)
