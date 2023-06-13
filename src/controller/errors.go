package controller

import (
	"errors"
	"net/http"

	"github.com/NoobforAl/BusinessActor/src/entity"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrBadRequest = errors.New("bad request ")
)

func handleErr(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, primitive.ErrInvalidHex) ||
		errors.Is(err, ErrBadRequest):
		status = http.StatusBadRequest

	case errors.Is(err, entity.ErrNotFoundPage):
		status = http.StatusNotFound
	}
	c.JSON(status, gin.H{"detail": err.Error()})
}
