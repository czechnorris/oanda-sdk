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

type GetAccountTradesRequest struct {
	// List of Trade IDs to retrieve.
	IDs []TradeID `url:"ids,comma"`

	// The state to filter the requested Trades by. [default=OPEN]
	State *TradeStateFilter `url:"state"`

	// The instrument to filter the requested Trades by.
	Instrument *string `url:"instrument"`

	// The maximum number of Trades to return. [default=50, maximum=500]
	Count int `url:"count"`

	// The maximum Trade ID to return. If not provided the most recent Trades in the Account are returned.
	BeforeID TradeID `url:"beforeID"`
}

type UpdateAccountTradeOrdersRequest struct {
	// The specification of the TakeProfit to create/modify/cancel. If
	// takeProfit is set to null, the TakeProfitOrder will be cancelled if it
	// exists. If takeProfit is not provided, the existing TakeProfitOrder
	// will not be modified. If a sub-field of takeProfit is not specified, that
	// field will be set to a default value on create, and be inherited by the
	// replacing order on modify.
	TakeProfit *TakeProfitDetails `json:"takeProfit"`

	// The specification of the StopLoss to create/modify/cancel. If stopLoss
	// is set to null, the StopLossOrder will be cancelled if it exists. If
	// stopLoss is not provided, the existing StopLossOrder will not be
	// modified. If a sub-field of stopLoss is not specified, that field will be
	// set to a default value on create, and be inherited by the replacing order
	// on modify.
	StopLoss *StopLossDetails `json:"stopLoss"`

	// The specification of the TrailingStopLoss to create/modify/cancel. If
	// trailingStopLoss is set to null, the TrailingStopLossOrder will be
	// cancelled if it exists. If trailingStopLoss is not provided, the existing
	// TrailingStopLossOrder will not be modified. If a sub-field of
	// trailingStopLoss is not specified, that field will be set to a default
	// value on create, and be inherited by the replacing order on modify.
	TrailingStopLoss *TrailingStopLossDetails `json:"trailingStopLoss"`

	// The specification of the GuaranteedStopLoss to create/modify/cancel. If
	// guaranteedStopLoss is set to null, the GuaranteedStopLossOrder will be
	// cancelled if it exists. If guaranteedStopLoss is not provided, the
	// existing GuaranteedStopLossOrder will not be modified. If a sub-field
	// of guaranteedStopLoss is not specified, that field will be set to a
	// default value on create, and be inherited by the replacing order on
	// modify.
	GuaranteedStopLoss *GuaranteedStopLossDetails `json:"guaranteedStopLoss"`
}

type CloseAccountInstrumentPositionRequest struct {
	// Indication of how much of the long Position to closeout. Either the
	// string “ALL”, the string “NONE”, or a DecimalNumber representing how many
	// units of the long position to close using a PositionCloseout MarketOrder.
	// The units specified must always be positive.
	// Default: ALL
	LongUnits *string `json:"longUnits,omitempty"`

	// The client extensions to add to the MarketOrder used to close the long position.
	LongClientExtensions ClientExtensions `json:"longClientExtensions"`

	// Indication of how much of the short Position to closeout. Either the
	// string “ALL”, the string “NONE”, or a DecimalNumber representing how many
	// units of the short position to close using a PositionCloseout
	// MarketOrder. The units specified must always be positive.
	// Default: ALL
	ShortUnits *string `json:"shortUnits,omitempty"`

	// The client extensions to add to the MarketOrder used to close the short position.
	ShortClientExtensions ClientExtensions
}
