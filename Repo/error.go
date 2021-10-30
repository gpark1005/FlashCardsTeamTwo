package Repo

import (
	"errors"
)

var (
	NotFound    error = errors.New("card not found")
	ServerError error = errors.New("internal server error")
	
)
