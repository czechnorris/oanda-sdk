package oanda_sdk

import (
	"github.com/shopspring/decimal"
	"time"
)

// AccountID is the string representation of an account identifier
// Format: “-“-delimited string with format “{siteID}-{divisionID}-{userID}-{accountNumber}”
// Example: 001-011-5838423-001
type AccountID string

// Account contains the full details of a client's Account. This includes full open Trade, open Position and pending
// Order representation
type Account struct {
	// The Account’s identifier
	Id AccountID `json:"id"`

	// Client-assigned alias for the Account. Only provided if the Account has an alias set
	Alias string `json:"alias"`

	// The home currency of the Account
	Currency Currency `json:"currency"`

	// ID of the user that created the Account
	CreatedByUserID int `json:"createdByUserID"`

	// The date/time when the Account was created.
	CreatedTime time.Time `json:"createdTime"`

	// The current guaranteed Stop Loss Order settings of the Account. This
	// field will only be present if the guaranteedStopLossOrderMode is not ‘DISABLED’.
	GuaranteedStopLossOrderParameters GuaranteedStopLossOrderParameters `json:"guaranteedStopLossOrderParameters"`

	// The current guaranteed Stop Loss Order mode of the Account.
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`

	// The date/time that the Account’s resettablePL was last reset.
	ResettablePLTime time.Time `json:"resettablePLTime"`

	// Client-provided margin rate override for the Account. The effective
	// margin rate of the Account is the lesser of this value and the OANDA
	// margin rate for the Account’s division. This value is only provided if a
	// margin rate override exists for the Account.
	MarginRate decimal.Decimal `json:"marginRate"`

	// The number of Trades currently open in the Account.
	OpenTradeCount int `json:"openTradeCount"`

	// The number of Positions currently open in the Account.
	OpenPositionCount int `json:"openPositionCount"`

	// The number of Orders currently pending in the Account.
	PendingOrderCount int `json:"pendingOrderCount"`

	// Flag indicating that the Account has hedging enabled.
	HedgingEnabled bool `json:"hedgingEnabled"`

	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL decimal.Decimal `json:"unrealizedPL"`

	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV decimal.Decimal `json:"NAV"`

	// Margin currently used for the Account.
	MarginUsed decimal.Decimal `json:"marginUsed"`

	// Margin available for Account currency.
	MarginAvailable decimal.Decimal `json:"marginAvailable"`

	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue decimal.Decimal `json:"positionValue"`

	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL decimal.Decimal `json:"marginCloseoutUnrealizedPL"`

	// The Account’s margin closeout NAV.
	MarginCloseoutNAV decimal.Decimal `json:"marginCloseoutNAV"`

	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed decimal.Decimal `json:"marginCloseoutMarginUsed"`

	// The Account’s margin closeout percentage. When this value is 1.0 or above
	// the Account is in a margin closeout situation.
	MarginCloseoutPercent decimal.Decimal `json:"marginCloseoutPercent"`

	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue decimal.Decimal `json:"marginCloseoutPositionValue"`

	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit decimal.Decimal `json:"withdrawalLimit"`

	// The Account’s margin call margin used.
	MarginCallMarginUsed decimal.Decimal `json:"marginCallMarginUsed"`

	// The Account’s margin call percentage. When this value is 1.0 or above the Account is in a margin call situation.
	MarginCallPercent decimal.Decimal `json:"marginCallPercent"`

	// The current balance of the account.
	Balance decimal.Decimal `json:"balance"`

	// The total profit/loss realized over the lifetime of the Account.
	PL decimal.Decimal `json:"pl"`

	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL decimal.Decimal `json:"resettablePL"`

	// The total amount of financing paid/collected over the lifetime of the account.
	Financing decimal.Decimal `json:"financing"`

	// The total amount of commission paid over the lifetime of the Account.
	Commission decimal.Decimal `json:"commission"`

	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The total amount of fees charged over the lifetime of the Account for the execution of
	// guaranteed Stop Loss Orders.
	GuaranteedExecutionFees decimal.Decimal `json:"guaranteedExecutionFees"`

	// The date/time when the Account entered a margin call state. Only provided if the Account is in a margin call.
	MarginCallEnterTime time.Time `json:"marginCallEnterTime"`

	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount int `json:"marginCallExtensionCount"`

	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime time.Time `json:"lastMarginCallExtensionTime"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The details of the Trades currently open in the Account.
	Trades []TradeSummary `json:"trades"`

	// The details all Account Positions.
	Positions []Position `json:"positions"`

	// The details of the Orders currently pending in the Account.
	Orders []Order `json:"orders"`
}

// AccountChangesState struct is used to represent an Account’s current price-dependent state.
// Price-dependent Account state is dependent on OANDA’s current Prices, and includes things like unrealized PL, NAV and Trailing Stop Loss Order state.
// Fields will be omitted if their value has not changed since the specified transaction ID.
type AccountChangesState struct {
	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL *decimal.Decimal `json:"unrealizedPL"`

	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV *decimal.Decimal `json:"NAV"`

	// Margin currently used for the Account.
	MarginUsed *decimal.Decimal `json:"marginUsed"`

	// Margin available for Account currency.
	MarginAvailable *decimal.Decimal `json:"marginAvailable"`

	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue *decimal.Decimal `json:"positionValue"`

	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL *decimal.Decimal `json:"marginCloseoutUnrealizedPL"`

	// The Account’s margin closeout NAV.
	MarginCloseoutNAV *decimal.Decimal `json:"marginCloseoutNAV"`

	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed *decimal.Decimal `json:"marginCloseoutMarginUsed"`

	// The Account’s margin closeout percentage. When this value is 1.0 or above the Account is in
	// a margin closeout situation.
	MarginCloseoutPercent *decimal.Decimal `json:"marginCloseoutPercent"`

	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue *decimal.Decimal `json:"marginCloseoutPositionValue"`

	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit *decimal.Decimal `json:"withdrawalLimit"`

	// The Account’s margin call margin used.
	MarginCallMarginUsed *decimal.Decimal `json:"marginCallMarginUsed"`

	// The Account’s margin call percentage. When this value is 1.0 or above the Account is in
	// a margin call situation.
	MarginCallPercent *decimal.Decimal `json:"marginCallPercent"`

	// The current balance of the account.
	Balance *decimal.Decimal `json:"balance"`

	// The total profit/loss realized over the lifetime of the Account.
	PL *decimal.Decimal `json:"pl"`

	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL *decimal.Decimal `json:"resettablePL"`

	// The total amount of financing paid/collected over the lifetime of the account.
	Financing *decimal.Decimal `json:"financing"`

	// The total amount of commission paid over the lifetime of the Account.
	Commission *decimal.Decimal `json:"commission"`

	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment *decimal.Decimal `json:"dividendAdjustment"`

	// The total amount of fees charged over the lifetime of the Account for the execution
	// of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees *decimal.Decimal `json:"guaranteedExecutionFees"`

	// The date/time when the Account entered a margin call state. Only provided if the Account is in a margin call.
	MarginCallEnterTime *time.Time `json:"marginCallEnterTime"`

	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount *int `json:"marginCallExtensionCount"`

	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime *time.Time `json:"lastMarginCallExtensionTime"`

	// The price-dependent state of each pending Order in the Account.
	Orders []DynamicOrderState `json:"orders"`

	// The price-dependent state for each open Trade in the Account.
	Trades []CalculatedTradeState `json:"trades"`

	// The price-dependent state for each open Position in the Account.
	Positions []CalculatedPositionState `json:"positions"`
}

// AccountProperties is a struct of properties related to an Account
type AccountProperties struct {
	// The Account’s identifier
	ID AccountID `json:"id"`

	// The Account’s associated MT4 Account ID. This field will not be present if the Account is not an MT4 account.
	Mt4AccountID *int `json:"mt4AccountID"`

	// The Account’s tags
	Tags []string `json:"tags"`
}

// GuaranteedStopLossOrderParameters is a struct carrying the information about the current mutability and hedging
// settings related to guaranteed Stop Loss orders
type GuaranteedStopLossOrderParameters struct {
	// The current guaranteed Stop Loss Order mutability setting of the Account when market is open.
	MutabilityMarketOpen GuaranteedStopLossOrderMutability `json:"mutabilityMarketOpen"`

	// The current guaranteed Stop Loss Order mutability setting of the Account when market is halted.
	MutabilityMarketHalted GuaranteedStopLossOrderMutability `json:"mutabilityMarketHalted"`
}

// GuaranteedStopLossOrderMode represents the overall behaviour of the Account regarding guaranteed Stop Loss Orders.
type GuaranteedStopLossOrderMode string

const (
	// Disabled represents the state when the Account is not permitted to create guaranteed Stop Loss Orders.
	Disabled = GuaranteedStopLossOrderMode("DISABLED")

	// Allowed represents the state when the Account is able, but not required to have guaranteed Stop Loss Orders
	// for open Trades.
	Allowed = GuaranteedStopLossOrderMode("ALLOWED")

	// Required represents the state when the Account is required to have guaranteed Stop Loss Orders
	// for all open Trades.
	Required = GuaranteedStopLossOrderMode("REQUIRED")
)

// GuaranteedStopLossOrderMutability for Accounts that support guaranteed Stop Loss Order, describes the actions
// that can be performed on guaranteed Stop Loss Orders.
type GuaranteedStopLossOrderMutability string

const (
	// Fixed represents mutability where once a guaranteed Stop Loss Order has been created it cannot be
	// replaced or cancelled.
	Fixed = GuaranteedStopLossOrderMutability("FIXED")

	// Replaceable represents mutability where an existing guaranteed Stop Loss Order can only be
	// replaced, not cancelled.
	Replaceable = GuaranteedStopLossOrderMutability("REPLACEABLE")

	// Cancelable represents mutability where once a guaranteed Stop Loss Order has been created it can be
	// either replaced or cancelled.
	Cancelable = GuaranteedStopLossOrderMutability("CANCELABLE")

	// PriceWidenOnly represents mutability where an existing guaranteed Stop Loss Order can only be
	// replaced to widen the gap from the current price, not cancelled.
	PriceWidenOnly = GuaranteedStopLossOrderMutability("PRICE_WIDEN_ONLY")
)

// AccountSummary is a summary representation of a client's Account. The AccountSummary does not provide to full
// specification of pending Orders, open Trades and Positions
type AccountSummary Account

// TODO: Maybe make Account an extension of AccountSummary?

type AccumulatedAccountState AccountSummary

// TODO: this is extended by AccountSummary and AccountChangesState

type CalculatedAccountState AccountSummary

// TODO: this is extended by AccountSummary and AccountChangesState

// AccountChanges struct is used to represent the changes to an Account's Orders, Trades and Positions since
// a specified Account TransactionID in the past.
type AccountChanges struct {
	// The Orders created. These Orders may have been filled, cancelled or triggered in the same period.
	OrdersCreated []Order `json:"ordersCreated"`

	// The Orders cancelled.
	OrdersCancelled []Order `json:"ordersCancelled"`

	// The Orders filled.
	OrdersFilled []Order `json:"ordersFilled"`

	// The Orders triggered.
	OrdersTriggered []Order `json:"ordersTriggered"`

	// The Trades opened.
	TradesOpened []TradeSummary `json:"tradesOpened"`

	// The Trades reduced.
	TradesReduced []TradeSummary `json:"tradesReduced"`

	// The Trades closed.
	TradesClosed []TradeSummary `json:"tradesClosed"`

	// The Positions changed.
	Positions []Position `json:"positions"`

	// The Transactions that have been generated.
	Transactions []Transaction `json:"transactions"`
}

// AccountFinancingMode represents the financing mode of an Account
type AccountFinancingMode string

const (
	// NoFinancing means no financing is paid/charged for open Trades in the Account
	NoFinancing = AccountFinancingMode("NO_FINANCING")

	// SecondBySecond financing is paid/charged for open Trades in the Account,
	// both daily and when the Trade is closed
	SecondBySecond = AccountFinancingMode("SECOND_BY_SECOND")

	// Daily means a full day’s worth of financing is paid/charged for open Trades
	// in the Account daily at 5pm New York time
	Daily = AccountFinancingMode("DAILY")
)

// UserAttributes contain the attributes of a user
type UserAttributes struct {
	// The user’s OANDA-assigned user ID.
	UserID int `json:"userID"`

	// The user-provided username.
	Username string `json:"username"`

	// The user’s title.
	Title string `json:"title"`

	// The user’s name.
	Name string `json:"name"`

	// The user’s email address.
	Email string `json:"email"`

	// The OANDA division the user belongs to.
	DivisionAbbreviation string `json:"divisionAbbreviation"`

	// The user’s preferred language.
	LanguageAbbreviation string `json:"languageAbbreviation"`

	// The home currency of the Account.
	HomeCurrency Currency `json:"homeCurrency"`
}

// PositionAggregationMode represents the way that position values for an Account are calculated and aggregated
type PositionAggregationMode string

const (
	// AbsoluteSum means the Position value or margin for each side (long and short) of the Position
	// are computed independently and added together.
	AbsoluteSum = PositionAggregationMode("ABSOLUTE_SUM")

	// MaximalSide means the Position value or margin for each side (long and short) of the Position
	// are computed independently. The Position value or margin chosen is the maximal absolute value of the two.
	MaximalSide = PositionAggregationMode("MAXIMAL_SIDE")

	// NetSum means the units for each side (long and short) of the Position are netted together and
	// the resulting value (long or short) is used to compute the Position value or margin.
	NetSum = PositionAggregationMode("NET_SUM")
)
