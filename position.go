package oanda_sdk

import "github.com/shopspring/decimal"

// Position specifies a Position within an Account.
type Position struct {
	// The Position’s Instrument.
	Instrument string `json:"instrument"`

	// Profit/loss realized by the Position over the lifetime of the Account.
	PL decimal.Decimal `json:"pl"`

	// The unrealized profit/loss of all open Trades that contribute to this Position.
	UnrealizedPL decimal.Decimal `json:"unrealizedPL"`

	// Margin currently used by the Position.
	MarginUsed decimal.Decimal `json:"marginUsed"`

	// Profit/loss realized by the Position since the Account’s resettablePL was last reset by the client.
	ResettablePL decimal.Decimal `json:"resettablePL"`

	// The total amount of financing paid/collected for this instrument over the lifetime of the Account.
	Financing decimal.Decimal `json:"financing"`

	// The total amount of commission paid for this instrument over the lifetime of the Account.
	Commission decimal.Decimal `json:"commission"`

	// The total amount of dividend adjustment paid for this instrument over the lifetime of the Account.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The total amount of fees charged over the lifetime of the Account for the execution of guaranteed StopLossOrders
	// for this instrument.
	GuaranteedExecutionFees decimal.Decimal `json:"guaranteedExecutionFees"`

	// The details of the long side of the Position.
	Long PositionSide `json:"long"`

	// The details of the short side of the Position.
	Short PositionSide `json:"short"`
}

// PositionSide is a representation of a Position for a single direction (long or short).
type PositionSide struct {
	// Number of units in the position (negative value indicates short position, positive indicates long position).
	Units decimal.Decimal `json:"units"`

	// Volume-weighted average of the underlying Trade open prices for the Position.
	AveragePrice decimal.Decimal `json:"averagePrice"`

	// List of the open Trade IDs which contribute to the open Position.
	TradeIDs []TradeID `json:"tradeIDs"`

	// Profit/loss realized by the PositionSide over the lifetime of the Account.
	PL decimal.Decimal `json:"pl"`

	// The unrealized profit/loss of all open Trades that contribute to this PositionSide.
	UnrealizedPL decimal.Decimal `json:"unrealizedPL"`

	// Profit/loss realized by the PositionSide since the Account’s resettablePL was last reset by the client.
	ResettablePL decimal.Decimal `json:"resettablePL"`

	// The total amount of financing paid/collected for this PositionSide over the lifetime of the Account.
	Financing decimal.Decimal `json:"financing"`

	// The total amount of dividend adjustment paid for the PositionSide over the lifetime of the Account.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The total amount of fees charged over the lifetime of the Account for the execution of guaranteed StopLossOrders
	// attached to Trades for this PositionSide.
	GuaranteedExecutionFees decimal.Decimal `json:"guaranteedExecutionFees"`
}

// CalculatedPositionState represents the dynamic (calculated) state of a Position
type CalculatedPositionState struct {
	// The Position’s Instrument.
	Instrument string `json:"instrument"`

	// The Position’s net unrealized profit/loss
	NetUnrealizedPL decimal.Decimal `json:"netUnrealizedPL"`

	// The unrealized profit/loss of the Position’s long open Trades
	LongUnrealizedPL decimal.Decimal `json:"longUnrealizedPL"`

	// The unrealized profit/loss of the Position’s short open Trades
	ShortUnrealizedPL decimal.Decimal `json:"shortUnrealizedPL"`

	// Margin currently used by the Position.
	MarginUsed decimal.Decimal `json:"marginUsed"`
}
