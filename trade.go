package oanda_sdk

import (
	"github.com/shopspring/decimal"
	"time"
)

// TradeID is a Trade identifier, unique within the Trade's Account.
type TradeID string

// TradeState represents the current state of the Trade.
type TradeState string

const (
	// TradeStateOpen represents the state of a Trade that is currently open
	TradeStateOpen = TradeState("OPEN")

	// TradeStateClosed represents the state of a Trade that has been fully closed
	TradeStateClosed = TradeState("CLOSED")

	// TradeStateCloseWhenTradeable represents the state of a Trade that will be closed as soon as the trade's
	// instrument becomes tradeable
	TradeStateCloseWhenTradeable = TradeState("CLOSE_WHEN_TRADEABLE")
)

// TradeStateFilter represents the state to filter the Trades by
type TradeStateFilter string

const (
	// TradeStateFilterOpen filters the Trades that are currently open
	TradeStateFilterOpen = TradeStateFilter("OPEN")

	// TradeStateFilterClosed filters the Trades that have been fully closed
	TradeStateFilterClosed = TradeStateFilter("CLOSED")

	// TradeStateFilterCloseWhenTradeable filters the Trades that will be closed as soon as the trades' instrument
	// becomes tradeable
	TradeStateFilterCloseWhenTradeable = TradeStateFilter("CLOSE_WHEN_TRADEABLE")

	// TradeStateFilterAll filters the Trades that are in any of the possible states listed above
	TradeStateFilterAll = TradeStateFilter("ALL")
)

// TradeSpecifier represents an identification of a Trade as referred to by clients
// Format: Either the Trade’s OANDA-assigned TradeID or the Trade’s client-provided ClientID prefixed by the “@” symbol
// Example: @my_trade_id
type TradeSpecifier string

// Trade struct specifies a Trade within an Account. This includes the full representation of the Trade's dependent
// Orders in addition to the IDs of those Orders.
type Trade struct {
	// The Trade’s identifier, unique within the Trade’s Account.
	Id TradeID `json:"id"`

	// The Trade’s Instrument.
	Instrument string `json:"instrument"`

	// The execution price of the Trade.
	Price decimal.Decimal `json:"price"`

	// The date/time when the Trade was opened.
	OpenTime time.Time `json:"openTime"`

	// The current state of the Trade.
	State TradeState `json:"state"`

	// The initial size of the Trade. Negative values indicate a short Trade, and positive values indicate a long Trade.
	InitialUnits decimal.Decimal `json:"initialUnits"`

	// The margin required at the time the Trade was created. Note, this is the
	// ‘pure’ margin required, it is not the ‘effective’ margin used that
	// factors in the trade risk if a GSLO is attached to the trade.
	InitialMarginRequired decimal.Decimal `json:"initialMarginRequired"`

	// The number of units currently open for the Trade. This value is reduced to 0.0 as the Trade is closed.
	CurrentUnits decimal.Decimal `json:"currentUnits"`

	// The total profit/loss realized on the closed portion of the Trade.
	RealizedPL decimal.Decimal `json:"realizedPL"`

	// The unrealized profit/loss on the open portion of the Trade.
	UnrealizedPL decimal.Decimal `json:"unrealizedPL"`

	// Margin currently used by the Trade.
	MarginUsed decimal.Decimal `json:"marginUsed"`

	// The average closing price of the Trade. Only present if the Trade has been closed or reduced at least once.
	AverageClosePrice decimal.Decimal `json:"averageClosePrice"`

	// The IDs of the Transactions that have closed portions of this Trade.
	ClosingTransactionIDs []TransactionID `json:"closingTransactionIDs"`

	// The financing paid/collected for this Trade.
	Financing decimal.Decimal `json:"financing"`

	// The dividend adjustment paid for this Trade.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The date/time when the Trade was fully closed. Only provided for Trades whose state is CLOSED.
	CloseTime time.Time `json:"closeTime"`

	// The client extensions of the Trade.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// Full representation of the Trade’s TakeProfitOrder, only provided if such an Order exists.
	TakeProfitOrder *TakeProfitOrder `json:"takeProfitOrder"`

	// Full representation of the Trade’s StopLossOrder, only provided if such an Order exists.
	StopLossOrder *StopLossOrder `json:"stopLossOrder"`

	// Full representation of the Trade’s TrailingStopLossOrder, only provided if such an Order exists.
	TrailingStopLossOrder *TrailingStopLossOrder `json:"trailingStopLossOrder"`
}

// TradeSummary represents the summary of a Trade within an Account. This representation does not provide the full
// details of the Trade's dependent Orders.
type TradeSummary struct {
	// The Trade’s identifier, unique within the Trade’s Account.
	Id TradeID `json:"id"`

	// The Trade’s Instrument.
	Instrument string `json:"instrument"`

	// The execution price of the Trade.
	Price decimal.Decimal `json:"price"`

	// The date/time when the Trade was opened.
	OpenTime time.Time `json:"openTime"`

	// The current state of the Trade.
	State TradeState `json:"state"`

	// The initial size of the Trade. Negative values indicate a short Trade, and positive values indicate a long Trade.
	InitialUnits decimal.Decimal `json:"initialUnits"`

	// The margin required at the time the Trade was created. Note, this is the
	// ‘pure’ margin required, it is not the ‘effective’ margin used that
	// factors in the trade risk if a GSLO is attached to the trade.
	InitialMarginRequired decimal.Decimal `json:"initialMarginRequired"`

	// The number of units currently open for the Trade. This value is reduced to 0.0 as the Trade is closed.
	CurrentUnits decimal.Decimal `json:"currentUnits"`

	// The total profit/loss realized on the closed portion of the Trade.
	RealizedPL decimal.Decimal `json:"realizedPL"`

	// The unrealized profit/loss on the open portion of the Trade.
	UnrealizedPL decimal.Decimal `json:"unrealizedPL"`

	// Margin currently used by the Trade.
	MarginUsed decimal.Decimal `json:"marginUsed"`

	// The average closing price of the Trade. Only present if the Trade has been closed or reduced at least once.
	AverageClosePrice decimal.Decimal `json:"averageClosePrice"`

	// The IDs of the Transactions that have closed portions of this Trade.
	ClosingTransactionIDs []TransactionID `json:"closingTransactionIDs"`

	// The financing paid/collected for this Trade.
	Financing decimal.Decimal `json:"financing"`

	// The dividend adjustment paid for this Trade.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The date/time when the Trade was fully closed. Only provided for Trades whose state is CLOSED.
	CloseTime time.Time `json:"closeTime"`

	// The client extensions of the Trade.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// ID of the Trade’s TakeProfitOrder, only provided if such an Order exists.
	TakeProfitOrderID *OrderID `json:"takeProfitOrderID"`

	// ID of the Trade’s StopLossOrder, only provided if such an Order exists.
	StopLossOrderID *OrderID `json:"stopLossOrderID"`

	// ID of the Trade’s GuaranteedStopLossOrder, only provided if such an Order exists.
	GuaranteedStopLossOrderID *OrderID `json:"guaranteedStopLossOrderID"`

	// ID of the Trade’s TrailingStopLossOrder, only provided if such an Order exists.
	TrailingStopLossOrderID *OrderID `json:"trailingStopLossOrderID"`
}

// CalculatedTradeState represents the dynamic (calculated) state of an open Trade
type CalculatedTradeState struct {
	// The Trade's ID.
	ID TradeID `json:"id"`

	// The Trade's unrealized profit/loss.
	UnrealizedPL decimal.Decimal `json:"unrealizedPL"`

	// Margin currently used by the Trade.
	MarginUsed decimal.Decimal `json:"marginUsed"`
}

// TradePL is the classification of TradePLs
type TradePL string

const (
	// TradePLPositive represents a state where an open Trade currently has a positive (profitable) unrealized P/L, or a
	// closed Trade realized a positive amount of P/L.
	TradePLPositive = TradePL("POSITIVE")

	// TradePLNegative represents a state where an open Trade currently has a negative (losing) unrealized P/L, or a
	// closed Trade realized a negative amount of P/L.
	TradePLNegative = TradePL("NEGATIVE")

	// TradePLZero represents a state where an open Trade currently has unrealized P/L of zero (neither profitable nor
	// losing), or a closed Trade realized a P/L amount of zero.
	TradePLZero = TradePL("ZERO")
)
