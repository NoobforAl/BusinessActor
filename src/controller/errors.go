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
	case errors.Is(primitive.ErrInvalidHex, err) ||
		errors.Is(ErrBadRequest, err):
		status = http.StatusBadRequest

	case errors.Is(entity.ErrNotFoundPage, err):
		status = http.StatusNotFound
	}
	c.JSON(status, gin.H{"detail": err.Error()})
}
