package controller

import (
	"errors"
	"net/http"

	"github.com/NoobforAl/BusinessActor/src/action"
	"github.com/NoobforAl/BusinessActor/src/contract"
	"github.com/gin-gonic/gin"

	. "github.com/NoobforAl/BusinessActor/src/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ba, err := action.NewBaActor(stor).Find(c, id)
		if err != nil {
			handleErr(c, err)
			return
		}
		c.JSON(http.StatusOK, ba)
	}
}

func GetManyBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := getQueryInt64(c, "page")
		if err != nil {
			handleErr(c, ErrBadRequest)
			return
		}

		size, err := getQueryInt64(c, "size")
		if err != nil {
			handleErr(c, ErrBadRequest)
			return
		}

		data, err := action.NewBaActor(stor).GetMany(c, page, size)
		if err != nil {
			handleErr(c, err)
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

func CreateBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		ac := action.NewBaActor(stor)
		var err error

		ba, err := ac.Pars(c)
		if err != nil {
			handleErr(c, errors.Join(ErrBadRequest, err))
			return
		}

		if err = ac.Create(c, ba); err != nil {
			handleErr(c, err)
			return
		}

		c.JSON(http.StatusOK, ba)
	}
}

// ! to be carful for run this func
// if you send a empty value
// in database changed value to empty !!
func UpdateBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		ac := action.NewBaActor(stor)
		id := c.Param("id")
		var err error

		ba, err := ac.Pars(c)
		if err != nil {
			handleErr(c, errors.Join(ErrBadRequest, err))
			return
		}

		if err = ac.Update(c, ba, id); err != nil {
			handleErr(c, err)
			return
		}

		c.JSON(http.StatusOK, ba)
	}
}

func DeleteBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		ac := action.NewBaActor(stor)
		id := c.Param("id")

		if err := ac.Delete(c, id); err != nil {
			handleErr(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}

func handleErr(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, primitive.ErrInvalidHex) ||
		errors.Is(err, ErrBadRequest):
		status = http.StatusBadRequest

	case errors.Is(err, ErrNotFoundPage) ||
		errors.Is(err, ErrNotFound):
		status = http.StatusNotFound
	}
	c.JSON(status, gin.H{"detail": err.Error()})
}
