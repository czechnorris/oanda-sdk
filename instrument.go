package oanda_sdk

import (
	"github.com/shopspring/decimal"
	"time"
)

// CandlestickGranularity represents the granularity of a candlestick
type CandlestickGranularity string

const (
	// 5 second candlesticks, minute alignment
	S5 = CandlestickGranularity("S5")

	// 10 second candlesticks, minute alignment
	S10 = CandlestickGranularity("S10")

	// 15 second candlesticks, minute alignment
	S15 = CandlestickGranularity("S15")

	// 30 second candlesticks, minute alignment
	S30 = CandlestickGranularity("S30")

	// 1 minute candlesticks, minute alignment
	M1 = CandlestickGranularity("M1")

	// 2 minute candlesticks, hour alignment
	M2 = CandlestickGranularity("M2")

	// 4 minute candlesticks, hour alignment
	M4 = CandlestickGranularity("M4")

	// 5 minute candlesticks, hour alignment
	M5 = CandlestickGranularity("M5")

	// 10 minute candlesticks, hour alignment
	M10 = CandlestickGranularity("M10")

	// 15 minute candlesticks, hour alignment
	M15 = CandlestickGranularity("M15")

	// 30 minute candlesticks, hour alignment
	M30 = CandlestickGranularity("M30")

	// 1 hour candlesticks, hour alignment
	H1 = CandlestickGranularity("H1")

	// 2 hour candlesticks, day alignment
	H2 = CandlestickGranularity("H2")

	// 3 hour candlesticks, day alignment
	H3 = CandlestickGranularity("H3")

	// 4 hour candlesticks, day alignment
	H4 = CandlestickGranularity("H4")

	// 6 hour candlesticks, day alignment
	H6 = CandlestickGranularity("H6")

	// 8 hour candlesticks, day alignment
	H8 = CandlestickGranularity("H8")

	// 12 hour candlesticks, day alignment
	H12 = CandlestickGranularity("H12")

	// 1 day candlesticks, day alignment
	D = CandlestickGranularity("D")

	// 1 week candlesticks, aligned to start of week
	W = CandlestickGranularity("W")

	// 1 month candlesticks, aligned to first day of the month
	M = CandlestickGranularity("M")
)

// WeeklyAlignment represents the day of the week to use for candlestick granularities with weekly alignment
type WeeklyAlignment string

const (
	Monday    = WeeklyAlignment("Monday")
	Tuesday   = WeeklyAlignment("Tuesday")
	Wednesday = WeeklyAlignment("Wednesday")
	Thursday  = WeeklyAlignment("Thursday")
	Friday    = WeeklyAlignment("Friday")
	Saturday  = WeeklyAlignment("Saturday")
	Sunday    = WeeklyAlignment("Sunday")
)

// Candlestick represents the price candlestick
type Candlestick struct {
	// The start time of the candlestick
	Time time.Time `json:"time"`

	// The candlestick data based on bids. Only provided if bid-based candles were requested.
	Bid *CandlestickData `json:"bid"`

	// The candlestick data based on asks. Only provided if ask-based candles were requested.
	Ask *CandlestickData `json:"ask"`

	// The candlestick data based on midpoints. Only provided if midpoint-based candles were requested.
	Mid *CandlestickData `json:"mid"`

	// The number of prices created during the time-range represented by the candlestick.
	Volume int `json:"volume"`

	// A flag indicating if the candlestick is complete. A complete candlestick is one whose
	// ending time is not in the future.
	Complete bool `json:"complete"`
}

// CandlestickData represents the price data (open, high, low, close) for the Candlestick
type CandlestickData struct {
	// The first (open) price in the time-range represented by the candlestick.
	Open decimal.Decimal `json:"o"`

	// The highest price in the time-range represented by the candlestick.
	High decimal.Decimal `json:"h"`

	// The lowest price in the time-range represented by the candlestick.
	Low decimal.Decimal `json:"l"`

	// The last (closing) price in the time-range represented by the candlestick.
	Close decimal.Decimal `json:"c"`
}

// CandlestickResponse is a struct containing instrument, granularity, and list of candles
type CandlestickResponse struct {
	// The instrument whose Prices are represented by the candlesticks.
	Instrument string `json:"instrument"`

	// The granularity of the candlesticks provided.
	Granularity CandlestickGranularity `json:"granularity"`

	// The list of candlesticks that satisfy the request.
	Candles []Candlestick `json:"candles"`
}

// OrderBook is a representation of an instrument's order book at a point in time
type OrderBook struct {
	// The order book’s instrument
	Instrument string `json:"instrument"`

	// The time when the order book snapshot was created.
	Time time.Time `json:"time"`

	// The price (midpoint) for the order book’s instrument at the time of the order book snapshot
	Price decimal.Decimal `json:"price"`

	// The price width for each bucket. Each bucket covers the price range from the bucket’s price
	// to the bucket’s price + bucketWidth.
	BucketWidth decimal.Decimal `json:"bucketWidth"`

	// The partitioned order book, divided into buckets using a default bucket
	// width. These buckets are only provided for price ranges which actually
	// contain order or position data.
	Buckets []OrderBookBucket `json:"buckets"`
}

// OrderBookBucket holds the data for a partition of the instrument's prices
type OrderBookBucket struct {
	// The lowest price (inclusive) covered by the bucket. The bucket covers the
	// price range from the price to price + the order book’s bucketWidth.
	Price decimal.Decimal `json:"price"`

	// The percentage of the total number of orders represented by the long orders found in this bucket.
	LongCountPercent decimal.Decimal `json:"longCountPercent"`

	// The percentage of the total number of orders represented by the short orders found in this bucket.
	ShortCountPercent decimal.Decimal `json:"shortCountPercent"`
}

// PositionBook is a representation of an instrument's position book at a point in time
type PositionBook struct {
	// The position book’s instrument
	Instrument string `json:"instrument"`

	// The time when the position book snapshot was created
	Time time.Time `json:"time"`

	// The price (midpoint) for the position book’s instrument at the time of the position book snapshot
	Price decimal.Decimal `json:"price"`

	// The price width for each bucket. Each bucket covers the price range from the bucket’s price
	// to the bucket’s price + bucketWidth.
	BucketWidth decimal.Decimal `json:"bucketWidth"`

	// The partitioned position book, divided into buckets using a default
	// bucket width. These buckets are only provided for price ranges which
	// actually contain order or position data.
	Buckets []PositionBookBucket `json:"buckets"`
}

// PositionBookBucket holds the data for a partition of the instrument's prices
type PositionBookBucket struct {
	// The lowest price (inclusive) covered by the bucket. The bucket covers the
	// price range from the price to price + the position book’s bucketWidth.
	Price decimal.Decimal `json:"price"`

	// The percentage of the total number of positions represented by the long positions found in this bucket.
	LongCountPercent decimal.Decimal `json:"longCountPercent"`

	// The percentage of the total number of positions represented by the short positions found in this bucket.
	ShortCountPercent decimal.Decimal `json:"shortCountPercent"`
}
