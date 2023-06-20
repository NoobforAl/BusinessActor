package tests

import (
	"context"
	"os"
	"testing"

	mock "github.com/NoobforAl/BusinessActor/src/Mock"
	loadCsv "github.com/NoobforAl/BusinessActor/src/businessActorCsv"
	"github.com/NoobforAl/BusinessActor/src/logger"
	"github.com/NoobforAl/BusinessActor/src/router"
	"github.com/gin-gonic/gin"
	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mo mock.Mock
var r *gin.Engine

func TestMain(m *testing.M) {
	col, err := setUpMongoDB()
	if err != nil {
		logger.Log.Fatal(err)
	}

	setUpGinApi()

	defer func() {
		if err = col.Drop(context.Background()); err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

func setUpMongoDB() (mongoifc.Collection, error) {
	var col mongoifc.Collection

	// test with your database
	dsn := "mongodb://mongoadmin:dasfa4523da3214esad@127.0.0.1:27017/admin"
	stor, err := mongoifc.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		return col, err
	}

	col = stor.Database("tests").Collection("records")
	if err = stor.Connect(context.Background()); err != nil {
		return col, err
	}

	if err = stor.Ping(context.TODO(), readpref.Primary()); err != nil {
		return col, err
	}

	pathFile := "../businessActorCsv/" +
		"business-financial-data-mar-2022-quarter-csv.csv"

	mo = mock.NewMock(col)
	loadCsv.InitData(mo, pathFile)
	return col, nil
}

func setUpGinApi() {
	r = gin.Default()
	api := r.Group("/api")
	router.BaRouter(api, mo)
}
