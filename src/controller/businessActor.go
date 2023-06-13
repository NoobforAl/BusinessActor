package controller

import (
	"errors"
	"net/http"

	"github.com/NoobforAl/BusinessActor/src/contract"
	"github.com/NoobforAl/BusinessActor/src/entity"
	"github.com/gin-gonic/gin"
)

func FindBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ba := entity.BusinessActor{}
		if err := ba.Find(stor, c, id); err != nil {
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

		ba := entity.BusinessActor{}

		businessActors, err := ba.GetMany(stor, c, page, size)
		if err != nil {
			handleErr(c, errors.Join(ErrBadRequest, err))
			return
		}
		c.JSON(http.StatusOK, businessActors)
	}
}

func CreateBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		ba := entity.BusinessActor{}
		var err error

		if err = ba.Pars(c); err != nil {
			handleErr(c, errors.Join(ErrBadRequest, err))
			return
		}

		if err = ba.Create(stor, c); err != nil {
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
		ba := entity.BusinessActor{}
		id := c.Param("id")
		var err error

		if err = ba.Pars(c); err != nil {
			handleErr(c, errors.Join(ErrBadRequest, err))
			return
		}

		if err = ba.Update(stor, c, id); err != nil {
			handleErr(c, err)
			return
		}

		c.JSON(http.StatusOK, ba)
	}
}

func DeleteBusinessActor(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		ba := entity.BusinessActor{}
		id := c.Param("id")

		if err := ba.Delete(stor, c, id); err != nil {
			handleErr(c, err)
			return
		}

		c.JSON(http.StatusOK, ba)
	}
}
