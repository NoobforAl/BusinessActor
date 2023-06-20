package action

import (
	"context"

	. "github.com/NoobforAl/BusinessActor/src/errors"

	"github.com/NoobforAl/BusinessActor/src/contract"
	"github.com/NoobforAl/BusinessActor/src/entity"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Actor struct {
	stor contract.Stor
}

func NewBaActor(s contract.Stor) Actor {
	return Actor{stor: s}
}

func (ac Actor) Pars(
	c *gin.Context,
) (entity.BusinessActor, error) {
	var ba entity.BusinessActor
	return ba, c.BindJSON(&ba)
}

func (ac Actor) Find(
	c context.Context,
	id string,
) (entity.BusinessActor, error) {

	var ba entity.BusinessActor
	ob, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ba, err
	}

	return ac.stor.GetOneBusinessActor(c,
		bson.M{"_id": ob})
}

func (ac Actor) GetMany(
	c context.Context,
	page, size int64,
) ([]entity.BusinessActor, error) {

	length, err := ac.stor.CountBusinessActor(c, bson.D{})
	if err != nil {
		return nil, err
	}

	if page < 1 || size < 1 {
		return nil, ErrNotFoundPage
	}

	MaxPage := (length / (size))
	if page > MaxPage {
		return nil, ErrNotFoundPage
	}

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * size)
	findOptions.SetLimit(page * size)

	return ac.stor.GetManyBusinessActor(c,
		bson.M{}, findOptions)
}

func (ac Actor) Create(
	c context.Context,
	ba entity.BusinessActor,
) error {
	return ac.stor.InsertBusinessActor(c, ba)
}

func (ac Actor) Update(
	c context.Context,
	ba entity.BusinessActor,
	id string,
) error {

	ob, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	query := bson.M{"_id": ob}
	return ac.stor.UpdateBusinessActor(c, query, ba)
}

func (ac Actor) Delete(
	c context.Context,
	id string,
) error {

	ob, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	query := bson.M{"_id": ob}
	return ac.stor.DeleteBusinessActor(c, query)
}
