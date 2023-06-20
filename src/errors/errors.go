package errorApp

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFoundPage = errors.New("not found page")
	ErrBadRequest   = errors.New("bad request ")

	ErrNotFound = mongo.ErrNoDocuments
)
