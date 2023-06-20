package contract

import (
	"context"

	"github.com/NoobforAl/BusinessActor/src/entity"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BusinessActor interface {
	CountBusinessActor(
		c context.Context, filter any,
	) (int64, error)

	GetOneBusinessActor(
		c context.Context, filter any,
	) (entity.BusinessActor, error)

	GetManyBusinessActor(
		c context.Context,
		filter any,
		opts ...*options.FindOptions,
	) ([]entity.BusinessActor, error)

	InsertBusinessActor(
		c context.Context, d entity.BusinessActor,
	) error

	InsertManyBusinessActor(
		c context.Context,
		d []entity.BusinessActor,
	) error

	UpdateBusinessActor(
		c context.Context,
		filter any,
		update entity.BusinessActor,
	) error

	UpdateManyBusinessActor(
		c context.Context,
		filter any,
		update entity.BusinessActor,
	) error

	DeleteBusinessActor(
		c context.Context, filter any,
	) error

	DeleteManyBusinessActor(
		c context.Context, filter any,
	) error
}

type Stor interface {
	BusinessActor
}
