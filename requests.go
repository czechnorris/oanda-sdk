package oanda_sdk

import (
	"github.com/shopspring/decimal"
	"time"
)

type SetAccountConfigurationRequest struct {
	// Client-defined alias (name) for the Account
	Alias string `json:"alias"`

	// The string representation of a decimal number. (serialization done by Client)
	MarginRate decimal.Decimal `json:"marginRate"`
}

type GetInstrumentCandlesRequest struct {
	// The Price component(s) to get candlestick data for. [default=M]
	Price *PricingComponent `url:"price,omitempty"`

	// The granularity of the candlesticks to fetch [default=S5]
	Granularity *CandlestickGranularity `url:"granularity,omitempty"`

	// The number of candlesticks to return in the response. Count should not be specified if both the start and end
	// parameters are provided, as the time range combined with the granularity will determine the number of
	// candlesticks to return. [default=500, maximum=5000]
	Count *int `url:"count,omitempty"`

	// The start of the time range to fetch candlesticks for.
	From *time.Time `url:"from,omitempty"`

	// The end of the time range to fetch candlesticks for.
	To *time.Time `url:"to,omitempty"`

	// A flag that controls whether the candlestick is “smoothed” or not. A smoothed candlestick uses the previous
	// candle’s close price as its open price, while an un-smoothed candlestick uses the first price from its time range
	// as its open price. [default=False]
	Smooth *bool `url:"smooth,omitempty"`

	// A flag that controls whether the candlestick that is covered by the from time should be included in the results.
	// This flag enables clients to use the timestamp of the last completed candlestick received to poll for future
	// candlesticks but avoid receiving the previous candlestick repeatedly. [default=True]
	IncludeFirst *bool `url:"includeFirst,omitempty"`

	// The hour of the day (in the specified timezone) to use for granularities that have daily alignments.
	// [default=17, minimum=0, maximum=23]
	DailyAlignment *int `url:"dailyAlignment,omitempty"`

	// The timezone to use for the dailyAlignment parameter. Candlesticks with daily alignment will be aligned to the
	// dailyAlignment hour within the alignmentTimezone. Note that the returned times will still be represented in UTC.
	// [default=America/New_York]
	AlignmentTimezone *string `url:"alignmentTimezone,omitempty"`

	// The day of the week used for granularities that have weekly alignment. [default=Friday]
	WeeklyAlignment *WeeklyAlignment `url:"weeklyAlignment,omitempty"`
}

type GetAccountOrdersRequest struct {
	// List of Order IDs to retrieve
	IDs []OrderID `url:"ids,comma"`

	// The state to filter the requested Orders by [default=PENDING]
	State *OrderStateFilter `url:"state"`

	// The instrument to filter the requested orders by
	Instrument *string `url:"instrument"`

	// The maximum number of Orders to return [default=50, maximum=500]
	Count int `url:"count"`

	// The maximum Order ID to return. If not provided the most recent Orders in the Account are returned
	BeforeID OrderID `url:"beforeID"`
}

type UpdateClientExtensionsRequest struct {
	// The ClientExtensions to update for the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The ClientExtensions to update for the Trade created when the Order is filled. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`
}
