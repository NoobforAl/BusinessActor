package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BusinessActor struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Series_reference string `bson:"series_reference"`

	Period     time.Time `bson:"period"`
	Data_value float64   `bson:"data_value"`

	Suppressed bool `bson:"suppressed"`

	STATUS    string `bson:"status"`
	UNITS     string `bson:"units"`
	Magnitude int    `bson:"magnitude"`

	Subject string `bson:"subject"`
	Group   string `bson:"group"`

	Series_title_1 string `bson:"series_title_1"`
	Series_title_2 string `bson:"series_title_2"`
	Series_title_3 string `bson:"series_title_3"`
	Series_title_4 string `bson:"series_title_4"`
	Series_title_5 string `bson:"series_title_5"`
}
