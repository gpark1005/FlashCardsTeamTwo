package Repo

import (
	"errors"
)

var (
	NotFound    error = errors.New("card not found")
	ServerError error = errors.New("internal server error")
	FileError error = errors.New("file cannot be read")
	StorageError error = errors.New("unable to store data")
	ConversionError error = errors.New("unable to convert data")
)
