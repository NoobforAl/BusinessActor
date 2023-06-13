package entity

import "time"

type BusinessActor struct {
	Id               string `json:"id"`
	Series_reference string `form:"series_reference" json:"series_reference" bind:"required"`

	Period     time.Time `form:"period" json:"period" bind:"required"`
	Data_value float64   `form:"data_value" json:"data_value" bind:"required"`

	Suppressed bool `form:"suppressed" json:"suppressed" bind:"required"`

	STATUS    string `form:"status" json:"status" bind:"required"`
	UNITS     string `form:"units" json:"units" bind:"required"`
	Magnitude int    `form:"magnitude" json:"magnitude" bind:"required"`

	Subject string `form:"subject" json:"subject" bind:"required"`
	Group   string `form:"group" json:"group" bind:"required"`

	Series_title_1 string `form:"series_title_1" json:"series_title_1" bind:"required"`
	Series_title_2 string `form:"series_title_2" json:"series_title_2" bind:"required"`
	Series_title_3 string `form:"series_title_3" json:"series_title_3" bind:"required"`
	Series_title_4 string `form:"series_title_4" json:"series_title_4" bind:"required"`
	Series_title_5 string `form:"series_title_5" json:"series_title_5"`
}
