package models

type CandleResolution string

const (
	Minute         CandleResolution = "1"
	FiveMinutes    CandleResolution = "5"
	FifteenMinutes CandleResolution = "15"
	Day            CandleResolution = "D"
	Week           CandleResolution = "W"
	Month          CandleResolution = "M"
)

type Candle struct {
	// Open - open prices
	Open []float32 `json:"o"`

	// Close - close prices
	Close []float32 `json:"c"`

	// High - high prices
	High []float32 `json:"h"`

	// Low - low prices
	Low []float32 `json:"l"`

	// Volume - volume data
	Volume []int `json:"v"`

	// List of timestamps in UNIX format
	Timestamps []int64 `json:"t"`

	// Status of the response.
	// This field can either be `ok` or `no_data`.
	Status string `json:"s"`
}
