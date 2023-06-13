package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s database) NewBA() *BusinessActor {
	return &BusinessActor{}
}

func (s database) CountBusinessActor(c context.Context, filter any) (int64, error) {
	return s.db.CountDocuments(c, filter)
}

func (s database) GetOneBusinessActor(c context.Context, filter any) (BusinessActor, error) {
	var businessActor BusinessActor
	err := s.db.FindOne(c, filter).Decode(&businessActor)
	if err != nil {
		return businessActor, err
	}

	return businessActor, nil
}

func (s database) GetManyBusinessActor(
	c context.Context,
	filter any,
	opts ...*options.FindOptions) ([]BusinessActor, error) {

	cur, err := s.db.Find(c, filter, opts...)
	if err != nil {
		return nil, err
	}

	var businessActors []BusinessActor
	if err = cur.All(c, &businessActors); err != nil {
		return nil, err
	}

	return businessActors, nil
}

func (s database) InsertBusinessActor(c context.Context, d any) error {
	_, err := s.db.InsertOne(c, d)
	return err
}

func (s database) InsertManyBusinessActor(c context.Context, d []BusinessActor) error {
	var err error
	for _, v := range d {
		err = s.InsertBusinessActor(c, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s database) UpdateBusinessActor(c context.Context, filter any, update any) error {
	_, err := s.db.UpdateOne(c, filter, update)
	return err
}

func (s database) UpdateManyBusinessActor(c context.Context, filter any, update any) error {
	_, err := s.db.UpdateMany(c, filter, update)
	return err
}

func (s database) DeleteBusinessActor(c context.Context, filter any) error {
	_, err := s.db.DeleteOne(c, filter)
	return err
}

func (s database) DeleteManyBusinessActor(c context.Context, filter any) error {
	_, err := s.db.DeleteMany(c, filter)
	return err
}
