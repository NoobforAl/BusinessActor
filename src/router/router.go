package router

import (
	"github.com/NoobforAl/BusinessActor/src/contract"
	"github.com/NoobforAl/BusinessActor/src/controller"
	"github.com/gin-gonic/gin"
)

// business actor router
func BaRouter(api *gin.RouterGroup, stor contract.Stor) {
	api.GET("/get/:id", controller.FindBusinessActor(stor))
	api.GET("/getMany", controller.GetManyBusinessActor(stor))

	api.POST("/create", controller.CreateBusinessActor(stor))

	api.PUT("/update/:id", controller.UpdateBusinessActor(stor))

	api.DELETE("/delete/:id", controller.DeleteBusinessActor(stor))
}
