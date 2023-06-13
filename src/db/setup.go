package db

import (
	"context"
	"sync"

	env "github.com/NoobforAl/BusinessActor/loadEnv"
	"github.com/NoobforAl/BusinessActor/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type database struct {
	cl *mongo.Client
	db *mongo.Collection
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

		if err = stor.cl.Ping(context.TODO(), readpref.Primary()); err != nil {
			logger.Log.Fatal(err)
		}

		stor.db = stor.cl.Database("BusinessActor").Collection("records")
	})

	return stor
}
