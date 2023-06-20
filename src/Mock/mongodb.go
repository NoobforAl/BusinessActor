package mock

import (
	"context"

	"github.com/NoobforAl/BusinessActor/src/db"
	"github.com/NoobforAl/BusinessActor/src/entity"
	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mock struct {
	BaCol mongoifc.Collection
}

func NewMock(cl mongoifc.Collection) Mock {
	return Mock{BaCol: cl}
}

func (s Mock) CountBusinessActor(
	c context.Context, filter any,
) (int64, error) {
	return s.BaCol.CountDocuments(c, filter)
}

func (s Mock) GetOneBusinessActor(
	c context.Context, filter any,
) (entity.BusinessActor, error) {
	var businessActor db.BusinessActor
	err := s.BaCol.FindOne(c, filter).Decode(&businessActor)
	if err != nil {
		return entity.BusinessActor{}, err
	}

	return bindModelBaToEntity(businessActor), nil
}

func (s Mock) GetManyBusinessActor(
	c context.Context,
	filter any,
	opts ...*options.FindOptions,
) ([]entity.BusinessActor, error) {

	cur, err := s.BaCol.Find(c, filter, opts...)
	if err != nil {
		return nil, err
	}

	var businessActors []db.BusinessActor
	if err = cur.All(c, &businessActors); err != nil {
		return nil, err
	}

	baEntity := make([]entity.BusinessActor, len(businessActors))
	for i, v := range businessActors {
		baEntity[i] = bindModelBaToEntity(v)
	}

	return baEntity, nil
}

func (s Mock) InsertBusinessActor(
	c context.Context, d entity.BusinessActor,
) error {
	ev, err := bindEntityToModelBA(d)
	if err != nil {
		return err
	}
	_, err = s.BaCol.InsertOne(c, ev)
	return err
}

func (s Mock) InsertManyBusinessActor(
	c context.Context,
	d []entity.BusinessActor,
) error {
	var err error
	for _, v := range d {
		err = s.InsertBusinessActor(c, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Mock) UpdateBusinessActor(
	c context.Context,
	filter any,
	update entity.BusinessActor,
) error {

	va, err := bindEntityToModelBA(update)
	if err != nil {
		return err
	}
	_, err = s.BaCol.UpdateOne(c, filter, bson.M{"$set": va})
	return err
}

func (s Mock) UpdateManyBusinessActor(
	c context.Context,
	filter any,
	update entity.BusinessActor,
) error {

	va, err := bindEntityToModelBA(update)
	if err != nil {
		return err
	}

	_, err = s.BaCol.UpdateMany(c, filter, va)
	return err
}

func (s Mock) DeleteBusinessActor(
	c context.Context, filter any,
) error {

	_, err := s.GetOneBusinessActor(c, filter)
	if err != nil {
		return err
	}
	_, err = s.BaCol.DeleteOne(c, filter)
	return err
}

func (s Mock) DeleteManyBusinessActor(
	c context.Context, filter any,
) error {
	_, err := s.BaCol.DeleteMany(c, filter)
	return err
}

func bindModelBaToEntity(
	ba db.BusinessActor,
) entity.BusinessActor {
	return entity.BusinessActor{
		Id:               ba.ID.Hex(),
		Series_reference: ba.Series_reference,

		Period:     ba.Period,
		Data_value: ba.Data_value,

		Suppressed: ba.Suppressed,

		STATUS:    ba.STATUS,
		UNITS:     ba.UNITS,
		Magnitude: ba.Magnitude,

		Subject: ba.Subject,
		Group:   ba.Group,

		Series_title_1: ba.Series_title_1,
		Series_title_2: ba.Series_title_2,
		Series_title_3: ba.Series_title_3,
		Series_title_4: ba.Series_title_4,
		Series_title_5: ba.Series_title_5,
	}
}

func bindEntityToModelBA(
	ba entity.BusinessActor,
) (db.BusinessActor, error) {
	id, err := primitive.ObjectIDFromHex(ba.Id)
	if err != nil && ba.Id != "" {
		return db.BusinessActor{}, err
	}

	return db.BusinessActor{
		ID:               id,
		Series_reference: ba.Series_reference,

		Period:     ba.Period,
		Data_value: ba.Data_value,

		Suppressed: ba.Suppressed,

		STATUS:    ba.STATUS,
		UNITS:     ba.UNITS,
		Magnitude: ba.Magnitude,

		Subject: ba.Subject,
		Group:   ba.Group,

		Series_title_1: ba.Series_title_1,
		Series_title_2: ba.Series_title_2,
		Series_title_3: ba.Series_title_3,
		Series_title_4: ba.Series_title_4,
		Series_title_5: ba.Series_title_5,
	}, nil
}
