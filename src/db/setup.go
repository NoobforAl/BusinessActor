package db

import (
	"context"
	"sync"

	env "github.com/NoobforAl/BusinessActor/src/loadEnv"
	"github.com/NoobforAl/BusinessActor/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type database struct {
	// DB client
	cl *mongo.Client

	// BusinessActor database
	BaDB *mongo.Database

	// BusinessActor collection
	BaCol *mongo.Collection
}

var onc sync.Once
var stor database

func GetDb() database {
	onc.Do(func() {
		var err error
		dsn := env.GetDsn()

		stor.cl, err = mongo.NewClient(options.Client().ApplyURI(dsn))
		if err != nil {
			logger.Log.Fatal(err)
		}

		if err = stor.cl.Connect(context.Background()); err != nil {
			logger.Log.Fatal(err)
		}

		if err = stor.cl.Ping(context.TODO(), readpref.Primary()); err != nil {
			logger.Log.Fatal(err)
		}

		stor.BaDB = stor.cl.Database("BusinessActor")
		stor.BaCol = stor.BaDB.Collection("records")
	})

	return stor
}
