package mock

import (
	"context"

	"github.com/NoobforAl/BusinessActor/src/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mock struct {
	cl *mongo.Collection
}

func NewMock(cl *mongo.Collection) Mock {
	return Mock{cl: cl}
}

func (m Mock) NewBA() *db.BusinessActor {
	return &db.BusinessActor{}
}

func (m Mock) CountBusinessActor(c context.Context, filter any) (int64, error) {
	return m.cl.CountDocuments(c, filter)
}

func (m Mock) GetOneBusinessActor(c context.Context, filter any) (db.BusinessActor, error) {
	var businessActor db.BusinessActor
	err := m.cl.FindOne(c, filter).Decode(&businessActor)
	if err != nil {
		return businessActor, err
	}

	return businessActor, nil
}

func (m Mock) GetManyBusinessActor(
	c context.Context,
	filter any,
	opts ...*options.FindOptions) ([]db.BusinessActor, error) {

	cur, err := m.cl.Find(c, filter, opts...)
	if err != nil {
		return nil, err
	}

	var businessActors []db.BusinessActor
	if err = cur.All(c, &businessActors); err != nil {
		return nil, err
	}

	return businessActors, nil
}

func (m Mock) InsertBusinessActor(c context.Context, d any) error {
	_, err := m.cl.InsertOne(c, d)
	return err
}

func (m Mock) InsertManyBusinessActor(c context.Context, d []db.BusinessActor) error {
	var err error
	for _, v := range d {
		err = m.InsertBusinessActor(c, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Mock) UpdateBusinessActor(c context.Context, filter any, update any) error {
	_, err := m.cl.UpdateOne(c, filter, update)
	return err
}

func (m Mock) UpdateManyBusinessActor(c context.Context, filter any, update any) error {
	_, err := m.cl.UpdateMany(c, filter, update)
	return err
}

func (m Mock) DeleteBusinessActor(c context.Context, filter any) error {
	_, err := m.cl.DeleteOne(c, filter)
	return err
}

func (m Mock) DeleteManyBusinessActor(c context.Context, filter any) error {
	_, err := m.cl.DeleteMany(c, filter)
	return err
}
