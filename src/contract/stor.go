package contract

import (
	"context"

	"github.com/NoobforAl/BusinessActor/src/db"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Stor interface {
	NewBA() *db.BusinessActor

	CountBusinessActor(c context.Context, filter any) (int64, error)

	GetOneBusinessActor(c context.Context, filter any) (db.BusinessActor, error)
	GetManyBusinessActor(c context.Context, filter any, opts ...*options.FindOptions) ([]db.BusinessActor, error)

	InsertBusinessActor(c context.Context, d any) error
	InsertManyBusinessActor(c context.Context, d []db.BusinessActor) error

	UpdateBusinessActor(c context.Context, filter any, update any) error
	UpdateManyBusinessActor(c context.Context, filter any, update any) error

	DeleteBusinessActor(c context.Context, filter any) error
	DeleteManyBusinessActor(c context.Context, filter any) error
}
