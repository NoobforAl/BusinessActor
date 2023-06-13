package main

import (
	"net/http"

	loadCsv "github.com/NoobforAl/BusinessActor/src/businessActorCsv"
	"github.com/NoobforAl/BusinessActor/src/controller"
	"github.com/NoobforAl/BusinessActor/src/db"
	env "github.com/NoobforAl/BusinessActor/src/loadEnv"
	"github.com/NoobforAl/BusinessActor/src/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	stor := db.GetDb()
	loadCsv.InitData(stor, env.GetCsvPath())

	api := r.Group("/api")
	{
		api.GET("/get/:id", controller.FindBusinessActor(stor))
		api.GET("/getMany", controller.GetManyBusinessActor(stor))

		api.POST("/create", controller.CreateBusinessActor(stor))

		api.PUT("/update/:id", controller.UpdateBusinessActor(stor))

		api.DELETE("/delete/:id", controller.DeleteBusinessActor(stor))
	}

	if err := http.ListenAndServe(env.GetAddrListen(), r); err != nil {
		logger.Log.Fatal(err)
	}
}
