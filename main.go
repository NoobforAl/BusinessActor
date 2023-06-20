package main

import (
	"net/http"

	loadCsv "github.com/NoobforAl/BusinessActor/src/businessActorCsv"
	"github.com/NoobforAl/BusinessActor/src/db"
	env "github.com/NoobforAl/BusinessActor/src/loadEnv"
	"github.com/NoobforAl/BusinessActor/src/logger"
	"github.com/NoobforAl/BusinessActor/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	stor := db.GetDb()
	loadCsv.InitData(stor, env.GetCsvPath())

	api := r.Group("/api")
	router.BaRouter(api, stor)

	if err := http.ListenAndServe(env.GetAddrListen(), r); err != nil {
		logger.Log.Fatal(err)
	}
}
