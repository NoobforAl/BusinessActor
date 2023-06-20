package tests

import (
	"context"
	"testing"
	"time"

	"github.com/NoobforAl/BusinessActor/src/action"
	"github.com/NoobforAl/BusinessActor/src/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCaseFindEntity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ac := action.NewBaActor(mo)
	datas, err := ac.GetMany(ctx, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	dataFind := datas[0]
	ba, err := ac.Find(ctx, dataFind.Id)
	if err != nil {
		t.Errorf("Get error for find value: %v", err)
	}

	if ba.Series_title_2 != dataFind.Series_title_2 {
		t.Errorf("%s != %s", ba.Series_title_2, dataFind.Series_title_2)
	}
}

func TestCASEUpdateEntity(t *testing.T) {
	ctx := context.Background()
	ac := action.NewBaActor(mo)
	datas, err := ac.GetMany(ctx, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	dataUpdate := datas[1]
	ba := entity.BusinessActor{Series_title_5: "TEST"}
	err = ac.Update(ctx, ba, dataUpdate.Id)
	if err != nil {
		t.Errorf("Get error for update value: %v", err)
	}
}

func TestCASECreateEntity(t *testing.T) {
	ctx := context.Background()
	ba := entity.BusinessActor{
		Id:               primitive.NewObjectID().Hex(),
		Series_reference: "BDCQ.SF1AA2CA",
		Period:           time.Now(),
		Data_value:       1116.386,
		Suppressed:       false,
		STATUS:           "F",
		UNITS:            "Dollars",
		Magnitude:        6,
		Subject:          "Business Data Collection - BDC",
		Group:            "Industry by financial variable (NZSIOC Level 2)",
		Series_title_1:   "Sales (operating income)",
		Series_title_2:   "Forestry and Logging",
		Series_title_3:   "Current prices",
		Series_title_4:   "Unadjusted",
		Series_title_5:   "test3343",
	}

	ac := action.NewBaActor(mo)
	err := ac.Create(ctx, ba)
	if err != nil {
		t.Errorf("Get error for create value: %v", err)
	}

	ba, err = ac.Find(ctx, ba.Id)
	if err != nil {
		t.Errorf("Get error for create value: %v", err)
	}
}

func TestCASEDeleteEntity(t *testing.T) {
	ctx := context.Background()
	ac := action.NewBaActor(mo)
	datas, err := ac.GetMany(ctx, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	dataDelete := datas[3]
	err = ac.Delete(ctx, dataDelete.Id)
	if err != nil {
		t.Errorf("Get error for delete value: %v", err)
	}

	err = ac.Delete(ctx, dataDelete.Id)
	if err == nil {
		t.Error("error not equal nil")
	}
}
