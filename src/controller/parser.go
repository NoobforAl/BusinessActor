package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getQueryInt(c *gin.Context, key string) (int, error) {
	val := c.Query(key)
	return strconv.Atoi(val)
}
