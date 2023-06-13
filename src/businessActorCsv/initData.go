package loadCsv

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/NoobforAl/BusinessActor/src/contract"
	"github.com/NoobforAl/BusinessActor/src/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var onc sync.Once
var wg sync.WaitGroup

func formatTime(s string) (time.Time, error) {
	const layout = "2006.01"
	return time.Parse(layout, s)
}

func convFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func convInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func combineInsert(c context.Context, s contract.Stor, ch <-chan []string) {
	var err error
	for v := range ch {
		ba := s.NewBA()
		ba.ID = primitive.NewObjectID()
		ba.Series_reference = v[0]

		ba.Period, err = formatTime(v[1])
		if err != nil {
			logger.Log.Fatal(err)
		}

		ba.Suppressed = v[3] == "Y"

		if !ba.Suppressed && v[2] != "" {
			ba.Data_value, err = convFloat64(v[2])
			if err != nil {
				logger.Log.Fatal(err)
			}
		}

		ba.STATUS = v[4]
		ba.UNITS = v[5]

		ba.Magnitude, err = convInt(v[6])
		if err != nil {
			logger.Log.Fatal(err)
		}

		ba.Subject = v[7]
		ba.Group = v[8]
		ba.Series_title_1 = v[9]
		ba.Series_title_2 = v[10]
		ba.Series_title_3 = v[11]
		ba.Series_title_4 = v[12]
		ba.Series_title_5 = v[13]

		if err = s.InsertBusinessActor(c, ba); err != nil {
			logger.Log.Fatal(err)
		}
	}
	wg.Done()
}

func InitData(s contract.Stor) {
	onc.Do(func() {
		ctx := context.TODO()
		c, err := s.CountBusinessActor(ctx, bson.M{})
		if err != nil {
			panic(err)
		}

		// if have any record in database
		// do nothing
		if c > 0 {
			logger.Log.Printf("In database have %d data. not need add again.", c)
			return
		}

		logger.Log.Println("not found data in database and now insert New!")
		csvPath := "./src/businessActorCsv/business-financial-data-mar-2022-quarter-csv.csv"
		f, err := os.Open(csvPath)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		csvData := csv.NewReader(f)
		ch := make(chan []string)

		wg.Add(4)
		go combineInsert(ctx, s, ch)
		go combineInsert(ctx, s, ch)
		go combineInsert(ctx, s, ch)
		go combineInsert(ctx, s, ch)

		// get first line
		_, _ = csvData.Read()

		for {
			d, err := csvData.Read()
			if err != nil || err == io.EOF {
				break
			}
			ch <- d
		}
		close(ch)
		wg.Wait()
		logger.Log.Println("insert data in database is Done!")
	})
}
