package entity

import (
	"context"

	"github.com/NoobforAl/BusinessActor/src/contract"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (ba *BusinessActor) Pars(c *gin.Context) error {
	return c.BindJSON(ba)
}

func (ba *BusinessActor) Find(s contract.Stor, c context.Context, id string) error {
	ob, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	data, err := s.GetOneBusinessActor(c, bson.M{"_id": ob})
	if err != nil {
		return err
	}

	ba.Id = data.ID.Hex()
	ba.Series_reference = data.Series_reference

	ba.Period = data.Period
	ba.Data_value = data.Data_value
	ba.Suppressed = data.Suppressed

	ba.STATUS = data.STATUS
	ba.UNITS = data.UNITS

	ba.Magnitude = data.Magnitude
	ba.Subject = data.Subject
	ba.Group = data.Group

	ba.Series_title_1 = data.Series_title_1
	ba.Series_title_2 = data.Series_title_2
	ba.Series_title_3 = data.Series_title_3
	ba.Series_title_4 = data.Series_title_4
	ba.Series_title_5 = data.Series_title_5

	return nil
}

func (ba *BusinessActor) GetMany(s contract.Stor, c context.Context, page, size int64) ([]BusinessActor, error) {
	length, err := s.CountBusinessActor(c, bson.D{})
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

	businessActors, err := s.GetManyBusinessActor(c, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	var entityBA []BusinessActor
	for _, data := range businessActors {
		entityBA = append(entityBA, BusinessActor{
			Id:               data.ID.Hex(),
			Series_reference: data.Series_reference,

			Period:     data.Period,
			Data_value: data.Data_value,
			Suppressed: data.Suppressed,

			STATUS: data.STATUS,
			UNITS:  data.UNITS,

			Magnitude: data.Magnitude,
			Subject:   data.Subject,
			Group:     data.Group,

			Series_title_1: data.Series_title_1,
			Series_title_2: data.Series_title_2,
			Series_title_3: data.Series_title_3,
			Series_title_4: data.Series_title_4,
			Series_title_5: data.Series_title_5,
		})
	}
	return entityBA, nil
}

func (ba *BusinessActor) Create(s contract.Stor, c context.Context) error {
	newBa := s.NewBA()
	newBa.Series_reference = ba.Series_reference

	newBa.ID = primitive.NewObjectID()
	ba.Id = newBa.ID.Hex()

	newBa.Period = ba.Period
	newBa.Data_value = ba.Data_value
	newBa.Suppressed = ba.Suppressed

	newBa.STATUS = ba.STATUS
	newBa.UNITS = ba.UNITS

	newBa.Magnitude = ba.Magnitude
	newBa.Subject = ba.Subject
	newBa.Group = ba.Group

	newBa.Series_title_1 = ba.Series_title_1
	newBa.Series_title_2 = ba.Series_title_2
	newBa.Series_title_3 = ba.Series_title_3
	newBa.Series_title_4 = ba.Series_title_4
	newBa.Series_title_5 = ba.Series_title_5
	return s.InsertBusinessActor(c, newBa)
}

func (ba *BusinessActor) Update(s contract.Stor, c context.Context, id string) error {
	ob, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	newBa := s.NewBA()
	query := bson.M{"_id": ob}

	_, err = s.GetOneBusinessActor(c, query)
	if err != nil {
		return err
	}

	newBa.Series_reference = ba.Series_reference

	newBa.Period = ba.Period
	newBa.Data_value = ba.Data_value
	newBa.Suppressed = ba.Suppressed

	newBa.STATUS = ba.STATUS
	newBa.UNITS = ba.UNITS

	newBa.Magnitude = ba.Magnitude
	newBa.Subject = ba.Subject
	newBa.Group = ba.Group

	newBa.Series_title_1 = ba.Series_title_1
	newBa.Series_title_2 = ba.Series_title_2
	newBa.Series_title_3 = ba.Series_title_3
	newBa.Series_title_4 = ba.Series_title_4
	newBa.Series_title_5 = ba.Series_title_5
	return s.UpdateBusinessActor(c, query, bson.M{"$set": newBa})
}

func (ba *BusinessActor) Delete(s contract.Stor, c context.Context, id string) error {
	ob, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.DeleteBusinessActor(c, bson.M{"_id": ob})
}
