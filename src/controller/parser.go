package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getQueryInt64(c *gin.Context, key string) (int64, error) {
	val := c.Query(key)
	return strconv.ParseInt(val, 10, 64)
}
