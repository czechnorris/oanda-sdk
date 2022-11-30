package oanda_sdk

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction interface {
	GetType() TransactionType
}

type TransactionBase struct {
	// The Transaction’s Identifier.
	Id TransactionID `json:"id"`

	// The date/time when the Transaction was created.
	Time time.Time `json:"time"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int `json:"userID"`

	// The ID of the Account the Transaction was created for.
	AccountID AccountID `json:"accountID"`

	// The ID of the “batch” that the Transaction belongs to. Transactions in the same batch are applied to the Account
	// simultaneously.
	BatchID TransactionID `json:"batchID"`

	// The Request ID of the request which generated the transaction.
	RequestID RequestID `json:"requestID"`

	// The Type of the Transaction.
	Type TransactionType `json:"type"`
}

// CreateTransaction represents the creation of an Account.
type CreateTransaction struct {
	TransactionBase

	// The ID of the Division that the Account is in
	DivisionID int `json:"divisionID"`

	// The ID of the Site that the Account was created at
	SiteID int `json:"siteID"`

	// The ID of the user that the Account was created for
	AccountUserID int `json:"accountUserID"`

	// The number of the Account within the site/division/user
	AccountNumber int `json:"accountNumber"`

	// The home currency of the Account
	HomeCurrency Currency `json:"homeCurrency"`
}

func (ct CreateTransaction) GetType() TransactionType {
	return TransactionTypeCreate
}

// CloseTransaction represents the closing of an Account.
type CloseTransaction struct {
	TransactionBase
}

func (ct CloseTransaction) GetType() TransactionType {
	return TransactionTypeClose
}

// ReopenTransaction represents the re-opening of a closed Account.
type ReopenTransaction struct {
	TransactionBase
}

func (rt ReopenTransaction) GetType() TransactionType {
	return TransactionTypeReopen
}

// ClientConfigureTransaction represents the configuration of an Account by a client.
type ClientConfigureTransaction struct {
	TransactionBase

	// The client-provided alias for the Account.
	Alias string `json:"alias"`

	// The margin rate override for the Account.
	MarginRage decimal.Decimal `json:"marginRage"`
}

func (cct ClientConfigureTransaction) GetType() TransactionType {
	return TransactionTypeClientConfigure
}

// ClientConfigureRejectTransaction represents the reject of configuration of an Account by a client
type ClientConfigureRejectTransaction struct {
	ClientConfigureTransaction
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (ccrt ClientConfigureRejectTransaction) GetType() TransactionType {
	return TransactionTypeClientConfigureReject
}

// TransferFundsTransaction represents the transfer of funds in/out of an Account.
type TransferFundsTransaction struct {
	TransactionBase

	// The amount to deposit/withdraw from the Account in the Account’s home
	// currency. A positive value indicates a deposit, a negative value
	// indicates a withdrawal.
	Amount decimal.Decimal `json:"amount"`

	// The reason that an Account is being funded.
	FundingReason FundingReason `json:"fundingReason"`

	// An optional comment that may be attached to a fund transfer for audit purposes
	Comment string `json:"comment"`

	// The Account’s balance after funds are transferred.
	AccountBalance decimal.Decimal `json:"accountBalance"`
}

func (tft TransferFundsTransaction) GetType() TransactionType {
	return TransactionTypeTransferFunds
}

// TransferFundsRejectTransaction represents the rejection of the transfer of funds in/out of an Account.
type TransferFundsRejectTransaction struct {
	TransactionBase

	// The amount to deposit/withdraw from the Account in the Account’s home
	// currency. A positive value indicates a deposit, a negative value
	// indicates a withdrawal.
	Amount decimal.Decimal `json:"amount"`

	// The reason that an Account is being funded.
	FundingReason FundingReason `json:"fundingReason"`

	// An optional comment that may be attached to a fund transfer for audit purposes
	Comment string `json:"comment"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// MarketOrderTransaction represents the creation of a MarketOrder in the user's account. A MarketOrder is an Order that
// is filled immediately at the current market price. MarketOrders can be specialized when they are created to
// accomplish a specific task: to close a Trade, to closeout a Position or to participate in a Margin closeout
type MarketOrderTransaction struct {
	TransactionBase

	// The Market Order’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the MarketOrder. A positive number of units results in a long Order, and
	// a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The time-in-force requested for the MarketOrder. Restricted to FOK or IOC for a MarketOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The worst price that the client is willing to have the MarketOrder filled at.
	PriceBound decimal.Decimal `json:"priceBound"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

	// Details of the Trade requested to be closed, only provided when the MarketOrder is being used to explicitly
	// close a Trade.
	TradeClose MarketOrderTradeClose `json:"tradeClose"`

	// Details of the long Position requested to be closed out, only provided
	// when a MarketOrder is being used to explicitly closeout a long Position.
	LongPositionCloseout MarketOrderPositionCloseout `json:"longPositionCloseout"`

	// Details of the short Position requested to be closed out, only provided when a MarketOrder is being used to
	// explicitly closeout a short Position.
	ShortPositionCloseout MarketOrderPositionCloseout `json:"shortPositionCloseout"`

	// Details of the Margin Closeout that this MarketOrder was created for
	MarginCloseout MarketOrderMarginCloseout `json:"marginCloseout"`

	// Details of the delayed Trade close that this MarketOrder was created for
	DelayedTradeClose MarketOrderDelayedTradeClose `json:"delayedTradeClose"`

	// The reason that the MarketOrder was created
	Reason MarketOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The specification of the TakeProfitOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// The specification of the StopLossOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// The specification of the TrailingStopLossOrder that should be created for a Trade that is opened when the
	// Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// The specification of the GuaranteedStopLossOrder that should be created for a Trade that is opened when the Order
	// is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled (if such a Trade is created).  Do not set,
	// modify, delete tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`
}

func (mot MarketOrderTransaction) GetType() TransactionType {
	return TransactionTypeMarketOrder
}

// MarketOrderRejectTransaction represents the rejection of the creation of a MarketOrder
type MarketOrderRejectTransaction struct {
	MarketOrderTransaction

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (mort MarketOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeMarketOrderReject
}

type FixedPriceOrderTransaction struct {
	TransactionBase

	// The FixedPriceOrder’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the Fixed Price Order. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price specified for the FixedPriceOrder. This price is the exact price that the FixedPriceOrder will be
	// filled at.
	Price decimal.Decimal `json:"price"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Default: OrderPositionFillDefault
	PositionFill *OrderPositionFill `json:"positionFill"`

	// The state that the trade resulting from the FixedPriceOrder should be set to.
	TradeState TradeState `json:"tradeState"`

	// The reason that the Fixed Price Order was created
	Reason *FixedPriceOrderReason `json:"reason"`

	// The client extensions for the Fixed Price Order.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the TakeProfitOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// The specification of the StopLossOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// The specification of the TrailingStopLossOrder that should be created for a Trade that is opened when the
	// Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// The specification of the GuaranteedStopLossOrder that should be created for a Trade that is opened when the Order
	// is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled (if such a Trade is created).  Do not set,
	// modify, delete tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
}

func (fpot FixedPriceOrderTransaction) GetType() TransactionType {
	return TransactionTypeFixedPriceOrder
}

// LimitOrderTransaction represents the creation of a LimitOrder in the user's Account
type LimitOrderTransaction struct {
	TransactionBase

	// The FixedPriceOrder’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the Fixed Price Order. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price specified for the FixedPriceOrder. This price is the exact price that the FixedPriceOrder will be
	// filled at.
	Price decimal.Decimal `json:"price"`

	// The time-in-force requested for the LimitOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the LimitOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Default: OrderPositionFillDefault
	PositionFill *OrderPositionFill `json:"positionFill"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The reason that the LimitOrder was initiated
	Reason *LimitOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The specification of the TakeProfitOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// The specification of the StopLossOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// The specification of the TrailingStopLossOrder that should be created for a Trade that is opened when the
	// Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// The specification of the GuaranteedStopLossOrder that should be created for a Trade that is opened when the Order
	// is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled (if such a Trade is created).  Do not set,
	// modify, delete tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (lot LimitOrderTransaction) GetType() TransactionType {
	return TransactionTypeLimitOrder
}

// LimitOrderRejectTransaction represents the rejection of the creation of a LimitOrder
type LimitOrderRejectTransaction struct {
	LimitOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (lort LimitOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeLimitOrderReject
}

// StopOrderTransaction represents the creation of a StopOrder in the user's Account.
type StopOrderTransaction struct {
	TransactionBase

	// The StopOrder’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the StopOrder. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price threshold specified for the StopOrder. The StopOrder will
	// only be filled by a market price that is equal to or worse than this
	// price.
	Price decimal.Decimal `json:"price"`

	// The worst market price that may be used to fill this StopOrder. If the
	// market gaps and crosses through both the price and the priceBound, the
	// StopOrder will be cancelled instead of being filled.
	PriceBound decimal.Decimal `json:"priceBound"`

	// The time-in-force requested for the StopOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the LimitOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Default: OrderPositionFillDefault
	PositionFill *OrderPositionFill `json:"positionFill"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The reason that the StopOrder was initiated
	Reason *StopOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The specification of the TakeProfitOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// The specification of the StopLossOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// The specification of the TrailingStopLossOrder that should be created for a Trade that is opened when the
	// Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// The specification of the GuaranteedStopLossOrder that should be created for a Trade that is opened when the Order
	// is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled (if such a Trade is created).  Do not set,
	// modify, delete tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (sot StopOrderTransaction) GetType() TransactionType {
	return TransactionTypeStopOrder
}

// StopOrderRejectTransaction represents the rejection of the creation of a StopOrder
type StopOrderRejectTransaction struct {
	StopOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (sort StopOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeStopOrderReject
}

// MarketIfTouchedOrderTransaction represents the creation of a MarketIfTouchedOrder in the user's Account.
type MarketIfTouchedOrderTransaction struct {
	TransactionBase

	// The MarketIfTouchedOrder’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the MarketIfTouchedOrder. A positive number
	// of units results in a long Order, and a negative number of units results
	// in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price threshold specified for the MarketIfTouchedOrder. The MarketIfTouchedOrder will
	// only be filled by a market price that is equal to or worse than this
	// price.
	Price decimal.Decimal `json:"price"`

	// The worst market price that may be used to fill this MarketIfTouchedOrder. If the
	// market gaps and crosses through both the price and the priceBound, the
	// MarketIfTouchedOrder will be cancelled instead of being filled.
	PriceBound decimal.Decimal `json:"priceBound"`

	// The time-in-force requested for the MarketIfTouchedOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the LimitOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Default: OrderPositionFillDefault
	PositionFill *OrderPositionFill `json:"positionFill"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The reason that the MarketIfTouchedOrder was initiated
	Reason *MarketIfTouchedOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The specification of the TakeProfitOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`

	// The specification of the StopLossOrder that should be created for a Trade opened when the Order is filled
	// (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`

	// The specification of the TrailingStopLossOrder that should be created for a Trade that is opened when the
	// Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// The specification of the GuaranteedStopLossOrder that should be created for a Trade that is opened when the Order
	// is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled (if such a Trade is created).  Do not set,
	// modify, delete tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (mitot MarketIfTouchedOrderTransaction) GetType() TransactionType {
	return TransactionTypeMarketIfTouchedOrder
}

// MarketIfTouchedOrderRejectTransaction represents the rejection of the creation of a MarketIfTouchedOrder
type MarketIfTouchedOrderRejectTransaction struct {
	MarketIfTouchedOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (mitort MarketIfTouchedOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeMarketIfTouchedOrderReject
}

// TakeProfitOrderTransaction represents the creation of a TakeProfitOrder in the user's Account
type TakeProfitOrderTransaction struct {
	TransactionBase

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID *ClientID `json:"clientTradeID"`

	// The price threshold specified for the TakeProfitOrder. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	Price decimal.Decimal `json:"price"`

	// The time-in-force requested for the TakeProfitOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the TakeProfitOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The reason that the TakeProfitOrder was initiated
	Reason *TakeProfitOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	OrderFillTransactionID *TransactionID `json:"orderFillTransactionID"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *TransactionID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (TakeProfitOrderTransaction) GetType() TransactionType {
	return TransactionTypeTakeProfitOrder
}

// TakeProfitOrderRejectTransaction represents the rejection of the creation of a TakeProfitOrder
type TakeProfitOrderRejectTransaction struct {
	TakeProfitOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (TakeProfitOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeTakeProfitOrderReject
}

// StopLossOrderTransaction represents the creation of a StopLossOrder in the user's Account
type StopLossOrderTransaction struct {
	TransactionBase

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID *ClientID `json:"clientTradeID"`

	// The price threshold specified for the StopLossOrder. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	Price decimal.Decimal `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the StopLossOrder price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance *decimal.Decimal `json:"distance"`

	// The time-in-force requested for the StopLossOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the StopLossOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The reason that the StopLossOrder was initiated
	Reason *StopLossOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	OrderFillTransactionID *TransactionID `json:"orderFillTransactionID"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *TransactionID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (StopLossOrderTransaction) GetType() TransactionType {
	return TransactionTypeStopLossOrder
}

// StopLossOrderRejectTransaction represents the rejection of the creation of a StopLossOrder
type StopLossOrderRejectTransaction struct {
	StopLossOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (StopLossOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeStopLossOrderReject
}

// GuaranteedStopLossOrderTransaction represents the creation of a GuaranteedStopLossOrder in the user's Account
type GuaranteedStopLossOrderTransaction struct {
	TransactionBase

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID *ClientID `json:"clientTradeID"`

	// The price threshold specified for the StopLossOrder. The associated
	// Trade will be closed by a market price that is equal to or better than
	// this threshold.
	Price decimal.Decimal `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the StopLossOrder price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance *decimal.Decimal `json:"distance"`

	// The time-in-force requested for the GuaranteedStopLossOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the GuaranteedStopLossOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The fee that will be charged if the GuaranteedStopLossOrder is filled
	// at the guaranteed price. The value is determined at Order creation time.
	// It is in price units and is charged for each unit of the Trade.
	GuaranteedExecutionPremium decimal.Decimal `json:"guaranteedExecutionPremium"`

	// The reason that the GuaranteedStopLossOrder was initiated
	Reason *GuaranteedStopLossOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	OrderFillTransactionID *TransactionID `json:"orderFillTransactionID"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *TransactionID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (GuaranteedStopLossOrderTransaction) GetType() TransactionType {
	return TransactionTypeGuaranteedStopLossOrder
}

// GuaranteedStopLossOrderRejectTransaction represents the rejection of the creation of a GuaranteedStopLossOrder.
type GuaranteedStopLossOrderRejectTransaction struct {
	GuaranteedStopLossOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (GuaranteedStopLossOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeGuaranteedStopLossOrderReject
}

// TrailingStopLossOrderTransaction represents the creation of a TrailingStopLossOrder in the user's Account.
type TrailingStopLossOrderTransaction struct {
	TransactionBase

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID *ClientID `json:"clientTradeID"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the TrailingStopLossOrder price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance *decimal.Decimal `json:"distance"`

	// The time-in-force requested for the TrailingStopLossOrder.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce"`

	// The date/time when the TrailingStopLossOrder will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of which price component should be used when determining if
	// an Order should be triggered and filled. This allows Orders to be
	// triggered based on the bid, ask, mid, default (ask for buy, bid for sell)
	// or inverse (ask for sell, bid for buy) price depending on the desired
	// behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to
	// specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction
	// history or their account statements. OANDA platforms always assume that
	// an Order’s trigger condition is set to the default value when indicating
	// the distance from an Order’s trigger price, and will always provide the
	// default trigger condition when creating or modifying an Order. A special
	// restriction applies when creating a GuaranteedStopLossOrder. In this
	// case the TriggerCondition value must either be “DEFAULT”, or the
	// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder
	// for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	// Default: OrderTriggerConditionDefault
	TriggerCondition *OrderTriggerCondition `json:"triggerCondition"`

	// The fee that will be charged if the TrailingStopLossOrder is filled
	// at the guaranteed price. The value is determined at Order creation time.
	// It is in price units and is charged for each unit of the Trade.
	GuaranteedExecutionPremium decimal.Decimal `json:"guaranteedExecutionPremium"`

	// The reason that the TrailingStopLossOrder was initiated
	Reason *TrailingStopLossOrderReason `json:"reason"`

	// Client Extensions to add to the Order (only provided if the Order is being created with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`

	// The ID of the OrderFill Transaction that caused this Order to be created
	// (only provided if this Order was created automatically when another Order
	// was filled).
	OrderFillTransactionID *TransactionID `json:"orderFillTransactionID"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID *TransactionID `json:"replacesOrderID"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`
}

func (TrailingStopLossOrderTransaction) GetType() TransactionType {
	return TransactionTypeTrailingStopLossOrder
}

// TrailingStopLossOrderRejectTransaction represents the rejection ot the creation of a TrailingStopLossOrder
type TrailingStopLossOrderRejectTransaction struct {
	TrailingStopLossOrderTransaction

	// The ID of the Order that this Order was intended to replace (only provided if this Order was intended to replace
	// an existing Order).
	IntendedReplacesOrderID *OrderID `json:"intendedReplacesOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (TrailingStopLossOrderRejectTransaction) GetType() TransactionType {
	return TransactionTypeTrailingStopLossOrderReject
}

// OrderFillTransaction represents the filling of an Order in the client's Account.
type OrderFillTransaction struct {
	TransactionBase

	// The ID of the Order filled.
	OrderID OrderID `json:"orderID"`

	// The client Order ID of the Order filled (only provided if the client has assigned one).
	ClientOrderID *ClientID `json:"clientOrderID"`

	// The name of the filled Order’s instrument.
	Instrument string `json:"instrument"`

	// The number of units filled by the OrderFill.
	Units decimal.Decimal `json:"units"`

	// The HomeConversionFactors in effect at the time of the OrderFill.
	HomeConversionFactors HomeConversionFactors `json:"homeConversionFactors"`

	// The price that all the units of the OrderFill should have been filled
	// at, in the absence of guaranteed price execution. This factors in the
	// Account’s current ClientPrice, used liquidity and the units of the
	// OrderFill only. If no Trades were closed with their price clamped for
	// guaranteed stop loss enforcement, then this value will match the price
	// fields of each Trade opened, closed, and reduced, and they will all be
	// the exact same.
	FullVWAP decimal.Decimal `json:"fullVWAP"`

	// The price in effect for the account at the time of the Order fill.
	FullPrice decimal.Decimal `json:"fullPrice"`

	// The reason that an Order was filled
	Reason OrderFillReason `json:"reason"`

	// The profit or loss incurred when the Order was filled.
	PL decimal.Decimal `json:"pl"`

	// The profit or loss incurred when the Order was filled, in the Instrument’s quote currency.
	QuotePL decimal.Decimal `json:"quotePL"`

	// The financing paid or collected when the Order was filled.
	Financing decimal.Decimal `json:"financing"`

	// The financing paid or collected when the Order was filled, in the Instrument’s base currency.
	BaseFinancing decimal.Decimal `json:"baseFinancing"`

	// The financing paid or collected when the Order was filled, in the Instrument’s quote currency.
	QuoteFinancing decimal.Decimal `json:"quoteFinancing"`

	// The commission charged in the Account’s home currency as a result of
	// filling the Order. The commission is always represented as a positive
	// quantity of the Account’s home currency, however it reduces the balance
	// in the Account.
	Commission decimal.Decimal `json:"commission"`

	// The total guaranteed execution fees charged for all Trades opened, closed or reduced with GuaranteedStopLossOrders.
	GuaranteedExecutionFee decimal.Decimal `json:"guaranteedExecutionFee"`

	// The total guaranteed execution fees charged for all Trades opened, closed
	// or reduced with GuaranteedStopLossOrders, expressed in the
	// Instrument’s quote currency.
	QuoteGuaranteedExecutionFee decimal.Decimal `json:"quoteGuaranteedExecutionFee"`

	// The Account’s balance after the Order was filled.
	AccountBalance decimal.Decimal `json:"accountBalance"`

	// The Trade that was opened when the Order was filled (only provided if filling the Order resulted in a new Trade).
	TradeOpened *TradeOpen `json:"tradeOpened"`

	// The Trades that were closed when the Order was filled (only provided if filling the Order resulted in a closing
	// open Trades).
	TradesClosed []TradeReduce `json:"tradesClosed"`

	// The Trade that was reduced when the Order was filled (only provided if filling the Order resulted in reducing an
	// open Trade).
	TradeReduced TradeReduce `json:"tradeReduced"`

	// The half spread cost for the OrderFill, which is the sum of the
	// halfSpreadCost values in the tradeOpened, tradesClosed and tradeReduced
	// fields. This can be a positive or negative value and is represented in
	// the home currency of the Account.
	HalfSpreadCost decimal.Decimal `json:"halfSpreadCost"`
}

func (OrderFillTransaction) GetType() TransactionType {
	return TransactionTypeOrderFill
}

// OrderCancelTransaction represents the cancellation of an Order in the client's Account
type OrderCancelTransaction struct {
	TransactionBase

	// The ID of the Order cancelled
	OrderID OrderID `json:"orderID"`

	// The client ID of the Order cancelled (only provided if the Order has a client Order ID).
	ClientOrderID ClientID `json:"clientOrderID"`

	// The reason that the Order was cancelled.
	Reason OrderCancelReason `json:"reason"`

	// The ID of the Order that replaced this Order (only provided if this Order was cancelled for replacement).
	ReplacedByOrderID *Order `json:"replacedByOrderID"`
}

func (OrderCancelTransaction) GetType() TransactionType {
	return TransactionTypeOrderCancel
}

// OrderCancelRejectTransaction represents the rejection of the cancellation of an Order in the client's Account.
type OrderCancelRejectTransaction struct {
	TransactionBase

	// The ID of the Order cancelled
	OrderID OrderID `json:"orderID"`

	// The client ID of the Order cancelled (only provided if the Order has a client Order ID).
	ClientOrderID ClientID `json:"clientOrderID"`

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (OrderCancelRejectTransaction) GetType() TransactionType {
	return TransactionTypeOrderCancelReject
}

// OrderClientExtensionsModifyTransaction represents the modification of the Order's ClientExtensions
type OrderClientExtensionsModifyTransaction struct {
	TransactionBase

	// The ID of the Order cancelled
	OrderID OrderID `json:"orderID"`

	// The client ID of the Order cancelled (only provided if the Order has a client Order ID).
	ClientOrderID ClientID `json:"clientOrderID"`

	// The new Client Extensions for the Order.
	ClientExtensionsModify ClientExtensions `json:"clientExtensionsModify"`

	// The new Client Extensions for the Order’s Trade on fill.
	TradeExtensionsModify ClientExtensions `json:"tradeExtensionsModify"`
}

func (OrderClientExtensionsModifyTransaction) GetType() TransactionType {
	return TransactionTypeOrderClientExtensionsModify
}

// OrderClientExtensionsModifyRejectTransaction represents the rejection of the modification of an Order's
// ClientExtensions
type OrderClientExtensionsModifyRejectTransaction struct {
	OrderClientExtensionsModifyTransaction

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (OrderClientExtensionsModifyRejectTransaction) GetType() TransactionType {
	return TransactionTypeOrderClientExtensionsModifyReject
}

// TradeClientExtensionsModifyTransaction represents the modification of a Trade's ClientExtensions
type TradeClientExtensionsModifyTransaction struct {
	TransactionBase

	// The ID of the Trade whose client extensions are to be modified.
	TradeID TradeID `json:"tradeID"`

	// The original Client ID of the Trade whose client extensions are to be modified.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The new Client Extensions for the Trade.
	TradeClientExtensionsModify ClientExtensions `json:"tradeClientExtensionsModify"`
}

func (TradeClientExtensionsModifyTransaction) GetType() TransactionType {
	return TransactionTypeTradeClientExtensionsModify
}

// TradeClientExtensionsModifyRejectTransaction represents the rejection of the modification of a Trade's
// ClientExtensions
type TradeClientExtensionsModifyRejectTransaction struct {
	TradeClientExtensionsModifyTransaction

	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

func (TradeClientExtensionsModifyRejectTransaction) GetType() TransactionType {
	return TransactionTypeTradeClientExtensionsModifyReject
}

// MarginCallEnterTransaction is created when an Account enters the margin call state.
type MarginCallEnterTransaction struct {
	TransactionBase
}

func (MarginCallEnterTransaction) GetType() TransactionType {
	return TransactionTypeMarginCallEnter
}

// MarginCallExtendTransaction is created when the margin call state for an Account has been extended
type MarginCallExtendTransaction struct {
	TransactionBase

	// The number of the extensions to the Account’s current margin call that
	// have been applied. This value will be set to 1 for the first
	// MarginCallExtendTransaction
	ExtensionNumber int `json:"extensionNumber"`
}

func (MarginCallExtendTransaction) GetType() TransactionType {
	return TransactionTypeMarginCallExtend
}

// MarginCallExitTransaction is created when an Account leaves the margin call state
type MarginCallExitTransaction struct {
	TransactionBase
}

func (MarginCallExitTransaction) GetType() TransactionType {
	return TransactionTypeMarginCallExit
}

// DelayedTradeClosureTransaction is created administratively to indicate open trades that should have been closed but
// weren't because the open trades' instruments were untradeable at the time. Open trades listed in this transaction
// will be closed once their respective instruments become available.
type DelayedTradeClosureTransaction struct {
	TransactionBase

	// The reason for the delayed trade closure
	Reason MarketOrderReason `json:"reason"`

	// List of Trade ID’s identifying the open trades that will be closed when their respective instruments become tradeable
	TradeIDs []TradeID `json:"tradeIDs"`
}

func (DelayedTradeClosureTransaction) GetType() TransactionType {
	return TransactionTypeDelayedTradeClosure
}

// DailyFinancingTransaction represents the daily payment/collection of financing for an Account.
type DailyFinancingTransaction struct {
	TransactionBase

	// The amount of financing paid/collected for the Account.
	Financing decimal.Decimal `json:"financing"`

	// The Account’s balance after daily financing.
	AccountBalance decimal.Decimal `json:"accountBalance"`

	// The financing paid/collected for each Position in the Account.
	PositionFinancings []PositionFinancing `json:"positionFinancings"`
}

func (DailyFinancingTransaction) GetType() TransactionType {
	return TransactionTypeDailyFinancing
}

// DividendAdjustmentTransaction is created administratively to pay or collect dividend adjustment amounts to or from an
// Account.
type DividendAdjustmentTransaction struct {
	TransactionBase

	// The name of the instrument for the dividendAdjustment transaction
	Instrument string `json:"instrument"`

	// The total dividend adjustment amount paid or collected in the Account’s
	// home currency for the Account as a result of applying the
	// DividendAdjustment Transaction. This is the sum of the dividend
	// adjustments paid/collected for each OpenTradeDividendAdjustment found
	// within the Transaction.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The total dividend adjustment amount paid or collected in the
	// Instrument’s quote currency for the Account as a result of applying the
	// DividendAdjustment Transaction. This is the sum of the quote dividend
	// adjustments paid/collected for each OpenTradeDividendAdjustment found
	// within the Transaction.
	QuoteDividendAdjustment decimal.Decimal `json:"quoteDividendAdjustment"`

	// The HomeConversionFactors in effect at the time of the DividendAdjustment.
	HomeConversionFactors HomeConversionFactors `json:"homeConversionFactors"`

	// The Account balance after applying the DividendAdjustment Transaction
	AccountBalance decimal.Decimal `json:"accountBalance"`

	// The dividend adjustment payment/collection details for each open Trade,
	// within the Account, for which a dividend adjustment is to be paid or
	// collected.
	OpenTradeDividendAdjustments []OpenTradeDividendAdjustment `json:"openTradeDividendAdjustments"`
}

func (DividendAdjustmentTransaction) GetType() TransactionType {
	return TransactionTypeDividendAdjustment
}

// ResetResettablePLTransaction represents the resetting of the Account's resettable PL counters
type ResetResettablePLTransaction struct {
	TransactionBase
}

func (ResetResettablePLTransaction) GetType() TransactionType {
	return TransactionTypeResetResettablePl
}

// TransactionID is the unique Transaction identifier within each Account.
// Format: String representation of the numerical OANDA-assigned TransactionID
// Example: 1523
type TransactionID string

// TransactionType covers the possible types of a Transaction
type TransactionType string

const (
	TransactionTypeCreate                            = TransactionType("CREATE")
	TransactionTypeClose                             = TransactionType("CLOSE")
	TransactionTypeReopen                            = TransactionType("REOPEN")
	TransactionTypeClientConfigure                   = TransactionType("CLIENT_CONFIGURE")
	TransactionTypeClientConfigureReject             = TransactionType("CLIENT_CONFIGURE_REJECT")
	TransactionTypeTransferFunds                     = TransactionType("TRANSFER_FUNDS")
	TransactionTypeTransferFundsReject               = TransactionType("TRANSFER_FUNDS_REJECT")
	TransactionTypeMarketOrder                       = TransactionType("MARKET_ORDER")
	TransactionTypeMarketOrderReject                 = TransactionType("MARKET_ORDER_REJECT")
	TransactionTypeFixedPriceOrder                   = TransactionType("FIXED_PRICE_ORDER")
	TransactionTypeLimitOrder                        = TransactionType("LIMIT_ORDER")
	TransactionTypeLimitOrderReject                  = TransactionType("LIMIT_ORDER_REJECT")
	TransactionTypeStopOrder                         = TransactionType("STOP_ORDER")
	TransactionTypeStopOrderReject                   = TransactionType("STOP_ORDER_REJECT")
	TransactionTypeMarketIfTouchedOrder              = TransactionType("MARKET_IF_TOUCHED_ORDER")
	TransactionTypeMarketIfTouchedOrderReject        = TransactionType("MARKET_IF_TOUCHED_ORDER_REJECT")
	TransactionTypeTakeProfitOrder                   = TransactionType("TAKE_PROFIT_ORDER")
	TransactionTypeTakeProfitOrderReject             = TransactionType("TAKE_PROFIT_ORDER_REJECT")
	TransactionTypeStopLossOrder                     = TransactionType("STOP_LOSS_ORDER")
	TransactionTypeStopLossOrderReject               = TransactionType("STOP_LOSS_ORDER_REJECT")
	TransactionTypeGuaranteedStopLossOrder           = TransactionType("GUARANTEED_STOP_LOSS_ORDER")
	TransactionTypeGuaranteedStopLossOrderReject     = TransactionType("GUARANTEED_STOP_LOSS_ORDER_REJECT")
	TransactionTypeTrailingStopLossOrder             = TransactionType("TRAILING_STOP_LOSS_ORDER")
	TransactionTypeTrailingStopLossOrderReject       = TransactionType("TRAILING_STOP_LOSS_ORDER_REJECT")
	TransactionTypeOrderFill                         = TransactionType("ORDER_FILL")
	TransactionTypeOrderCancel                       = TransactionType("ORDER_CANCEL")
	TransactionTypeOrderCancelReject                 = TransactionType("ORDER_CANCEL_REJECT")
	TransactionTypeOrderClientExtensionsModify       = TransactionType("ORDER_CLIENT_EXTENSIONS_MODIFY")
	TransactionTypeOrderClientExtensionsModifyReject = TransactionType("ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT")
	TransactionTypeTradeClientExtensionsModify       = TransactionType("TRADE_CLIENT_EXTENSIONS_MODIFY")
	TransactionTypeTradeClientExtensionsModifyReject = TransactionType("TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT")
	TransactionTypeMarginCallEnter                   = TransactionType("MARGIN_CALL_ENTER")
	TransactionTypeMarginCallExtend                  = TransactionType("MARGIN_CALL_EXTEND")
	TransactionTypeMarginCallExit                    = TransactionType("MARGIN_CALL_EXIT")
	TransactionTypeDelayedTradeClosure               = TransactionType("DELAYED_TRADE_CLOSURE")
	TransactionTypeDailyFinancing                    = TransactionType("DAILY_FINANCING")
	TransactionTypeDividendAdjustment                = TransactionType("DIVIDEND_ADJUSTMENT")
	TransactionTypeResetResettablePl                 = TransactionType("RESET_RESETTABLE_PL")
)

// FundingReason is the reason that an Account is being funded
type FundingReason string

const (
	// FundingReasonClientFunding means the client has initiated a funds transfer
	FundingReasonClientFunding = FundingReason("CLIENT_FUNDING")

	// FundingReasonAccountTransfer means funds are being transferred between two Accounts.
	FundingReasonAccountTransfer = FundingReason("ACCOUNT_TRANSFER")

	// FundingReasonDivisionMigration means funds are being transferred as part of a Division migration
	FundingReasonDivisionMigration = FundingReason("DIVISION_MIGRATION")

	// FundingReasonSiteMigration means funds are being transferred as part of a Site migration
	FundingReasonSiteMigration = FundingReason("SITE_MIGRATION")

	// FundingReasonAdjustment means funds are being transferred as part of an Account adjustment
	FundingReasonAdjustment = FundingReason("ADJUSTMENT")
)

// MarketOrderReason specifies the reason that the MarketOrder was created
type MarketOrderReason string

const (
	// MarketOrderReasonClientOrder means the MarketOrder was created at the request of a client
	MarketOrderReasonClientOrder = MarketOrderReason("CLIENT_ORDER")

	// MarketOrderReasonTradeClose means the MarketOrder was created to close a Trade at the request of a client
	MarketOrderReasonTradeClose = MarketOrderReason("TRADE_CLOSE")

	// MarketOrderReasonPositionCloseout means the MarketOrder was created to close a Position at the request of a
	// client
	MarketOrderReasonPositionCloseout = MarketOrderReason("POSITION_CLOSEOUT")

	// MarketOrderReasonMarginCloseout means the MarketOrder was created as part of a Margin Closeout
	MarketOrderReasonMarginCloseout = MarketOrderReason("MARGIN_CLOSEOUT")

	// MarketOrderReasonDelayedTradeClose means the MarketOrder was created to close a trade marked for delayed closure
	MarketOrderReasonDelayedTradeClose = MarketOrderReason("DELAYED_TRADE_CLOSE")
)

// FixedPriceOrderReason specifies the reason that the FixedPriceOrder was created
type FixedPriceOrderReason string

const (
	// FixedPriceOrderReasonPlatformAccountMigration means the FixedPriceOrder was created as part of a platform
	// account migration
	FixedPriceOrderReasonPlatformAccountMigration = FixedPriceOrderReason("PLATFORM_ACCOUNT_MIGRATION")

	// FixedPriceOrderReasonTradeCloseDivisionAccountMigration means the FixedPriceOrder was created to close a Trade as
	// part of division account migration
	FixedPriceOrderReasonTradeCloseDivisionAccountMigration = FixedPriceOrderReason("TRADE_CLOSE_DIVISION_ACCOUNT_MIGRATION")

	// FixedPriceOrderReasonTradeCloseAdministrativeAction means the FixedPriceOrder was created to close a Trade
	// administratively
	FixedPriceOrderReasonTradeCloseAdministrativeAction = FixedPriceOrderReason("TRADE_CLOSE_ADMINISTRATIVE_ACTION")
)

// LimitOrderReason specifies the reason that the LimitOrder was initiated
type LimitOrderReason string

const (
	// LimitOrderReasonClientOrder means the LimitOrder was initiated at the request of a client
	LimitOrderReasonClientOrder = LimitOrderReason("CLIENT_ORDER")

	// LimitOrderReasonReplacement means the LimitOrder was initiated as a replacement for an existing Order
	LimitOrderReasonReplacement = LimitOrderReason("REPLACEMENT")
)

// StopOrderReason specifies the reason that the StopOrder was initiated
type StopOrderReason string

const (
	// StopOrderReasonClientOrder means the StopOrder was initiated at the request of a client
	StopOrderReasonClientOrder = StopOrderReason("CLIENT_ORDER")

	// StopOrderReasonReplacement means the StopOrder was initiated as a replacement for an existing Order
	StopOrderReasonReplacement = StopOrderReason("REPLACEMENT")
)

// MarketIfTouchedOrderReason specifies the reason that the MarketIfTouchedOrder was initiated
type MarketIfTouchedOrderReason string

const (
	// MarketIfTouchedOrderReasonClientOrder means the MarketIfTouchedOrder was initiated at the request of a client
	MarketIfTouchedOrderReasonClientOrder = MarketIfTouchedOrderReason("CLIENT_ORDER")

	// MarketIfTouchedOrderReasonReplacement means the MarketIfTouchedOrder was initiated as a replacement for an
	// existing Order
	MarketIfTouchedOrderReasonReplacement = MarketIfTouchedOrderReason("REPLACEMENT")
)

// TakeProfitOrderReason specifies the reason that the TakeProfitOrder was initiated
type TakeProfitOrderReason string

const (
	// TakeProfitOrderReasonClientOrder means the TakeProfitOrder was initiated at the request of a client
	TakeProfitOrderReasonClientOrder = TakeProfitOrderReason("CLIENT_ORDER")

	// TakeProfitOrderReasonReplacement means the TakeProfitOrder was initiated as a replacement for an existing Order
	TakeProfitOrderReasonReplacement = TakeProfitOrderReason("REPLACEMENT")

	// TakeProfitOrderReasonOnFill means the TakeProfitOrder was initiated automatically when an Order was filled that
	// opened a new Trade requiring a TakeProfitOrder
	TakeProfitOrderReasonOnFill = TakeProfitOrderReason("ON_FILL")
)

// StopLossOrderReason specifies the reason that the StopLossOrder was initiated
type StopLossOrderReason string

const (
	// StopLossOrderReasonClientOrder means the StopLossOrder was initiated at the request of a client
	StopLossOrderReasonClientOrder = StopLossOrderReason("CLIENT_ORDER")

	// StopLossOrderReasonReplacement means the StopLossOrder was initiated as a replacement for an existing Order
	StopLossOrderReasonReplacement = StopLossOrderReason("REPLACEMENT")

	// StopLossOrderReasonOnFill means the StopLossOrder was initiated automatically when an Order was filled that
	// opened a new Trade requiring a StopLossOrder
	StopLossOrderReasonOnFill = StopLossOrderReason("ON_FILL")
)

// GuaranteedStopLossOrderReason specifies the reason that the GuaranteedStopLossOrder was initiated
type GuaranteedStopLossOrderReason string

const (
	// GuaranteedStopLossOrderReasonClientOrder means the GuaranteedStopLossOrder was initiated at the request of a client
	GuaranteedStopLossOrderReasonClientOrder = GuaranteedStopLossOrderReason("CLIENT_ORDER")

	// GuaranteedStopLossOrderReasonReplacement means the GuaranteedStopLossOrder was initiated as a replacement for an existing Order
	GuaranteedStopLossOrderReasonReplacement = GuaranteedStopLossOrderReason("REPLACEMENT")

	// GuaranteedStopLossOrderReasonOnFill means the GuaranteedStopLossOrder was initiated automatically when an Order was filled that
	// opened a new Trade requiring a GuaranteedStopLossOrder
	GuaranteedStopLossOrderReasonOnFill = GuaranteedStopLossOrderReason("ON_FILL")
)

// TrailingStopLossOrderReason specifies the reason that the TrailingStopLossOrder was initiated
type TrailingStopLossOrderReason string

const (
	// TrailingStopLossOrderReasonClientOrder means the TrailingStopLossOrder was initiated at the request of a client
	TrailingStopLossOrderReasonClientOrder = TrailingStopLossOrderReason("CLIENT_ORDER")

	// TrailingStopLossOrderReasonReplacement means the TrailingStopLossOrder was initiated as a replacement for an existing Order
	TrailingStopLossOrderReasonReplacement = TrailingStopLossOrderReason("REPLACEMENT")

	// TrailingStopLossOrderReasonOnFill means the TrailingStopLossOrder was initiated automatically when an Order was filled that
	// opened a new Trade requiring a TrailingStopLossOrder
	TrailingStopLossOrderReasonOnFill = TrailingStopLossOrderReason("ON_FILL")
)

// OrderFillReason specifies the reason that an Order was filled
type OrderFillReason string

const (
	// OrderFillReasonLimitOrder means the Order filled was a LimitOrder
	OrderFillReasonLimitOrder = OrderFillReason("LIMIT_ORDER")

	// OrderFillReasonStopOrder means the Order filled was a StopOrder
	OrderFillReasonStopOrder = OrderFillReason("STOP_ORDER")

	// OrderFillReasonMarketIfTouchedOrder means the Order filled was a MarketIfTouchedOrder
	OrderFillReasonMarketIfTouchedOrder = OrderFillReason("MARKET_IF_TOUCHED_ORDER")

	// OrderFillReasonTakeProfitOrder means the Order filled was a TakeProfitOrder
	OrderFillReasonTakeProfitOrder = OrderFillReason("TAKE_PROFIT_ORDER")

	// OrderFillReasonStopLossOrder means the Order filled was a StopLossOrder
	OrderFillReasonStopLossOrder = OrderFillReason("STOP_LOSS_ORDER")

	// OrderFillReasonGuaranteedStopLossOrder means the Order filled was a GuaranteedStopLossOrder
	OrderFillReasonGuaranteedStopLossOrder = OrderFillReason("GUARANTEED_STOP_LOSS_ORDER")

	// OrderFillReasonTrailingStopLossOrder means the Order filled was a TrailingStopLossOrder
	OrderFillReasonTrailingStopLossOrder = OrderFillReason("TRAILING_STOP_LOSS_ORDER")

	// OrderFillReasonMarketOrder means the Order filled was a MarketOrder
	OrderFillReasonMarketOrder = OrderFillReason("MARKET_ORDER")

	// OrderFillReasonMarketOrderTradeClose means the Order filled was a MarketOrder used to explicitly close a Trade
	OrderFillReasonMarketOrderTradeClose = OrderFillReason("MARKET_ORDER_TRADE_CLOSE")

	// OrderFillReasonMarketOrderPositionCloseout means the Order filled was a MarketOrder used to explicitly close a Position
	OrderFillReasonMarketOrderPositionCloseout = OrderFillReason("MARKET_ORDER_POSITION_CLOSEOUT")

	// OrderFillReasonMarketOrderMarginCloseout means the Order filled was a MarketOrder used for a Margin Closeout
	OrderFillReasonMarketOrderMarginCloseout = OrderFillReason("MARKET_ORDER_MARGIN_CLOSEOUT")

	// OrderFillReasonMarketOrderDelayedTradeClose means the Order filled was a MarketOrder used for a delayed Trade close
	OrderFillReasonMarketOrderDelayedTradeClose = OrderFillReason("MARKET_ORDER_DELAYED_TRADE_CLOSE")

	// OrderFillReasonFixedPriceOrder means the Order filled was a FixedPriceOrder
	OrderFillReasonFixedPriceOrder = OrderFillReason("FIXED_PRICE_ORDER")

	// OrderFillReasonFixedPriceOrderPlatformAccountMigration means the Order filled was a FixedPriceOrder created as part of a platform account migration
	OrderFillReasonFixedPriceOrderPlatformAccountMigration = OrderFillReason("FIXED_PRICE_ORDER_PLATFORM_ACCOUNT_MIGRATION")

	// OrderFillReasonFixedPriceOrderDivisionAccountMigration means the Order filled was a FixedPriceOrder created to close a Trade as part of division account migration
	OrderFillReasonFixedPriceOrderDivisionAccountMigration = OrderFillReason("FIXED_PRICE_ORDER_DIVISION_ACCOUNT_MIGRATION")

	// OrderFillReasonFixedPriceOrderAdministrativeAction means the Order filled was a FixedPriceOrder created to close a Trade administratively
	OrderFillReasonFixedPriceOrderAdministrativeAction = OrderFillReason("FIXED_PRICE_ORDER_ADMINISTRATIVE_ACTION")
)

// OrderCancelReason specifies the reason that an Order was cancelled
type OrderCancelReason string

const (
	// OrderCancelReasonInternalServerError means the Order was cancelled because at the time of filling, an unexpected
	// internal server error occurred.
	OrderCancelReasonInternalServerError = OrderCancelReason("INTERNAL_SERVER_ERROR")

	// OrderCancelReasonAccountLocked means the Order was cancelled because at the time of filling the account was
	// locked.
	OrderCancelReasonAccountLocked = OrderCancelReason("ACCOUNT_LOCKED")

	// OrderCancelReasonAccountNewPositionsLocked means the order was to be filled, however the account is configured to
	// not allow new positions to be created.
	OrderCancelReasonAccountNewPositionsLocked = OrderCancelReason("ACCOUNT_NEW_POSITIONS_LOCKED")

	// OrderCancelReasonAccountOrderCreationLocked means filling the Order wasn’t possible because it required the
	// creation of a dependent Order and the Account is locked for Order creation.
	OrderCancelReasonAccountOrderCreationLocked = OrderCancelReason("ACCOUNT_ORDER_CREATION_LOCKED")

	// OrderCancelReasonAccountOrderFillLocked means filling the Order was not possible because the Account is locked
	// for filling Orders.
	OrderCancelReasonAccountOrderFillLocked = OrderCancelReason("ACCOUNT_ORDER_FILL_LOCKED")

	// OrderCancelReasonClientRequest means the Order was cancelled explicitly at the request of the client.
	OrderCancelReasonClientRequest = OrderCancelReason("CLIENT_REQUEST")

	// OrderCancelReasonMigration means the Order cancelled because it is being migrated to another account.
	OrderCancelReasonMigration = OrderCancelReason("MIGRATION")

	// OrderCancelReasonMarketHalted means filling the Order wasn’t possible because the Order’s instrument was halted.
	OrderCancelReasonMarketHalted = OrderCancelReason("MARKET_HALTED")

	// OrderCancelReasonLinkedTradeClosed means the Order is linked to an open Trade that was closed.
	OrderCancelReasonLinkedTradeClosed = OrderCancelReason("LINKED_TRADE_CLOSED")

	// OrderCancelReasonTimeInForceExpired means the time in force specified for this order has passed.
	OrderCancelReasonTimeInForceExpired = OrderCancelReason("TIME_IN_FORCE_EXPIRED")

	// OrderCancelReasonInsufficientMargin means filling the Order wasn’t possible because the Account had insufficient
	// margin.
	OrderCancelReasonInsufficientMargin = OrderCancelReason("INSUFFICIENT_MARGIN")

	// OrderCancelReasonFifoViolation means filling the Order would have resulted in a FIFO violation.
	OrderCancelReasonFifoViolation = OrderCancelReason("FIFO_VIOLATION")

	// OrderCancelReasonBoundsViolation means filling the Order would have violated the Order’s price bound.
	OrderCancelReasonBoundsViolation = OrderCancelReason("BOUNDS_VIOLATION")

	// OrderCancelReasonClientRequestReplaced means the Order was cancelled for replacement at the request of the client.
	OrderCancelReasonClientRequestReplaced = OrderCancelReason("CLIENT_REQUEST_REPLACED")

	// OrderCancelReasonDividendAdjustmentReplaced means the Order was cancelled for replacement with an adjusted
	// fillPrice to accommodate for the price movement caused by a dividendAdjustment.
	OrderCancelReasonDividendAdjustmentReplaced = OrderCancelReason("DIVIDEND_ADJUSTMENT_REPLACED")

	// OrderCancelReasonInsufficientLiquidity means filling the Order wasn’t possible because enough liquidity available.
	OrderCancelReasonInsufficientLiquidity = OrderCancelReason("INSUFFICIENT_LIQUIDITY")

	// OrderCancelReasonTakeProfitOnFillGtdTimestampInPast means filling the Order would have resulted in the creation
	// of a TakeProfitOrder with a GTD time in the past.
	OrderCancelReasonTakeProfitOnFillGtdTimestampInPast = OrderCancelReason("TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST")

	// OrderCancelReasonTakeProfitOnFillLoss means filling the Order would result in the creation of a TakeProfitOrder
	// that would have been filled immediately, closing the new Trade at a loss.
	OrderCancelReasonTakeProfitOnFillLoss = OrderCancelReason("TAKE_PROFIT_ON_FILL_LOSS")

	// OrderCancelReasonLosingTakeProfit means filling the Order would result in the creation of a TakeProfitLossOrder
	// that would close the new Trade at a loss when filled.
	OrderCancelReasonLosingTakeProfit = OrderCancelReason("LOSING_TAKE_PROFIT")

	// OrderCancelReasonStopLossOnFillGtdTimestampInPast means filling the Order would have resulted in the creation of
	// a StopLossOrder with a GTD time in the past.
	OrderCancelReasonStopLossOnFillGtdTimestampInPast = OrderCancelReason("STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST")

	// OrderCancelReasonStopLossOnFillLoss means filling the Order would result in the creation of a StopLossOrder that
	// would have been filled immediately, closing the new Trade at a loss.
	OrderCancelReasonStopLossOnFillLoss = OrderCancelReason("STOP_LOSS_ON_FILL_LOSS")

	// OrderCancelReasonStopLossOnFillPriceDistanceMaximumExceeded means filling the Order would result in the creation
	// of a StopLossOrder whose price would be zero or negative due to the specified distance.
	OrderCancelReasonStopLossOnFillPriceDistanceMaximumExceeded = OrderCancelReason("STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED")

	// OrderCancelReasonStopLossOnFillRequired means filling the Order would not result in the creation of StopLossOrder,
	// however the Account’s configuration requires that all Trades have a StopLossOrder attached to them.
	OrderCancelReasonStopLossOnFillRequired = OrderCancelReason("STOP_LOSS_ON_FILL_REQUIRED")

	// OrderCancelReasonStopLossOnFillGuaranteedRequired means filling the Order would not result in the creation of a
	// GuaranteedStopLossOrder, however the Account’s configuration requires that all Trades have a
	// GuaranteedStopLossOrder attached to them.
	OrderCancelReasonStopLossOnFillGuaranteedRequired = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED")

	// OrderCancelReasonStopLossOnFillGuaranteedNotAllowed means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder, however the Account’s configuration does not allow GuaranteedStopLossOrders.
	OrderCancelReasonStopLossOnFillGuaranteedNotAllowed = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED")

	// OrderCancelReasonStopLossOnFillGuaranteedMinimumDistanceNotMet means filling the Order would result in the
	// creation of a GuaranteedStopLossOrder with a distance smaller than the configured minimum distance.
	OrderCancelReasonStopLossOnFillGuaranteedMinimumDistanceNotMet = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET")

	// OrderCancelReasonStopLossOnFillGuaranteedLevelRestrictionExceeded means filling the Order would result in the
	// creation of a GuaranteedStopLossOrder with trigger price and number of units that that violates the account’s
	// GuaranteedStopLossOrderLevelRestriction.
	OrderCancelReasonStopLossOnFillGuaranteedLevelRestrictionExceeded = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED")

	// OrderCancelReasonStopLossOnFillGuaranteedHedgingNotAllowed means filling the Order would result in the creation
	// of a GuaranteedStopLossOrder for a hedged Trade, however the Account’s configuration does not allow
	// GuaranteedStopLossOrders for hedged Trades/Positions.
	OrderCancelReasonStopLossOnFillGuaranteedHedgingNotAllowed = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED")

	// OrderCancelReasonStopLossOnFillTimeInForceInvalid means filling the Order would result in the creation of a
	// StopLossOrder whose TimeInForce value is invalid. A likely cause would be if the Account requires
	// GuaranteedStopLossOrders and the TimeInForce value were not GTC.
	OrderCancelReasonStopLossOnFillTimeInForceInvalid = OrderCancelReason("STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID")

	// OrderCancelReasonStopLossOnFillTriggerConditionInvalid means filling the Order would result in the creation of a
	// StopLossOrder whose TriggerCondition value is invalid. A likely cause would be if the stop loss order is
	// guaranteed and the TimeInForce is not TRIGGER_DEFAULT or TRIGGER_BID for a long trade, or not TRIGGER_DEFAULT or
	// TRIGGER_ASK for a short trade.
	OrderCancelReasonStopLossOnFillTriggerConditionInvalid = OrderCancelReason("STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID")

	// OrderCancelReasonGuaranteedStopLossOnFillGtdTimestampInPast means filling the Order would have resulted in the
	// creation of a GuaranteedStopLossOrder with a GTD time in the past.
	OrderCancelReasonGuaranteedStopLossOnFillGtdTimestampInPast = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST")

	// OrderCancelReasonGuaranteedStopLossOnFillLoss means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder that would have been filled immediately, closing the new Trade at a loss.
	OrderCancelReasonGuaranteedStopLossOnFillLoss = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_LOSS")

	// OrderCancelReasonGuaranteedStopLossOnFillPriceDistanceMaximumExceeded means filling the Order would result in the
	// creation of a GuaranteedStopLossOrder whose price would be zero or negative due to the specified distance.
	OrderCancelReasonGuaranteedStopLossOnFillPriceDistanceMaximumExceeded = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED")

	// OrderCancelReasonGuaranteedStopLossOnFillRequired means filling the Order would not result in the creation of a
	// GuaranteedStopLossOrder, however the Account’s configuration requires that all Trades have a
	// GuaranteedStopLossOrder attached to them.
	OrderCancelReasonGuaranteedStopLossOnFillRequired = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED")

	// OrderCancelReasonGuaranteedStopLossOnFillNotAllowed means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder, however the Account’s configuration does not allow GuaranteedStopLossOrders.
	OrderCancelReasonGuaranteedStopLossOnFillNotAllowed = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED")

	// OrderCancelReasonGuaranteedStopLossOnFillMinimumDistanceNotMet means filling the Order would result in the
	// creation of a GuaranteedStopLossOrder with a distance smaller than the configured minimum distance.
	OrderCancelReasonGuaranteedStopLossOnFillMinimumDistanceNotMet = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET")

	// OrderCancelReasonGuaranteedStopLossOnFillLevelRestrictionVolumeExceeded means filling the Order would result in
	// the creation of a GuaranteedStopLossOrder with trigger number of units that violates the account’s
	// GuaranteedStopLossOrderLevelRestriction volume.
	OrderCancelReasonGuaranteedStopLossOnFillLevelRestrictionVolumeExceeded = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED")

	// OrderCancelReasonGuaranteedStopLossOnFillLevelRestrictionPriceRangeExceeded means filling the Order would result
	// in the creation of a GuaranteedStopLossOrder with trigger price that violates the account’s
	// GuaranteedStopLossOrder level restriction price range.
	OrderCancelReasonGuaranteedStopLossOnFillLevelRestrictionPriceRangeExceeded = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED")

	// OrderCancelReasonGuaranteedStopLossOnFillHedgingNotAllowed means filling the Order would result in the creation
	// of a GuaranteedStopLossOrder for a hedged Trade, however the Account’s configuration does not allow
	// GuaranteedStopLossOrders for hedged Trades/Positions.
	OrderCancelReasonGuaranteedStopLossOnFillHedgingNotAllowed = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_HEDGING_NOT_ALLOWED")

	// OrderCancelReasonGuaranteedStopLossOnFillTimeInForceInvalid means filling the Order would result in the creation
	// of a GuaranteedStopLossOrder whose TimeInForce value is invalid. A likely cause would be if the Account requires
	// guaranteed stop loss orders and the TimeInForce value were not GTC.
	OrderCancelReasonGuaranteedStopLossOnFillTimeInForceInvalid = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID")

	// OrderCancelReasonGuaranteedStopLossOnFillTriggerConditionInvalid means filling the Order would result in the
	// creation of a GuaranteedStopLossOrder whose TriggerCondition value is invalid. A likely cause would be the
	// TimeInForce is not TRIGGER_DEFAULT or TRIGGER_BID for a long trade, or not TRIGGER_DEFAULT or TRIGGER_ASK for a
	// short trade.
	OrderCancelReasonGuaranteedStopLossOnFillTriggerConditionInvalid = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID")

	// OrderCancelReasonTakeProfitOnFillPriceDistanceMaximumExceeded means filling the Order would result in the
	// creation of a TakeProfitOrder whose price would be zero or negative due to the specified distance.
	OrderCancelReasonTakeProfitOnFillPriceDistanceMaximumExceeded = OrderCancelReason("TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED")

	// OrderCancelReasonTrailingStopLossOnFillGtdTimestampInPast means filling the Order would have resulted in the
	// creation of a TrailingStopLossOrder with a GTD time in the past.
	OrderCancelReasonTrailingStopLossOnFillGtdTimestampInPast = OrderCancelReason("TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST")

	// OrderCancelReasonClientTradeIdAlreadyExists means filling the Order would result in the creation of a new Open
	// Trade with a client Trade ID already in use.
	OrderCancelReasonClientTradeIdAlreadyExists = OrderCancelReason("CLIENT_TRADE_ID_ALREADY_EXISTS")

	// OrderCancelReasonPositionCloseoutFailed means closing out a position wasn’t fully possible.
	OrderCancelReasonPositionCloseoutFailed = OrderCancelReason("POSITION_CLOSEOUT_FAILED")

	// OrderCancelReasonOpenTradesAllowedExceeded means filling the Order would cause the maximum open trades allowed
	// for the Account to be exceeded.
	OrderCancelReasonOpenTradesAllowedExceeded = OrderCancelReason("OPEN_TRADES_ALLOWED_EXCEEDED")

	// OrderCancelReasonPendingOrdersAllowedExceeded means filling the Order would have resulted in exceeding the number
	// of pending Orders allowed for the Account.
	OrderCancelReasonPendingOrdersAllowedExceeded = OrderCancelReason("PENDING_ORDERS_ALLOWED_EXCEEDED")

	// OrderCancelReasonTakeProfitOnFillClientOrderIdAlreadyExists means filling the Order would have resulted in the
	// creation of a TakeProfitOrder with a ClientOrderID that is already in use.
	OrderCancelReasonTakeProfitOnFillClientOrderIdAlreadyExists = OrderCancelReason("TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS")

	// OrderCancelReasonStopLossOnFillClientOrderIdAlreadyExists means filling the Order would have resulted in the
	// creation of a StopLossOrder with a ClientOrderID that is already in use.
	OrderCancelReasonStopLossOnFillClientOrderIdAlreadyExists = OrderCancelReason("STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS")

	// OrderCancelReasonGuaranteedStopLossOnFillClientOrderIdAlreadyExists means filling the Order would have resulted
	// in the creation of a GuaranteedStopLossOrder with a ClientOrderID that is already in use.
	OrderCancelReasonGuaranteedStopLossOnFillClientOrderIdAlreadyExists = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS")

	// OrderCancelReasonTrailingStopLossOnFillClientOrderIdAlreadyExists means filling the Order would have resulted in
	// the creation of a TrailingStopLossOrder with a ClientOrderID that is already in use.
	OrderCancelReasonTrailingStopLossOnFillClientOrderIdAlreadyExists = OrderCancelReason("TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS")

	// OrderCancelReasonPositionSizeExceeded means filling the Order would have resulted in the Account’s maximum
	// position size limit being exceeded for the Order’s instrument.
	OrderCancelReasonPositionSizeExceeded = OrderCancelReason("POSITION_SIZE_EXCEEDED")

	// OrderCancelReasonHedgingGsloViolation means filling the Order would result in the creation of a Trade, however
	// there already exists an opposing (hedged) Trade that has a GuaranteedStopLossOrder attached to it.
	// GuaranteedStopLossOrders cannot be combined with hedged positions.
	OrderCancelReasonHedgingGsloViolation = OrderCancelReason("HEDGING_GSLO_VIOLATION")

	// OrderCancelReasonAccountPositionValueLimitExceeded means filling the order would cause the maximum position value
	// allowed for the account to be exceeded. The Order has been cancelled as a result.
	OrderCancelReasonAccountPositionValueLimitExceeded = OrderCancelReason("ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED")

	// OrderCancelReasonInstrumentBidReduceOnly means filling the order would require the creation of a short trade,
	// however the instrument is configured such that orders being filled using bid prices can only reduce existing
	// positions. New short positions cannot be created, but existing long positions may be reduced or closed.
	OrderCancelReasonInstrumentBidReduceOnly = OrderCancelReason("INSTRUMENT_BID_REDUCE_ONLY")

	// OrderCancelReasonInstrumentAskReduceOnly means filling the order would require the creation of a long trade,
	// however the instrument is configured such that orders being filled using ask prices can only reduce existing
	// positions. New long positions cannot be created, but existing short positions may be reduced or closed.
	OrderCancelReasonInstrumentAskReduceOnly = OrderCancelReason("INSTRUMENT_ASK_REDUCE_ONLY")

	// OrderCancelReasonInstrumentBidHalted means filling the order would require using the bid, however the instrument
	// is configured such that the bids are halted, and so no short orders may be filled.
	OrderCancelReasonInstrumentBidHalted = OrderCancelReason("INSTRUMENT_BID_HALTED")

	// OrderCancelReasonInstrumentAskHalted means filling the order would require using the ask, however the instrument
	// is configured such that the asks are halted, and so no long orders may be filled.
	OrderCancelReasonInstrumentAskHalted = OrderCancelReason("INSTRUMENT_ASK_HALTED")

	// OrderCancelReasonStopLossOnFillGuaranteedBidHalted means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder (GSLO). Since the trade is long the GSLO would be short, however the bid side is
	// currently halted. GSLOs cannot be created in this situation.
	OrderCancelReasonStopLossOnFillGuaranteedBidHalted = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED")

	// OrderCancelReasonStopLossOnFillGuaranteedAskHalted means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder (GSLO). Since the trade is short the GSLO would be long, however the ask side is
	// currently halted. GSLOs cannot be created in this situation.
	OrderCancelReasonStopLossOnFillGuaranteedAskHalted = OrderCancelReason("STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED")

	// OrderCancelReasonGuaranteedStopLossOnFillBidHalted means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder (GSLO). Since the trade is long the GSLO would be short, however the bid side is
	// currently halted. GSLOs cannot be created in this situation.
	OrderCancelReasonGuaranteedStopLossOnFillBidHalted = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_BID_HALTED")

	// OrderCancelReasonGuaranteedStopLossOnFillAskHalted means filling the Order would result in the creation of a
	// GuaranteedStopLossOrder (GSLO). Since the trade is short the GSLO would be long, however the ask side is
	// currently halted. GSLOs cannot be created in this situation.
	OrderCancelReasonGuaranteedStopLossOnFillAskHalted = OrderCancelReason("GUARANTEED_STOP_LOSS_ON_FILL_ASK_HALTED")

	// OrderCancelReasonFifoViolationSafeguardViolation means filling the Order would have resulted in a new Trade that
	// violates the FIFO violation safeguard constraints.
	OrderCancelReasonFifoViolationSafeguardViolation = OrderCancelReason("FIFO_VIOLATION_SAFEGUARD_VIOLATION")

	// OrderCancelReasonFifoViolationSafeguardPartialCloseViolation means filling the Order would have reduced an
	// existing Trade such that the reduced Trade violates the FIFO violation safeguard constraints.
	OrderCancelReasonFifoViolationSafeguardPartialCloseViolation = OrderCancelReason("FIFO_VIOLATION_SAFEGUARD_PARTIAL_CLOSE_VIOLATION")

	// OrderCancelReasonOrdersOnFillRmoMutualExclusivityMutuallyExclusiveViolation means the Orders on fill would be in
	// violation of the risk management Order mutual exclusivity configuration specifying that only one risk management
	// Order can be attached to a Trade.
	OrderCancelReasonOrdersOnFillRmoMutualExclusivityMutuallyExclusiveViolation = OrderCancelReason("ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION")
)

// OpenTradeDividendAdjustment is used to pay or collect adjustment amount for an open Trade within the Account
type OpenTradeDividendAdjustment struct {
	// The ID of the Trade for which the dividend adjustment is to be paid or collected.
	TradeID TradeID `json:"tradeID"`

	// The dividend adjustment amount to pay or collect for the Trade.
	DividendAdjustment decimal.Decimal `json:"dividendAdjustment"`

	// The dividend adjustment amount to pay or collect for the Trade, in the Instrument’s quote currency.
	QuoteDividendAdjustment decimal.Decimal `json:"quoteDividendAdjustment"`
}

// ClientID is a client-provided identifier, used by clients to refer to their Order or Trades with an identifier that
// they have provided
type ClientID string

// ClientTag is a client-provided tag that can contain any data and may be assigned to their Orders or Trades. Tags are
// typically used to associate groups of Trades and/or Orders together.
type ClientTag string

// ClientComment is a client-provided comment that can contain any data and may be assigned to their Orders or Trades.
// Comments are typically used to provide extra context or meaning to an Order or Trade.
type ClientComment string

// ClientExtensions object allows a client to attach a clientID, tag and comment to Orders and Trades in their Account.
// Do not set, modify, or delete this field if your account is associated with MT4.
type ClientExtensions struct {
	// The Client ID of the Order/Trade
	Id ClientID `json:"id"`

	// A tag associated with the Order/Trade
	Tag ClientTag `json:"tag"`

	// A comment associated with the Order/Trade
	Comment ClientComment `json:"comment"`
}

// TakeProfitDetails specifies the details of a TakeProfitOrder to be created on behalf of a client. This may happen
// when an Order is filled that opens a Trade requiring a TakeProfit, or when a Trade's dependent TakeProfitOrder is
// modified directly though the Trade
type TakeProfitDetails struct {
	// The price that the TakeProfitOrder will be triggered at. Only one of the price and distance fields may be
	// specified.
	Price decimal.Decimal `json:"price"`

	// The time in force for the created TakeProfitOrder. This may only be GTC, GTD or GFD.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`

	// The date when the TakeProfitOrder will be cancelled on if timeInForce is GTD.
	GtdTime *time.Time `json:"gtdTime,omitempty"`

	// The ClientExtensions to add to the TakeProfitOrder when created.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

// StopLossDetails specifies the details of a StopLossOrder to be created on behalf of a client. This may happen when an
// Order is filled that opens a Trade requiring a StopLoss, or when a Trade's dependent StopLossOrder is modified
// directly through the Trade
type StopLossDetails struct {

	// The price that the StopLossOrder will be triggered at. Only one of the price and distance fields may be specified.
	Price *decimal.Decimal `json:"price,omitempty"`

	// Specifies the distance (in price units) from the Trade’s open price to use as the StopLossOrder price. Only one
	// of the distance and price fields may be specified.
	Distance *decimal.Decimal `json:"distance,omitempty"`

	// The time in force for the created StopLossOrder. This may only be GTC, GTD or GFD.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`

	// The date when the StopLossOrder will be cancelled on if timeInForce is GTD.
	GtdTime *time.Time `json:"gtdTime,omitempty"`

	// The ClientExtensions to add to the StopLossOrder when created.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

// GuaranteedStopLossDetails specifies the details of a GuaranteedStopLossOrder to be created on behalf of a client.
// This may happen when an Order is filled that opens a Trade requiring a GuaranteedStopLoss, or when a Trade's
// dependent GuaranteedStopLossOrder is modified directly through the Trade
type GuaranteedStopLossDetails struct {

	// The price that the GuaranteedStopLossOrder will be triggered at. Only one of the price and distance fields may be
	// specified.
	Price *decimal.Decimal `json:"price,omitempty"`

	// Specifies the distance (in price units) from the Trade’s open price to use as the GuaranteedStopLossOrder price.
	// Only one of the distance and price fields may be specified.
	Distance *decimal.Decimal `json:"distance,omitempty"`

	// The time in force for the created GuaranteedStopLossOrder. This may only be GTC, GTD or GFD.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`

	// The date when the GuaranteedStopLossOrder will be cancelled on if timeInForce is GTD.
	GtdTime *time.Time `json:"gtdTime,omitempty"`

	// The ClientExtensions to add to the GuaranteedStopLossOrder when created.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

// TrailingStopLossDetails specifies the details of a TrailingStopLossOrder to be created on behalf of a client. This may happen when an
// Order is filled that opens a Trade requiring a TrailingStopLoss, or when a Trade's dependent TrailingStopLossOrder is modified
// directly through the Trade
type TrailingStopLossDetails struct {

	// The distance (in price units) from the Trade’s fill price that the TrailingStopLossOrder will be triggered at.
	Distance decimal.Decimal `json:"distance"`

	// The time in force for the created TrailingStopLossOrder. This may only be GTC, GTD or GFD.
	// Default: GTC
	TimeInForce *TimeInForce `json:"timeInForce,omitempty"`

	// The date when the TrailingStopLossOrder will be cancelled on if timeInForce is GTD.
	GtdTime *time.Time `json:"gtdTime,omitempty"`

	// The ClientExtensions to add to the TrailingStopLossOrder when created.
	ClientExtensions ClientExtensions `json:"clientExtensions"`
}

// TradeOpen object represents a Trade for an instrument that was opened in an Account. It is found embedded in
// Transactions that affect the position of an instrument in the Account, specifically the OrderFillTransaction.
type TradeOpen struct {
	// The ID of the Trade that was opened
	TradeID TradeID `json:"tradeID"`

	// The number of units opened by the Trade
	Units decimal.Decimal `json:"units"`

	// The average price that the units were opened at.
	Price decimal.Decimal `json:"price"`

	// This is the fee charged for opening the Trade if it has a GuaranteedStopLossOrder attached to it.
	GuaranteedExecutionFee decimal.Decimal `json:"guaranteedExecutionFee"`

	// This is the fee charged for opening the trade if it has a GuaranteedStopLossOrder attached to it, expressed in
	// the Instrument’s quote currency.
	QuoteGuaranteedExecutionFee decimal.Decimal `json:"quoteGuaranteedExecutionFee"`

	// The client extensions for the newly opened Trade
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The half spread cost for the trade open. This can be a positive or negative value and is represented in the home
	// currency of the Account.
	HalfSpreadCost decimal.Decimal `json:"halfSpreadCost"`

	// The margin required at the time the Trade was created. Note, this is the ‘pure’ margin required, it is not the
	// ‘effective’ margin used that factors in the trade risk if a GSLO is attached to the trade.
	InitialMarginRequired decimal.Decimal `json:"initialMarginRequired"`
}

// TradeReduce object represents a Trade for an instrument that was reduced (either partially or fully) in an Account.
// It is found embedded in Transactions that affect the position of an instrument in the account, specifically the
// OrderFillTransaction
type TradeReduce struct {
	// The ID of the Trade that was reduced or closed
	TradeID TradeID `json:"tradeID"`

	// The number of units that the Trade was reduced by
	Units decimal.Decimal `json:"units"`

	// The average price that the units were closed at. This price may be clamped for GuaranteedStopLossOrders.
	Price decimal.Decimal `json:"price"`

	// The PL realized when reducing the Trade
	RealizedPL decimal.Decimal `json:"realizedPL"`

	// The financing paid/collected when reducing the Trade
	Financing decimal.Decimal `json:"financing"`

	// The base financing paid/collected when reducing the Trade
	BaseFinancing decimal.Decimal `json:"baseFinancing"`

	// The quote financing paid/collected when reducing the Trade
	QuoteFinancing decimal.Decimal `json:"quoteFinancing"`

	// The financing rate in effect for the instrument used to calculate the
	// amount of financing paid/collected when reducing the Trade. This field
	// will only be set if the AccountFinancingMode at the time of the order
	// fill is SECOND_BY_SECOND_INSTRUMENT. The value is in decimal rather than
	// percentage points, e.g. 5% is represented as 0.05.
	FinancingRate decimal.Decimal `json:"financingRate"`

	// This is the fee charged for closing the Trade if it has a GuaranteedStopLossOrder attached to it.
	GuaranteedExecutionFee decimal.Decimal `json:"guaranteedExecutionFee"`

	// This is the fee charged for closing the trade if it has a GuaranteedStopLossOrder attached to it, expressed in
	// the Instrument’s quote currency.
	QuoteGuaranteedExecutionFee decimal.Decimal `json:"quoteGuaranteedExecutionFee"`

	// The client extensions for the newly opened Trade
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The half spread cost for the trade reduce/close. This can be a positive or negative value and is represented in
	// the home currency of the Account.
	HalfSpreadCost decimal.Decimal `json:"halfSpreadCost"`
}

// MarketOrderTradeClose specifies the extensions to a MarketOrder that has been created specifically to close a Trade.
type MarketOrderTradeClose struct {
	// The ID of the Trade requested to be closed
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade requested to be closed
	ClientTradeID string `json:"clientTradeID"`

	// Indication of how much of the Trade to close. Either “ALL”, or a DecimalNumber reflection a partial close of the
	// Trade.
	Units string `json:"units"`
}

// MarketOrderMarginCloseout carries the details for the MarketOrder extensions specific to a MarketOrder placed that is
// part of a MarketOrderMarginCloseout ina client's account
type MarketOrderMarginCloseout struct {
	// The reason the MarketOrder was created to perform a margin closeout
	Reason MarketOrderMarginCloseoutReason `json:"reason"`
}

// MarketOrderMarginCloseoutReason represents the reason that the MarketOrder was created to perform a margin closeout
type MarketOrderMarginCloseoutReason string

const (
	// MarketOrderMarginCloseoutReasonMarginCheckViolation represents trade closures resulted from violating OANDA’s margin policy
	MarketOrderMarginCloseoutReasonMarginCheckViolation = MarketOrderMarginCloseoutReason("MARGIN_CHECK_VIOLATION")

	// MarketOrderMarginCloseoutReasonRegulatoryMarginCallViolation represents trade closures came from a margin closeout event resulting from regulatory conditions placed on the Account’s margin call
	MarketOrderMarginCloseoutReasonRegulatoryMarginCallViolation = MarketOrderMarginCloseoutReason("REGULATORY_MARGIN_CALL_VIOLATION")

	// MarketOrderMarginCloseoutReasonRegulatoryMarginCheckViolation represents trade closures resulted from violating the margin policy imposed by regulatory requirements
	MarketOrderMarginCloseoutReasonRegulatoryMarginCheckViolation = MarketOrderMarginCloseoutReason("REGULATORY_MARGIN_CHECK_VIOLATION")
)

// MarketOrderDelayedTradeClose carries the details for the MarketOrder extensions specific to a MarketOrder placed with
// the intent of fully closing a specific open trade that should have already been closed but wasn't due to halted
// market conditions
type MarketOrderDelayedTradeClose struct {
	// The ID of the Trade being closed
	TradeID TradeID `json:"tradeID"`

	// The Client ID of the Trade being closed
	ClientTradeID string `json:"clientTradeID"`

	// The TransactionID of the DelayedTradeClosureTransaction to which this DelayedTradeClose belongs to
	SourceTransactionID TransactionID `json:"sourceTransactionID"`
}

// MarketOrderPositionCloseout specifies the extensions to a MarketOrder when it has been created to closeout a specific
// Position.
type MarketOrderPositionCloseout struct {
	// The instrument of the Position being closed out.
	Instrument string `json:"instrument"`

	// Indication of how much of the Position to close. Either “ALL”, or a
	// DecimalNumber reflection a partial close of the Trade. The DecimalNumber
	// must always be positive, and represent a number that doesn’t exceed the
	// absolute size of the Position.
	Units string `json:"units"`
}

// LiquidityRegenerationSchedule indicates how liquidity that is used when filling an Order for an instrument is
// regenerated following the fill. A liquidity regeneration schedule will be in effect until the timestamp of its final
// step, but may be replaced by a schedule created for an Order of the same instrument that is filled while it is still
// in effect.
type LiquidityRegenerationSchedule struct {
	// The steps in the LiquidityRegenerationSchedule
	Steps []LiquidityRegenerationScheduleStep `json:"steps"`
}

// LiquidityRegenerationScheduleStep indicates the amount of bid and ask liquidity that is used by the Account at a
// certain time. These amounts will only change at the timestamp of the following step.
type LiquidityRegenerationScheduleStep struct {
	// The timestamp of the schedule stop
	Timestamp time.Time `json:"timestamp"`

	// The amount of bid liquidity used at this step in the schedule
	BidLiquidityUsed decimal.Decimal `json:"bidLiquidityUsed"`

	// The amount of ask liquidity used at this step in the schedule
	AskLiquidityUsed decimal.Decimal `json:"askLiquidityUsed"`
}

// OpenTradeFinancing is used to pay/collect daily financing charge for an open Trade within an Account
type OpenTradeFinancing struct {
	// The ID of the Trade that financing is being paid/collected for.
	TradeID TradeID `json:"tradeID"`

	// The amount of financing paid/collected for the Trade.
	Financing decimal.Decimal `json:"financing"`

	// The amount of financing paid/collected in the Instrument’s base currency for the Trade.
	BaseFinancing decimal.Decimal `json:"baseFinancing"`

	// The amount of financing paid/collected in the Instrument’s quote currency for the Trade.
	QuoteFinancing decimal.Decimal `json:"quoteFinancing"`

	// The financing rate in effect for the instrument used to calculate the
	// amount of financing paid/collected for the Trade. This field will only be
	// set if the AccountFinancingMode at the time of the daily financing is
	// DAILY_INSTRUMENT or SECOND_BY_SECOND_INSTRUMENT. The value is in decimal
	// rather than percentage points, e.g. 5% is represented as 0.05.
	FinancingRate decimal.Decimal `json:"financingRate"`
}

// PositionFinancing is used to pay/collect daily financing charge for a Position within an Account
type PositionFinancing struct {
	// The instrument of the Position that financing is being paid/collected for.
	Instrument string `json:"instrument"`

	// The amount of financing paid/collected for the Position.
	Financing decimal.Decimal `json:"financing"`

	// The amount of financing paid/collected in the Instrument’s base currency for the Position.
	BaseFinancing decimal.Decimal `json:"baseFinancing"`

	// The amount of financing paid/collected in the Instrument’s quote currency for the Position.
	QuoteFinancing decimal.Decimal `json:"quoteFinancing"`

	// The HomeConversionFactors in effect for the Position’s Instrument at the time of the DailyFinancing.
	HomeConversionFactors HomeConversionFactors `json:"homeConversionFactors"`

	// The financing paid/collected for each open Trade within the Position.
	OpenTradeFinancings []OpenTradeFinancing `json:"openTradeFinancings"`

	// The account financing mode at the time of the daily financing.
	AccountFinancingMode AccountFinancingMode `json:"accountFinancingMode"`
}

// RequestID is the request identifier
type RequestID string

// ClientRequestID is a client provided request identifier
type ClientRequestID string

// TransactionRejectReason is the reason that a Transaction was rejected
type TransactionRejectReason string

const (
	// TransactionRejectReasonInternalServerError means an unexpected internal server error has occurred
	TransactionRejectReasonInternalServerError = TransactionRejectReason("INTERNAL_SERVER_ERROR")
	// TransactionRejectReasonInstrumentPriceUnknown means the system was unable to determine the current price for the Order’s instrument
	TransactionRejectReasonInstrumentPriceUnknown = TransactionRejectReason("INSTRUMENT_PRICE_UNKNOWN")
	// TransactionRejectReasonAccountNotActive means the Account is not active
	TransactionRejectReasonAccountNotActive = TransactionRejectReason("ACCOUNT_NOT_ACTIVE")
	// TransactionRejectReasonAccountLocked means the Account is locked
	TransactionRejectReasonAccountLocked = TransactionRejectReason("ACCOUNT_LOCKED")
	// TransactionRejectReasonAccountOrderCreationLocked means the Account is locked for Order creation
	TransactionRejectReasonAccountOrderCreationLocked = TransactionRejectReason("ACCOUNT_ORDER_CREATION_LOCKED")
	// TransactionRejectReasonAccountConfigurationLocked means the Account is locked for configuration
	TransactionRejectReasonAccountConfigurationLocked = TransactionRejectReason("ACCOUNT_CONFIGURATION_LOCKED")
	// TransactionRejectReasonAccountDepositLocked means the Account is locked for deposits
	TransactionRejectReasonAccountDepositLocked = TransactionRejectReason("ACCOUNT_DEPOSIT_LOCKED")
	// TransactionRejectReasonAccountWithdrawalLocked means the Account is locked for withdrawals
	TransactionRejectReasonAccountWithdrawalLocked = TransactionRejectReason("ACCOUNT_WITHDRAWAL_LOCKED")
	// TransactionRejectReasonAccountOrderCancelLocked means the Account is locked for Order cancellation
	TransactionRejectReasonAccountOrderCancelLocked = TransactionRejectReason("ACCOUNT_ORDER_CANCEL_LOCKED")
	// TransactionRejectReasonInstrumentNotTradeable means the instrument specified is not tradeable by the Account
	TransactionRejectReasonInstrumentNotTradeable = TransactionRejectReason("INSTRUMENT_NOT_TRADEABLE")
	// TransactionRejectReasonPendingOrdersAllowedExceeded means creating the Order would result in the maximum number of allowed pending Orders being exceeded
	TransactionRejectReasonPendingOrdersAllowedExceeded = TransactionRejectReason("PENDING_ORDERS_ALLOWED_EXCEEDED")
	// TransactionRejectReasonOrderIdUnspecified means neither the Order ID nor client Order ID are specified
	TransactionRejectReasonOrderIdUnspecified = TransactionRejectReason("ORDER_ID_UNSPECIFIED")
	// TransactionRejectReasonOrderDoesntExist means the Order specified does not exist
	TransactionRejectReasonOrderDoesntExist = TransactionRejectReason("ORDER_DOESNT_EXIST")
	// TransactionRejectReasonOrderIdentifierInconsistency means the Order ID and client Order ID specified do not identify the same Order
	TransactionRejectReasonOrderIdentifierInconsistency = TransactionRejectReason("ORDER_IDENTIFIER_INCONSISTENCY")
	// TransactionRejectReasonTradeIdUnspecified means neither the Trade ID nor client Trade ID are specified
	TransactionRejectReasonTradeIdUnspecified = TransactionRejectReason("TRADE_ID_UNSPECIFIED")
	// TransactionRejectReasonTradeDoesntExist means the Trade specified does not exist
	TransactionRejectReasonTradeDoesntExist = TransactionRejectReason("TRADE_DOESNT_EXIST")
	// TransactionRejectReasonTradeIdentifierInconsistency means the Trade ID and client Trade ID specified do not identify the same Trade
	TransactionRejectReasonTradeIdentifierInconsistency = TransactionRejectReason("TRADE_IDENTIFIER_INCONSISTENCY")
	// TransactionRejectReasonInsufficientMargin means the Account had insufficient margin to perform the action specified. One possible reason for this is due to the creation or modification of a guaranteed StopLoss Order.
	TransactionRejectReasonInsufficientMargin = TransactionRejectReason("INSUFFICIENT_MARGIN")
	// TransactionRejectReasonInstrumentMissing means order instrument has not been specified
	TransactionRejectReasonInstrumentMissing = TransactionRejectReason("INSTRUMENT_MISSING")
	// TransactionRejectReasonInstrumentUnknown means the instrument specified is unknown
	TransactionRejectReasonInstrumentUnknown = TransactionRejectReason("INSTRUMENT_UNKNOWN")
	// TransactionRejectReasonUnitsMissing means order units have not been not specified
	TransactionRejectReasonUnitsMissing = TransactionRejectReason("UNITS_MISSING")
	// TransactionRejectReasonUnitsInvalid means Order units specified are invalid
	TransactionRejectReasonUnitsInvalid = TransactionRejectReason("UNITS_INVALID")
	// TransactionRejectReasonUnitsPrecisionExceeded means the units specified contain more precision than is allowed for the Order’s instrument
	TransactionRejectReasonUnitsPrecisionExceeded = TransactionRejectReason("UNITS_PRECISION_EXCEEDED")
	// TransactionRejectReasonUnitsLimitExceeded means the units specified exceeds the maximum number of units allowed
	TransactionRejectReasonUnitsLimitExceeded = TransactionRejectReason("UNITS_LIMIT_EXCEEDED")
	// TransactionRejectReasonUnitsMinimumNotMet means the units specified is less than the minimum number of units required
	TransactionRejectReasonUnitsMinimumNotMet = TransactionRejectReason("UNITS_MINIMUM_NOT_MET")
	// TransactionRejectReasonPriceMissing means the price has not been specified
	TransactionRejectReasonPriceMissing = TransactionRejectReason("PRICE_MISSING")
	// TransactionRejectReasonPriceInvalid means the price specified is invalid
	TransactionRejectReasonPriceInvalid = TransactionRejectReason("PRICE_INVALID")
	// TransactionRejectReasonPricePrecisionExceeded means the price specified contains more precision than is allowed for the instrument
	TransactionRejectReasonPricePrecisionExceeded = TransactionRejectReason("PRICE_PRECISION_EXCEEDED")
	// TransactionRejectReasonPriceDistanceMissing means the price distance has not been specified
	TransactionRejectReasonPriceDistanceMissing = TransactionRejectReason("PRICE_DISTANCE_MISSING")
	// TransactionRejectReasonPriceDistanceInvalid means the price distance specified is invalid
	TransactionRejectReasonPriceDistanceInvalid = TransactionRejectReason("PRICE_DISTANCE_INVALID")
	// TransactionRejectReasonPriceDistancePrecisionExceeded means the price distance specified contains more precision than is allowed for the instrument
	TransactionRejectReasonPriceDistancePrecisionExceeded = TransactionRejectReason("PRICE_DISTANCE_PRECISION_EXCEEDED")
	// TransactionRejectReasonPriceDistanceMaximumExceeded means the price distance exceeds that maximum allowed amount
	TransactionRejectReasonPriceDistanceMaximumExceeded = TransactionRejectReason("PRICE_DISTANCE_MAXIMUM_EXCEEDED")
	// TransactionRejectReasonPriceDistanceMinimumNotMet means the price distance does not meet the minimum allowed amount
	TransactionRejectReasonPriceDistanceMinimumNotMet = TransactionRejectReason("PRICE_DISTANCE_MINIMUM_NOT_MET")
	// TransactionRejectReasonTimeInForceMissing means the TimeInForce field has not been specified
	TransactionRejectReasonTimeInForceMissing = TransactionRejectReason("TIME_IN_FORCE_MISSING")
	// TransactionRejectReasonTimeInForceInvalid means the TimeInForce specified is invalid
	TransactionRejectReasonTimeInForceInvalid = TransactionRejectReason("TIME_IN_FORCE_INVALID")
	// TransactionRejectReasonTimeInForceGtdTimestampMissing means the TimeInForce is GTD but no GTD timestamp is provided
	TransactionRejectReasonTimeInForceGtdTimestampMissing = TransactionRejectReason("TIME_IN_FORCE_GTD_TIMESTAMP_MISSING")
	// TransactionRejectReasonTimeInForceGtdTimestampInPast means the TimeInForce is GTD but the GTD timestamp is in the past
	TransactionRejectReasonTimeInForceGtdTimestampInPast = TransactionRejectReason("TIME_IN_FORCE_GTD_TIMESTAMP_IN_PAST")
	// TransactionRejectReasonPriceBoundInvalid means the price bound specified is invalid
	TransactionRejectReasonPriceBoundInvalid = TransactionRejectReason("PRICE_BOUND_INVALID")
	// TransactionRejectReasonPriceBoundPrecisionExceeded means the price bound specified contains more precision than is allowed for the Order’s instrument
	TransactionRejectReasonPriceBoundPrecisionExceeded = TransactionRejectReason("PRICE_BOUND_PRECISION_EXCEEDED")
	// TransactionRejectReasonOrdersOnFillDuplicateClientOrderIds means multiple Orders on fill share the same client Order ID
	TransactionRejectReasonOrdersOnFillDuplicateClientOrderIds = TransactionRejectReason("ORDERS_ON_FILL_DUPLICATE_CLIENT_ORDER_IDS")
	// TransactionRejectReasonTradeOnFillClientExtensionsNotSupported means the Order does not support Trade on fill client extensions because it cannot create a new Trade
	TransactionRejectReasonTradeOnFillClientExtensionsNotSupported = TransactionRejectReason("TRADE_ON_FILL_CLIENT_EXTENSIONS_NOT_SUPPORTED")
	// TransactionRejectReasonClientOrderIdInvalid means the client Order ID specified is invalid
	TransactionRejectReasonClientOrderIdInvalid = TransactionRejectReason("CLIENT_ORDER_ID_INVALID")
	// TransactionRejectReasonClientOrderIdAlreadyExists means the client Order ID specified is already assigned to another pending Order
	TransactionRejectReasonClientOrderIdAlreadyExists = TransactionRejectReason("CLIENT_ORDER_ID_ALREADY_EXISTS")
	// TransactionRejectReasonClientOrderTagInvalid means the client Order tag specified is invalid
	TransactionRejectReasonClientOrderTagInvalid = TransactionRejectReason("CLIENT_ORDER_TAG_INVALID")
	// TransactionRejectReasonClientOrderCommentInvalid means the client Order comment specified is invalid
	TransactionRejectReasonClientOrderCommentInvalid = TransactionRejectReason("CLIENT_ORDER_COMMENT_INVALID")
	// TransactionRejectReasonClientTradeIdInvalid means the client Trade ID specified is invalid
	TransactionRejectReasonClientTradeIdInvalid = TransactionRejectReason("CLIENT_TRADE_ID_INVALID")
	// TransactionRejectReasonClientTradeIdAlreadyExists means the client Trade ID specified is already assigned to another open Trade
	TransactionRejectReasonClientTradeIdAlreadyExists = TransactionRejectReason("CLIENT_TRADE_ID_ALREADY_EXISTS")
	// TransactionRejectReasonClientTradeTagInvalid means the client Trade tag specified is invalid
	TransactionRejectReasonClientTradeTagInvalid = TransactionRejectReason("CLIENT_TRADE_TAG_INVALID")
	// TransactionRejectReasonClientTradeCommentInvalid means the client Trade comment is invalid
	TransactionRejectReasonClientTradeCommentInvalid = TransactionRejectReason("CLIENT_TRADE_COMMENT_INVALID")
	// TransactionRejectReasonOrderFillPositionActionMissing means the OrderFillPositionAction field has not been specified
	TransactionRejectReasonOrderFillPositionActionMissing = TransactionRejectReason("ORDER_FILL_POSITION_ACTION_MISSING")
	// TransactionRejectReasonOrderFillPositionActionInvalid means the OrderFillPositionAction specified is invalid
	TransactionRejectReasonOrderFillPositionActionInvalid = TransactionRejectReason("ORDER_FILL_POSITION_ACTION_INVALID")
	// TransactionRejectReasonTriggerConditionMissing means the TriggerCondition field has not been specified
	TransactionRejectReasonTriggerConditionMissing = TransactionRejectReason("TRIGGER_CONDITION_MISSING")
	// TransactionRejectReasonTriggerConditionInvalid means the TriggerCondition specified is invalid
	TransactionRejectReasonTriggerConditionInvalid = TransactionRejectReason("TRIGGER_CONDITION_INVALID")
	// TransactionRejectReasonOrderPartialFillOptionMissing means the OrderFillPositionAction field has not been specified
	TransactionRejectReasonOrderPartialFillOptionMissing = TransactionRejectReason("ORDER_PARTIAL_FILL_OPTION_MISSING")
	// TransactionRejectReasonOrderPartialFillOptionInvalid means the OrderFillPositionAction specified is invalid.
	TransactionRejectReasonOrderPartialFillOptionInvalid = TransactionRejectReason("ORDER_PARTIAL_FILL_OPTION_INVALID")
	// TransactionRejectReasonInvalidReissueImmediatePartialFill means when attempting to reissue an order (currently only a MarketIfTouched) that was immediately partially filled, it is not possible to create a correct pending Order.
	TransactionRejectReasonInvalidReissueImmediatePartialFill = TransactionRejectReason("INVALID_REISSUE_IMMEDIATE_PARTIAL_FILL")
	// TransactionRejectReasonOrdersOnFillRmoMutualExclusivityMutuallyExclusiveViolation means the Orders on fill would be in violation of the risk management Order mutual exclusivity configuration specifying that only one risk management Order can be attached to a Trade.
	TransactionRejectReasonOrdersOnFillRmoMutualExclusivityMutuallyExclusiveViolation = TransactionRejectReason("ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION")
	// TransactionRejectReasonOrdersOnFillRmoMutualExclusivityGsloExcludesOthersViolation means the Orders on fill would be in violation of the risk management Order mutual exclusivity configuration specifying that if a GSLO is already attached to a Trade, no other risk management Order can be attached to a Trade.
	TransactionRejectReasonOrdersOnFillRmoMutualExclusivityGsloExcludesOthersViolation = TransactionRejectReason("ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION")
	// TransactionRejectReasonTakeProfitOrderAlreadyExists means a Take Profit Order for the specified Trade already exists
	TransactionRejectReasonTakeProfitOrderAlreadyExists = TransactionRejectReason("TAKE_PROFIT_ORDER_ALREADY_EXISTS")
	// TransactionRejectReasonTakeProfitOrderWouldViolateFifoViolationSafeguard means the Take Profit Order would cause the associated Trade to be in violation of the FIFO violation safeguard constraints.
	TransactionRejectReasonTakeProfitOrderWouldViolateFifoViolationSafeguard = TransactionRejectReason("TAKE_PROFIT_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD")
	// TransactionRejectReasonTakeProfitOnFillPriceMissing means the Take Profit on fill specified does not provide a price
	TransactionRejectReasonTakeProfitOnFillPriceMissing = TransactionRejectReason("TAKE_PROFIT_ON_FILL_PRICE_MISSING")
	// TransactionRejectReasonTakeProfitOnFillPriceInvalid means the Take Profit on fill specified contains an invalid price
	TransactionRejectReasonTakeProfitOnFillPriceInvalid = TransactionRejectReason("TAKE_PROFIT_ON_FILL_PRICE_INVALID")
	// TransactionRejectReasonTakeProfitOnFillPricePrecisionExceeded means the Take Profit on fill specified contains a price with more precision than is allowed by the Order’s instrument
	TransactionRejectReasonTakeProfitOnFillPricePrecisionExceeded = TransactionRejectReason("TAKE_PROFIT_ON_FILL_PRICE_PRECISION_EXCEEDED")
	// TransactionRejectReasonTakeProfitOnFillTimeInForceMissing means the Take Profit on fill specified does not provide a TimeInForce
	TransactionRejectReasonTakeProfitOnFillTimeInForceMissing = TransactionRejectReason("TAKE_PROFIT_ON_FILL_TIME_IN_FORCE_MISSING")
	// TransactionRejectReasonTakeProfitOnFillTimeInForceInvalid means the Take Profit on fill specifies an invalid TimeInForce
	TransactionRejectReasonTakeProfitOnFillTimeInForceInvalid = TransactionRejectReason("TAKE_PROFIT_ON_FILL_TIME_IN_FORCE_INVALID")
	// TransactionRejectReasonTakeProfitOnFillGtdTimestampMissing means the Take Profit on fill specifies a GTD TimeInForce but does not provide a GTD timestamp
	TransactionRejectReasonTakeProfitOnFillGtdTimestampMissing = TransactionRejectReason("TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_MISSING")
	// TransactionRejectReasonTakeProfitOnFillGtdTimestampInPast means the Take Profit on fill specifies a GTD timestamp that is in the past
	TransactionRejectReasonTakeProfitOnFillGtdTimestampInPast = TransactionRejectReason("TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST")
	// TransactionRejectReasonTakeProfitOnFillClientOrderIdInvalid means the Take Profit on fill client Order ID specified is invalid
	TransactionRejectReasonTakeProfitOnFillClientOrderIdInvalid = TransactionRejectReason("TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_INVALID")
	// TransactionRejectReasonTakeProfitOnFillClientOrderTagInvalid means the Take Profit on fill client Order tag specified is invalid
	TransactionRejectReasonTakeProfitOnFillClientOrderTagInvalid = TransactionRejectReason("TAKE_PROFIT_ON_FILL_CLIENT_ORDER_TAG_INVALID")
	// TransactionRejectReasonTakeProfitOnFillClientOrderCommentInvalid means the Take Profit on fill client Order comment specified is invalid
	TransactionRejectReasonTakeProfitOnFillClientOrderCommentInvalid = TransactionRejectReason("TAKE_PROFIT_ON_FILL_CLIENT_ORDER_COMMENT_INVALID")
	// TransactionRejectReasonTakeProfitOnFillTriggerConditionMissing means the Take Profit on fill specified does not provide a TriggerCondition
	TransactionRejectReasonTakeProfitOnFillTriggerConditionMissing = TransactionRejectReason("TAKE_PROFIT_ON_FILL_TRIGGER_CONDITION_MISSING")
	// TransactionRejectReasonTakeProfitOnFillTriggerConditionInvalid means the Take Profit on fill specifies an invalid TriggerCondition
	TransactionRejectReasonTakeProfitOnFillTriggerConditionInvalid = TransactionRejectReason("TAKE_PROFIT_ON_FILL_TRIGGER_CONDITION_INVALID")
	// TransactionRejectReasonStopLossOrderAlreadyExists means a Stop Loss Order for the specified Trade already exists
	TransactionRejectReasonStopLossOrderAlreadyExists = TransactionRejectReason("STOP_LOSS_ORDER_ALREADY_EXISTS")
	// TransactionRejectReasonStopLossOrderGuaranteedRequired means an attempt was made to to create a non-guaranteed stop loss order in an account that requires all stop loss orders to be guaranteed.
	TransactionRejectReasonStopLossOrderGuaranteedRequired = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_REQUIRED")
	// TransactionRejectReasonStopLossOrderGuaranteedPriceWithinSpread means an attempt to create a guaranteed stop loss order with a price that is within the current tradeable spread.
	TransactionRejectReasonStopLossOrderGuaranteedPriceWithinSpread = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_PRICE_WITHIN_SPREAD")
	// TransactionRejectReasonStopLossOrderGuaranteedNotAllowed means an attempt was made to create a guaranteed Stop Loss Order, however the Account’s configuration does not allow guaranteed Stop Loss Orders.
	TransactionRejectReasonStopLossOrderGuaranteedNotAllowed = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_NOT_ALLOWED")
	// TransactionRejectReasonStopLossOrderGuaranteedHaltedCreateViolation means an attempt was made to create a guaranteed Stop Loss Order when the market was halted.
	TransactionRejectReasonStopLossOrderGuaranteedHaltedCreateViolation = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_HALTED_CREATE_VIOLATION")
	// TransactionRejectReasonStopLossOrderGuaranteedHaltedTightenViolation means an attempt was made to re-create a guaranteed Stop Loss Order with a tighter fill price when the market was halted.
	TransactionRejectReasonStopLossOrderGuaranteedHaltedTightenViolation = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_HALTED_TIGHTEN_VIOLATION")
	// TransactionRejectReasonStopLossOrderGuaranteedHedgingNotAllowed means an attempt was made to create a guaranteed Stop Loss Order on a hedged Trade (ie there is an existing open Trade in the opposing direction), however the Account’s configuration does not allow guaranteed Stop Loss Orders for hedged Trades/Positions.
	TransactionRejectReasonStopLossOrderGuaranteedHedgingNotAllowed = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_HEDGING_NOT_ALLOWED")
	// TransactionRejectReasonStopLossOrderGuaranteedMinimumDistanceNotMet means an attempt was made to create a guaranteed Stop Loss Order, however the distance between the current price and the trigger price does not meet the Account’s configured minimum Guaranteed Stop Loss distance.
	TransactionRejectReasonStopLossOrderGuaranteedMinimumDistanceNotMet = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_MINIMUM_DISTANCE_NOT_MET")
	// TransactionRejectReasonStopLossOrderNotCancelable means an attempt was made to cancel a Stop Loss Order, however the Account’s configuration requires every Trade have an associated Stop Loss Order.
	TransactionRejectReasonStopLossOrderNotCancelable = TransactionRejectReason("STOP_LOSS_ORDER_NOT_CANCELABLE")
	// TransactionRejectReasonStopLossOrderNotReplaceable means an attempt was made to cancel and replace a Stop Loss Order, however the Account’s configuration prevents the modification of Stop Loss Orders.
	TransactionRejectReasonStopLossOrderNotReplaceable = TransactionRejectReason("STOP_LOSS_ORDER_NOT_REPLACEABLE")
	// TransactionRejectReasonStopLossOrderGuaranteedLevelRestrictionExceeded means an attempt was made to create a guaranteed Stop Loss Order, however doing so would exceed the Account’s configured guaranteed StopLoss Order level restriction volume.
	TransactionRejectReasonStopLossOrderGuaranteedLevelRestrictionExceeded = TransactionRejectReason("STOP_LOSS_ORDER_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED")
	// TransactionRejectReasonStopLossOrderPriceAndDistanceBothSpecified means the Stop Loss Order request contains both the price and distance fields.
	TransactionRejectReasonStopLossOrderPriceAndDistanceBothSpecified = TransactionRejectReason("STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_SPECIFIED")
	// TransactionRejectReasonStopLossOrderPriceAndDistanceBothMissing means the Stop Loss Order request contains neither the price nor distance fields.
	TransactionRejectReasonStopLossOrderPriceAndDistanceBothMissing = TransactionRejectReason("STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_MISSING")
	// TransactionRejectReasonStopLossOrderWouldViolateFifoViolationSafeguard means the Stop Loss Order would cause the associated Trade to be in violation of the FIFO violation safeguard constraints
	TransactionRejectReasonStopLossOrderWouldViolateFifoViolationSafeguard = TransactionRejectReason("STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD")
	// TransactionRejectReasonStopLossOrderRmoMutualExclusivityMutuallyExclusiveViolation means the Stop Loss Order would be in violation of the risk management Order mutual exclusivity configuration specifying that only one risk management order can be attached to a Trade.
	TransactionRejectReasonStopLossOrderRmoMutualExclusivityMutuallyExclusiveViolation = TransactionRejectReason("STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION")
	// TransactionRejectReasonStopLossOrderRmoMutualExclusivityGsloExcludesOthersViolation means the Stop Loss Order would be in violation of the risk management Order mutual exclusivity configuration specifying that if a GSLO is already attached to a Trade, no other risk management Order can be attached to the same Trade.
	TransactionRejectReasonStopLossOrderRmoMutualExclusivityGsloExcludesOthersViolation = TransactionRejectReason("STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION")
	// TransactionRejectReasonStopLossOnFillRequiredForPendingOrder means an attempt to create a pending Order was made with no Stop Loss Order on fill specified and the Account’s configuration requires that every Trade have an associated Stop Loss Order.
	TransactionRejectReasonStopLossOnFillRequiredForPendingOrder = TransactionRejectReason("STOP_LOSS_ON_FILL_REQUIRED_FOR_PENDING_ORDER")
	// TransactionRejectReasonStopLossOnFillGuaranteedNotAllowed means an attempt to create a pending Order was made with a Stop Loss Order on fill that was explicitly configured to be guaranteed, however the Account’s configuration does not allow guaranteed Stop Loss Orders.
	TransactionRejectReasonStopLossOnFillGuaranteedNotAllowed = TransactionRejectReason("STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED")
	// TransactionRejectReasonStopLossOnFillGuaranteedRequired means an attempt to create a pending Order was made with a Stop Loss Order on fill that was explicitly configured to be not guaranteed, however the Account’s configuration requires guaranteed Stop Loss Orders.
	TransactionRejectReasonStopLossOnFillGuaranteedRequired = TransactionRejectReason("STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED")
	// TransactionRejectReasonStopLossOnFillPriceMissing means the Stop Loss on fill specified does not provide a price
	TransactionRejectReasonStopLossOnFillPriceMissing = TransactionRejectReason("STOP_LOSS_ON_FILL_PRICE_MISSING")
	// TransactionRejectReasonStopLossOnFillPriceInvalid means the Stop Loss on fill specifies an invalid price
	TransactionRejectReasonStopLossOnFillPriceInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_PRICE_INVALID")
	// TransactionRejectReasonStopLossOnFillPricePrecisionExceeded means the Stop Loss on fill specifies a price with more precision than is allowed by the Order’s instrument
	TransactionRejectReasonStopLossOnFillPricePrecisionExceeded = TransactionRejectReason("STOP_LOSS_ON_FILL_PRICE_PRECISION_EXCEEDED")
	// TransactionRejectReasonStopLossOnFillGuaranteedMinimumDistanceNotMet means an attempt to create a pending Order was made with the distance between the guaranteed Stop Loss Order on fill’s price and the pending Order’s price is less than the Account’s configured minimum guaranteed stop loss distance.
	TransactionRejectReasonStopLossOnFillGuaranteedMinimumDistanceNotMet = TransactionRejectReason("STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET")
	// TransactionRejectReasonStopLossOnFillGuaranteedLevelRestrictionExceeded means an attempt to create a pending Order was made with a guaranteed Stop Loss Order on fill configured, and the Order’s units exceed the Account’s configured guaranteed StopLoss Order level restriction volume.
	TransactionRejectReasonStopLossOnFillGuaranteedLevelRestrictionExceeded = TransactionRejectReason("STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED")
	// TransactionRejectReasonStopLossOnFillDistanceInvalid means the Stop Loss on fill distance is invalid
	TransactionRejectReasonStopLossOnFillDistanceInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_DISTANCE_INVALID")
	// TransactionRejectReasonStopLossOnFillPriceDistanceMaximumExceeded means the Stop Loss on fill price distance exceeds the maximum allowed amount
	TransactionRejectReasonStopLossOnFillPriceDistanceMaximumExceeded = TransactionRejectReason("STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED")
	// TransactionRejectReasonStopLossOnFillDistancePrecisionExceeded means the Stop Loss on fill distance contains more precision than is allowed by the instrument
	TransactionRejectReasonStopLossOnFillDistancePrecisionExceeded = TransactionRejectReason("STOP_LOSS_ON_FILL_DISTANCE_PRECISION_EXCEEDED")
	// TransactionRejectReasonStopLossOnFillPriceAndDistanceBothSpecified means the Stop Loss on fill contains both the price and distance fields.
	TransactionRejectReasonStopLossOnFillPriceAndDistanceBothSpecified = TransactionRejectReason("STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_SPECIFIED")
	// TransactionRejectReasonStopLossOnFillPriceAndDistanceBothMissing means the Stop Loss on fill contains neither the price nor distance fields.
	TransactionRejectReasonStopLossOnFillPriceAndDistanceBothMissing = TransactionRejectReason("STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_MISSING")
	// TransactionRejectReasonStopLossOnFillTimeInForceMissing means the Stop Loss on fill specified does not provide a TimeInForce
	TransactionRejectReasonStopLossOnFillTimeInForceMissing = TransactionRejectReason("STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING")
	// TransactionRejectReasonStopLossOnFillTimeInForceInvalid means the Stop Loss on fill specifies an invalid TimeInForce
	TransactionRejectReasonStopLossOnFillTimeInForceInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID")
	// TransactionRejectReasonStopLossOnFillGtdTimestampMissing means the Stop Loss on fill specifies a GTD TimeInForce but does not provide a GTD timestamp
	TransactionRejectReasonStopLossOnFillGtdTimestampMissing = TransactionRejectReason("STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING")
	// TransactionRejectReasonStopLossOnFillGtdTimestampInPast means the Stop Loss on fill specifies a GTD timestamp that is in the past
	TransactionRejectReasonStopLossOnFillGtdTimestampInPast = TransactionRejectReason("STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST")
	// TransactionRejectReasonStopLossOnFillClientOrderIdInvalid means the Stop Loss on fill client Order ID specified is invalid
	TransactionRejectReasonStopLossOnFillClientOrderIdInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID")
	// TransactionRejectReasonStopLossOnFillClientOrderTagInvalid means the Stop Loss on fill client Order tag specified is invalid
	TransactionRejectReasonStopLossOnFillClientOrderTagInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID")
	// TransactionRejectReasonStopLossOnFillClientOrderCommentInvalid means the Stop Loss on fill client Order comment specified is invalid
	TransactionRejectReasonStopLossOnFillClientOrderCommentInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID")
	// TransactionRejectReasonStopLossOnFillTriggerConditionMissing means the Stop Loss on fill specified does not provide a TriggerCondition
	TransactionRejectReasonStopLossOnFillTriggerConditionMissing = TransactionRejectReason("STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING")
	// TransactionRejectReasonStopLossOnFillTriggerConditionInvalid means the Stop Loss on fill specifies an invalid TriggerCondition
	TransactionRejectReasonStopLossOnFillTriggerConditionInvalid = TransactionRejectReason("STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOrderAlreadyExists means a Guaranteed Stop Loss Order for the specified Trade already exists
	TransactionRejectReasonGuaranteedStopLossOrderAlreadyExists = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_ALREADY_EXISTS")
	// TransactionRejectReasonGuaranteedStopLossOrderRequired means an attempt was made to to create a non-guaranteed stop loss order in an account that requires all stop loss orders to be guaranteed.
	TransactionRejectReasonGuaranteedStopLossOrderRequired = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_REQUIRED")
	// TransactionRejectReasonGuaranteedStopLossOrderPriceWithinSpread means an attempt to create a guaranteed stop loss order with a price that is within the current tradeable spread.
	TransactionRejectReasonGuaranteedStopLossOrderPriceWithinSpread = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_PRICE_WITHIN_SPREAD")
	// TransactionRejectReasonGuaranteedStopLossOrderNotAllowed means an attempt was made to create a Guaranteed Stop Loss Order, however the Account’s configuration does not allow Guaranteed Stop Loss Orders.
	TransactionRejectReasonGuaranteedStopLossOrderNotAllowed = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_NOT_ALLOWED")
	// TransactionRejectReasonGuaranteedStopLossOrderHaltedCreateViolation means an attempt was made to create a Guaranteed Stop Loss Order when the market was halted.
	TransactionRejectReasonGuaranteedStopLossOrderHaltedCreateViolation = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_HALTED_CREATE_VIOLATION")
	// TransactionRejectReasonGuaranteedStopLossOrderCreateViolation means an attempt was made to create a Guaranteed Stop Loss Order when the market was open.
	TransactionRejectReasonGuaranteedStopLossOrderCreateViolation = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_CREATE_VIOLATION")
	// TransactionRejectReasonGuaranteedStopLossOrderHaltedTightenViolation means an attempt was made to re-create a Guaranteed Stop Loss Order with a tighter fill price when the market was halted.
	TransactionRejectReasonGuaranteedStopLossOrderHaltedTightenViolation = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_HALTED_TIGHTEN_VIOLATION")
	// TransactionRejectReasonGuaranteedStopLossOrderTightenViolation means an attempt was made to re-create a Guaranteed Stop Loss Order with a tighter fill price when the market was open.
	TransactionRejectReasonGuaranteedStopLossOrderTightenViolation = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_TIGHTEN_VIOLATION")
	// TransactionRejectReasonGuaranteedStopLossOrderHedgingNotAllowed means an attempt was made to create a Guaranteed Stop Loss Order on a hedged Trade (ie there is an existing open Trade in the opposing direction), however the Account’s configuration does not allow Guaranteed Stop Loss Orders for hedged Trades/Positions.
	TransactionRejectReasonGuaranteedStopLossOrderHedgingNotAllowed = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_HEDGING_NOT_ALLOWED")
	// TransactionRejectReasonGuaranteedStopLossOrderMinimumDistanceNotMet means an attempt was made to create a Guaranteed Stop Loss Order, however the distance between the current price and the trigger price does not meet the Account’s configured minimum Guaranteed Stop Loss distance.
	TransactionRejectReasonGuaranteedStopLossOrderMinimumDistanceNotMet = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_MINIMUM_DISTANCE_NOT_MET")
	// TransactionRejectReasonGuaranteedStopLossOrderNotCancelable means an attempt was made to cancel a Guaranteed Stop Loss Order when the market is open, however the Account’s configuration requires every Trade have an associated Guaranteed Stop Loss Order.
	TransactionRejectReasonGuaranteedStopLossOrderNotCancelable = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_NOT_CANCELABLE")
	// TransactionRejectReasonGuaranteedStopLossOrderHaltedNotCancelable means an attempt was made to cancel a Guaranteed Stop Loss Order when the market is halted, however the Account’s configuration requires every Trade have an associated Guaranteed Stop Loss Order.
	TransactionRejectReasonGuaranteedStopLossOrderHaltedNotCancelable = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_HALTED_NOT_CANCELABLE")
	// TransactionRejectReasonGuaranteedStopLossOrderNotReplaceable means an attempt was made to cancel and replace a Guaranteed Stop Loss Order when the market is open, however the Account’s configuration prevents the modification of Guaranteed Stop Loss Orders.
	TransactionRejectReasonGuaranteedStopLossOrderNotReplaceable = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_NOT_REPLACEABLE")
	// TransactionRejectReasonGuaranteedStopLossOrderHaltedNotReplaceable means an attempt was made to cancel and replace a Guaranteed Stop Loss Order when the market is halted, however the Account’s configuration prevents the modification of Guaranteed Stop Loss Orders.
	TransactionRejectReasonGuaranteedStopLossOrderHaltedNotReplaceable = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_HALTED_NOT_REPLACEABLE")
	// TransactionRejectReasonGuaranteedStopLossOrderLevelRestrictionVolumeExceeded means an attempt was made to create a Guaranteed Stop Loss Order, however doing so would exceed the Account’s configured guaranteed StopLoss Order level restriction volume.
	TransactionRejectReasonGuaranteedStopLossOrderLevelRestrictionVolumeExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_LEVEL_RESTRICTION_VOLUME_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOrderLevelRestrictionPriceRangeExceeded means an attempt was made to create a Guaranteed Stop Loss Order, however doing so would exceed the Account’s configured guaranteed StopLoss Order level restriction price range.
	TransactionRejectReasonGuaranteedStopLossOrderLevelRestrictionPriceRangeExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOrderPriceAndDistanceBothSpecified means the Guaranteed Stop Loss Order request contains both the price and distance fields.
	TransactionRejectReasonGuaranteedStopLossOrderPriceAndDistanceBothSpecified = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_SPECIFIED")
	// TransactionRejectReasonGuaranteedStopLossOrderPriceAndDistanceBothMissing means the Guaranteed Stop Loss Order request contains neither the price nor distance fields.
	TransactionRejectReasonGuaranteedStopLossOrderPriceAndDistanceBothMissing = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_MISSING")
	// TransactionRejectReasonGuaranteedStopLossOrderWouldViolateFifoViolationSafeguard means the Guaranteed Stop Loss Order would cause the associated Trade to be in violation of the FIFO violation safeguard constraints
	TransactionRejectReasonGuaranteedStopLossOrderWouldViolateFifoViolationSafeguard = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD")
	// TransactionRejectReasonGuaranteedStopLossOrderRmoMutualExclusivityMutuallyExclusiveViolation means the Guaranteed Stop Loss Order would be in violation of the risk management Order mutual exclusivity configuration specifying that only one risk management order can be attached to a Trade.
	TransactionRejectReasonGuaranteedStopLossOrderRmoMutualExclusivityMutuallyExclusiveViolation = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION")
	// TransactionRejectReasonGuaranteedStopLossOrderRmoMutualExclusivityGsloExcludesOthersViolation means the Guaranteed Stop Loss Order would be in violation of the risk management Order mutual exclusivity configuration specifying that if a GSLO is already attached to a Trade, no other risk management Order can be attached to the same Trade.
	TransactionRejectReasonGuaranteedStopLossOrderRmoMutualExclusivityGsloExcludesOthersViolation = TransactionRejectReason("GUARANTEED_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION")
	// TransactionRejectReasonGuaranteedStopLossOnFillRequiredForPendingOrder means an attempt to create a pending Order was made with no Guaranteed Stop Loss Order on fill specified and the Account’s configuration requires that every Trade have an associated Guaranteed Stop Loss Order.
	TransactionRejectReasonGuaranteedStopLossOnFillRequiredForPendingOrder = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED_FOR_PENDING_ORDER")
	// TransactionRejectReasonGuaranteedStopLossOnFillNotAllowed means an attempt to create a pending Order was made with a Guaranteed Stop Loss Order on fill that was explicitly configured to be guaranteed, however the Account’s configuration does not allow guaranteed Stop Loss Orders.
	TransactionRejectReasonGuaranteedStopLossOnFillNotAllowed = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED")
	// TransactionRejectReasonGuaranteedStopLossOnFillRequired means an attempt to create a pending Order was made with a Guaranteed Stop Loss Order on fill that was explicitly configured to be not guaranteed, however the Account’s configuration requires Guaranteed Stop Loss Orders.
	TransactionRejectReasonGuaranteedStopLossOnFillRequired = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED")
	// TransactionRejectReasonGuaranteedStopLossOnFillPriceMissing means the Guaranteed Stop Loss on fill specified does not provide a price
	TransactionRejectReasonGuaranteedStopLossOnFillPriceMissing = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_MISSING")
	// TransactionRejectReasonGuaranteedStopLossOnFillPriceInvalid means the Guaranteed Stop Loss on fill specifies an invalid price
	TransactionRejectReasonGuaranteedStopLossOnFillPriceInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOnFillPricePrecisionExceeded means the Guaranteed Stop Loss on fill specifies a price with more precision than is allowed by the Order’s instrument
	TransactionRejectReasonGuaranteedStopLossOnFillPricePrecisionExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_PRECISION_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOnFillMinimumDistanceNotMet means an attempt to create a pending Order was made with the distance between the Guaranteed Stop Loss Order on fill’s price and the pending Order’s price is less than the Account’s configured minimum guaranteed stop loss distance.
	TransactionRejectReasonGuaranteedStopLossOnFillMinimumDistanceNotMet = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET")
	// TransactionRejectReasonGuaranteedStopLossOnFillLevelRestrictionVolumeExceeded means filling the Order would result in the creation of a Guaranteed Stop Loss Order with trigger number of units that violates the account’s Guaranteed Stop Loss Order level restriction volume.
	TransactionRejectReasonGuaranteedStopLossOnFillLevelRestrictionVolumeExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOnFillLevelRestrictionPriceRangeExceeded means filling the Order would result in the creation of a Guaranteed Stop Loss Order with trigger price that violates the account’s Guaranteed Stop Loss Order level restriction price range.
	TransactionRejectReasonGuaranteedStopLossOnFillLevelRestrictionPriceRangeExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOnFillDistanceInvalid means the Guaranteed Stop Loss on fill distance is invalid
	TransactionRejectReasonGuaranteedStopLossOnFillDistanceInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_DISTANCE_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOnFillPriceDistanceMaximumExceeded means the Guaranteed Stop Loss on fill price distance exceeds the maximum allowed amount.
	TransactionRejectReasonGuaranteedStopLossOnFillPriceDistanceMaximumExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOnFillDistancePrecisionExceeded means the Guaranteed Stop Loss on fill distance contains more precision than is allowed by the instrument
	TransactionRejectReasonGuaranteedStopLossOnFillDistancePrecisionExceeded = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_DISTANCE_PRECISION_EXCEEDED")
	// TransactionRejectReasonGuaranteedStopLossOnFillPriceAndDistanceBothSpecified means the Guaranteed Stop Loss on fill contains both the price and distance fields.
	TransactionRejectReasonGuaranteedStopLossOnFillPriceAndDistanceBothSpecified = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_SPECIFIED")
	// TransactionRejectReasonGuaranteedStopLossOnFillPriceAndDistanceBothMissing means the Guaranteed Stop Loss on fill contains neither the price nor distance fields.
	TransactionRejectReasonGuaranteedStopLossOnFillPriceAndDistanceBothMissing = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_MISSING")
	// TransactionRejectReasonGuaranteedStopLossOnFillTimeInForceMissing means the Guaranteed Stop Loss on fill specified does not provide a TimeInForce
	TransactionRejectReasonGuaranteedStopLossOnFillTimeInForceMissing = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING")
	// TransactionRejectReasonGuaranteedStopLossOnFillTimeInForceInvalid means the Guaranteed Stop Loss on fill specifies an invalid TimeInForce
	TransactionRejectReasonGuaranteedStopLossOnFillTimeInForceInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOnFillGtdTimestampMissing means the Guaranteed Stop Loss on fill specifies a GTD TimeInForce but does not provide a GTD timestamp
	TransactionRejectReasonGuaranteedStopLossOnFillGtdTimestampMissing = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING")
	// TransactionRejectReasonGuaranteedStopLossOnFillGtdTimestampInPast means the Guaranteed Stop Loss on fill specifies a GTD timestamp that is in the past.
	TransactionRejectReasonGuaranteedStopLossOnFillGtdTimestampInPast = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST")
	// TransactionRejectReasonGuaranteedStopLossOnFillClientOrderIdInvalid means the Guaranteed Stop Loss on fill client Order ID specified is invalid
	TransactionRejectReasonGuaranteedStopLossOnFillClientOrderIdInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOnFillClientOrderTagInvalid means the Guaranteed Stop Loss on fill client Order tag specified is invalid
	TransactionRejectReasonGuaranteedStopLossOnFillClientOrderTagInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOnFillClientOrderCommentInvalid means the Guaranteed Stop Loss on fill client Order comment specified is invalid.
	TransactionRejectReasonGuaranteedStopLossOnFillClientOrderCommentInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID")
	// TransactionRejectReasonGuaranteedStopLossOnFillTriggerConditionMissing means the Guaranteed Stop Loss on fill specified does not provide a TriggerCondition.
	TransactionRejectReasonGuaranteedStopLossOnFillTriggerConditionMissing = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING")
	// TransactionRejectReasonGuaranteedStopLossOnFillTriggerConditionInvalid means the Guaranteed Stop Loss on fill specifies an invalid TriggerCondition.
	TransactionRejectReasonGuaranteedStopLossOnFillTriggerConditionInvalid = TransactionRejectReason("GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID")
	// TransactionRejectReasonTrailingStopLossOrderAlreadyExists means a Trailing Stop Loss Order for the specified Trade already exists
	TransactionRejectReasonTrailingStopLossOrderAlreadyExists = TransactionRejectReason("TRAILING_STOP_LOSS_ORDER_ALREADY_EXISTS")
	// TransactionRejectReasonTrailingStopLossOrderWouldViolateFifoViolationSafeguard means the Trailing Stop Loss Order would cause the associated Trade to be in violation of the FIFO violation safeguard constraints
	TransactionRejectReasonTrailingStopLossOrderWouldViolateFifoViolationSafeguard = TransactionRejectReason("TRAILING_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD")
	// TransactionRejectReasonTrailingStopLossOrderRmoMutualExclusivityMutuallyExclusiveViolation means the Trailing Stop Loss Order would be in violation of the risk management Order mutual exclusivity configuration specifying that only one risk management order can be attached to a Trade.
	TransactionRejectReasonTrailingStopLossOrderRmoMutualExclusivityMutuallyExclusiveViolation = TransactionRejectReason("TRAILING_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION")
	// TransactionRejectReasonTrailingStopLossOrderRmoMutualExclusivityGsloExcludesOthersViolation means the Trailing Stop Loss Order would be in violation of the risk management Order mutual exclusivity configuration specifying that if a GSLO is already attached to a Trade, no other risk management Order can be attached to the same Trade.
	TransactionRejectReasonTrailingStopLossOrderRmoMutualExclusivityGsloExcludesOthersViolation = TransactionRejectReason("TRAILING_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION")
	// TransactionRejectReasonTrailingStopLossOnFillPriceDistanceMissing means the Trailing Stop Loss on fill specified does not provide a distance
	TransactionRejectReasonTrailingStopLossOnFillPriceDistanceMissing = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MISSING")
	// TransactionRejectReasonTrailingStopLossOnFillPriceDistanceInvalid means the Trailing Stop Loss on fill distance is invalid
	TransactionRejectReasonTrailingStopLossOnFillPriceDistanceInvalid = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_INVALID")
	// TransactionRejectReasonTrailingStopLossOnFillPriceDistancePrecisionExceeded means the Trailing Stop Loss on fill distance contains more precision than is allowed by the instrument
	TransactionRejectReasonTrailingStopLossOnFillPriceDistancePrecisionExceeded = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_PRECISION_EXCEEDED")
	// TransactionRejectReasonTrailingStopLossOnFillPriceDistanceMaximumExceeded means the Trailing Stop Loss on fill price distance exceeds the maximum allowed amount
	TransactionRejectReasonTrailingStopLossOnFillPriceDistanceMaximumExceeded = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED")
	// TransactionRejectReasonTrailingStopLossOnFillPriceDistanceMinimumNotMet means the Trailing Stop Loss on fill price distance does not meet the minimum allowed amount
	TransactionRejectReasonTrailingStopLossOnFillPriceDistanceMinimumNotMet = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MINIMUM_NOT_MET")
	// TransactionRejectReasonTrailingStopLossOnFillTimeInForceMissing means the Trailing Stop Loss on fill specified does not provide a TimeInForce
	TransactionRejectReasonTrailingStopLossOnFillTimeInForceMissing = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING")
	// TransactionRejectReasonTrailingStopLossOnFillTimeInForceInvalid means the Trailing Stop Loss on fill specifies an invalid TimeInForce
	TransactionRejectReasonTrailingStopLossOnFillTimeInForceInvalid = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID")
	// TransactionRejectReasonTrailingStopLossOnFillGtdTimestampMissing means the Trailing Stop Loss on fill TimeInForce is specified as GTD but no GTD timestamp is provided
	TransactionRejectReasonTrailingStopLossOnFillGtdTimestampMissing = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING")
	// TransactionRejectReasonTrailingStopLossOnFillGtdTimestampInPast means the Trailing Stop Loss on fill GTD timestamp is in the past
	TransactionRejectReasonTrailingStopLossOnFillGtdTimestampInPast = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST")
	// TransactionRejectReasonTrailingStopLossOnFillClientOrderIdInvalid means the Trailing Stop Loss on fill client Order ID specified is invalid
	TransactionRejectReasonTrailingStopLossOnFillClientOrderIdInvalid = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID")
	// TransactionRejectReasonTrailingStopLossOnFillClientOrderTagInvalid means the Trailing Stop Loss on fill client Order tag specified is invalid
	TransactionRejectReasonTrailingStopLossOnFillClientOrderTagInvalid = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID")
	// TransactionRejectReasonTrailingStopLossOnFillClientOrderCommentInvalid means the Trailing Stop Loss on fill client Order comment specified is invalid
	TransactionRejectReasonTrailingStopLossOnFillClientOrderCommentInvalid = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID")
	// TransactionRejectReasonTrailingStopLossOrdersNotSupported means a client attempted to create either a Trailing Stop Loss order or an order with a Trailing Stop Loss On Fill specified, which may not yet be supported.
	TransactionRejectReasonTrailingStopLossOrdersNotSupported = TransactionRejectReason("TRAILING_STOP_LOSS_ORDERS_NOT_SUPPORTED")
	// TransactionRejectReasonTrailingStopLossOnFillTriggerConditionMissing means the Trailing Stop Loss on fill specified does not provide a TriggerCondition
	TransactionRejectReasonTrailingStopLossOnFillTriggerConditionMissing = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING")
	// TransactionRejectReasonTrailingStopLossOnFillTriggerConditionInvalid means the Tailing Stop Loss on fill specifies an invalid TriggerCondition
	TransactionRejectReasonTrailingStopLossOnFillTriggerConditionInvalid = TransactionRejectReason("TRAILING_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID")
	// TransactionRejectReasonCloseTradeTypeMissing means the request to close a Trade does not specify a full or partial close
	TransactionRejectReasonCloseTradeTypeMissing = TransactionRejectReason("CLOSE_TRADE_TYPE_MISSING")
	// TransactionRejectReasonCloseTradePartialUnitsMissing means the request to close a Trade partially did not specify the number of units to close
	TransactionRejectReasonCloseTradePartialUnitsMissing = TransactionRejectReason("CLOSE_TRADE_PARTIAL_UNITS_MISSING")
	// TransactionRejectReasonCloseTradeUnitsExceedTradeSize means the request to partially close a Trade specifies a number of units that exceeds the current size of the given Trade
	TransactionRejectReasonCloseTradeUnitsExceedTradeSize = TransactionRejectReason("CLOSE_TRADE_UNITS_EXCEED_TRADE_SIZE")
	// TransactionRejectReasonCloseoutPositionDoesntExist means the Position requested to be closed out does not exist
	TransactionRejectReasonCloseoutPositionDoesntExist = TransactionRejectReason("CLOSEOUT_POSITION_DOESNT_EXIST")
	// TransactionRejectReasonCloseoutPositionIncompleteSpecification means the request to closeout a Position was specified incompletely
	TransactionRejectReasonCloseoutPositionIncompleteSpecification = TransactionRejectReason("CLOSEOUT_POSITION_INCOMPLETE_SPECIFICATION")
	// TransactionRejectReasonCloseoutPositionUnitsExceedPositionSize means a partial Position closeout request specifies a number of units that exceeds the current Position
	TransactionRejectReasonCloseoutPositionUnitsExceedPositionSize = TransactionRejectReason("CLOSEOUT_POSITION_UNITS_EXCEED_POSITION_SIZE")
	// TransactionRejectReasonCloseoutPositionReject means the request to closeout a Position could not be fully satisfied
	TransactionRejectReasonCloseoutPositionReject = TransactionRejectReason("CLOSEOUT_POSITION_REJECT")
	// TransactionRejectReasonCloseoutPositionPartialUnitsMissing means the request to partially closeout a Position did not specify the number of units to close.
	TransactionRejectReasonCloseoutPositionPartialUnitsMissing = TransactionRejectReason("CLOSEOUT_POSITION_PARTIAL_UNITS_MISSING")
	// TransactionRejectReasonMarkupGroupIdInvalid means the markup group ID provided is invalid
	TransactionRejectReasonMarkupGroupIdInvalid = TransactionRejectReason("MARKUP_GROUP_ID_INVALID")
	// TransactionRejectReasonPositionAggregationModeInvalid means the PositionAggregationMode provided is not supported/valid.
	TransactionRejectReasonPositionAggregationModeInvalid = TransactionRejectReason("POSITION_AGGREGATION_MODE_INVALID")
	// TransactionRejectReasonAdminConfigureDataMissing means no configuration parameters provided
	TransactionRejectReasonAdminConfigureDataMissing = TransactionRejectReason("ADMIN_CONFIGURE_DATA_MISSING")
	// TransactionRejectReasonMarginRateInvalid means the margin rate provided is invalid
	TransactionRejectReasonMarginRateInvalid = TransactionRejectReason("MARGIN_RATE_INVALID")
	// TransactionRejectReasonMarginRateWouldTriggerCloseout means the margin rate provided would cause an immediate margin closeout
	TransactionRejectReasonMarginRateWouldTriggerCloseout = TransactionRejectReason("MARGIN_RATE_WOULD_TRIGGER_CLOSEOUT")
	// TransactionRejectReasonAliasInvalid means the account alias string provided is invalid
	TransactionRejectReasonAliasInvalid = TransactionRejectReason("ALIAS_INVALID")
	// TransactionRejectReasonClientConfigureDataMissing means no configuration parameters provided
	TransactionRejectReasonClientConfigureDataMissing = TransactionRejectReason("CLIENT_CONFIGURE_DATA_MISSING")
	// TransactionRejectReasonMarginRateWouldTriggerMarginCall means the margin rate provided would cause the Account to enter a margin call state.
	TransactionRejectReasonMarginRateWouldTriggerMarginCall = TransactionRejectReason("MARGIN_RATE_WOULD_TRIGGER_MARGIN_CALL")
	// TransactionRejectReasonAmountInvalid means funding is not possible because the requested transfer amount is invalid
	TransactionRejectReasonAmountInvalid = TransactionRejectReason("AMOUNT_INVALID")
	// TransactionRejectReasonInsufficientFunds means the Account does not have sufficient balance to complete the funding request
	TransactionRejectReasonInsufficientFunds = TransactionRejectReason("INSUFFICIENT_FUNDS")
	// TransactionRejectReasonAmountMissing means funding amount has not been specified
	TransactionRejectReasonAmountMissing = TransactionRejectReason("AMOUNT_MISSING")
	// TransactionRejectReasonFundingReasonMissing means funding reason has not been specified
	TransactionRejectReasonFundingReasonMissing = TransactionRejectReason("FUNDING_REASON_MISSING")
	// TransactionRejectReasonOcaOrderIdsStopLossNotAllowed means the list of Order Identifiers provided for a One Cancels All Order contains an Order Identifier that refers to a Stop Loss Order. OCA groups cannot contain Stop Loss Orders.
	TransactionRejectReasonOcaOrderIdsStopLossNotAllowed = TransactionRejectReason("OCA_ORDER_IDS_STOP_LOSS_NOT_ALLOWED")
	// TransactionRejectReasonClientExtensionsDataMissing means neither Order nor Trade on Fill client extensions were provided for modification
	TransactionRejectReasonClientExtensionsDataMissing = TransactionRejectReason("CLIENT_EXTENSIONS_DATA_MISSING")
	// TransactionRejectReasonReplacingOrderInvalid means the Order to be replaced has a different type than the replacing Order.
	TransactionRejectReasonReplacingOrderInvalid = TransactionRejectReason("REPLACING_ORDER_INVALID")
	// TransactionRejectReasonReplacingTradeIdInvalid means the replacing Order refers to a different Trade than the Order that is being replaced.
	TransactionRejectReasonReplacingTradeIdInvalid = TransactionRejectReason("REPLACING_TRADE_ID_INVALID")
	// TransactionRejectReasonOrderCancelWouldTriggerCloseout means canceling the order would cause an immediate margin closeout.
	TransactionRejectReasonOrderCancelWouldTriggerCloseout = TransactionRejectReason("ORDER_CANCEL_WOULD_TRIGGER_CLOSEOUT")
)

// TransactionFilter as a filter that can be used when fetching Transactions
type TransactionFilter string

const (
	// TransactionFilterOrder filters order-related Transactions. These are the Transactions that create, cancel, fill or trigger Orders
	TransactionFilterOrder = TransactionFilter("ORDER")
	// TransactionFilterFunding filters funding-related Transactions
	TransactionFilterFunding = TransactionFilter("FUNDING")
	// TransactionFilterAdmin filters administrative Transactions
	TransactionFilterAdmin = TransactionFilter("ADMIN")
	// TransactionFilterCreate filters account Create Transaction
	TransactionFilterCreate = TransactionFilter("CREATE")
	// TransactionFilterClose filters account Close Transaction
	TransactionFilterClose = TransactionFilter("CLOSE")
	// TransactionFilterReopen filters account Reopen Transaction
	TransactionFilterReopen = TransactionFilter("REOPEN")
	// TransactionFilterClientConfigure filters client Configuration Transaction
	TransactionFilterClientConfigure = TransactionFilter("CLIENT_CONFIGURE")
	// TransactionFilterClientConfigureReject filters client Configuration Reject Transaction
	TransactionFilterClientConfigureReject = TransactionFilter("CLIENT_CONFIGURE_REJECT")
	// TransactionFilterTransferFunds filters transfer Funds Transaction
	TransactionFilterTransferFunds = TransactionFilter("TRANSFER_FUNDS")
	// TransactionFilterTransferFundsReject filters transfer Funds Reject Transaction
	TransactionFilterTransferFundsReject = TransactionFilter("TRANSFER_FUNDS_REJECT")
	// TransactionFilterMarketOrder filters market Order Transaction
	TransactionFilterMarketOrder = TransactionFilter("MARKET_ORDER")
	// TransactionFilterMarketOrderReject filters market Order Reject Transaction
	TransactionFilterMarketOrderReject = TransactionFilter("MARKET_ORDER_REJECT")
	// TransactionFilterLimitOrder filters limit Order Transaction
	TransactionFilterLimitOrder = TransactionFilter("LIMIT_ORDER")
	// TransactionFilterLimitOrderReject filters limit Order Reject Transaction
	TransactionFilterLimitOrderReject = TransactionFilter("LIMIT_ORDER_REJECT")
	// TransactionFilterStopOrder filters stop Order Transaction
	TransactionFilterStopOrder = TransactionFilter("STOP_ORDER")
	// TransactionFilterStopOrderReject filters stop Order Reject Transaction
	TransactionFilterStopOrderReject = TransactionFilter("STOP_ORDER_REJECT")
	// TransactionFilterMarketIfTouchedOrder filters market if Touched Order Transaction
	TransactionFilterMarketIfTouchedOrder = TransactionFilter("MARKET_IF_TOUCHED_ORDER")
	// TransactionFilterMarketIfTouchedOrderReject filters market if Touched Order Reject Transaction
	TransactionFilterMarketIfTouchedOrderReject = TransactionFilter("MARKET_IF_TOUCHED_ORDER_REJECT")
	// TransactionFilterTakeProfitOrder filters take Profit Order Transaction
	TransactionFilterTakeProfitOrder = TransactionFilter("TAKE_PROFIT_ORDER")
	// TransactionFilterTakeProfitOrderReject filters take Profit Order Reject Transaction
	TransactionFilterTakeProfitOrderReject = TransactionFilter("TAKE_PROFIT_ORDER_REJECT")
	// TransactionFilterStopLossOrder filters stop Loss Order Transaction
	TransactionFilterStopLossOrder = TransactionFilter("STOP_LOSS_ORDER")
	// TransactionFilterStopLossOrderReject filters stop Loss Order Reject Transaction
	TransactionFilterStopLossOrderReject = TransactionFilter("STOP_LOSS_ORDER_REJECT")
	// TransactionFilterGuaranteedStopLossOrder filters guaranteed Stop Loss Order Transaction
	TransactionFilterGuaranteedStopLossOrder = TransactionFilter("GUARANTEED_STOP_LOSS_ORDER")
	// TransactionFilterGuaranteedStopLossOrderReject filters guaranteed Stop Loss Order Reject Transaction
	TransactionFilterGuaranteedStopLossOrderReject = TransactionFilter("GUARANTEED_STOP_LOSS_ORDER_REJECT")
	// TransactionFilterTrailingStopLossOrder filters trailing Stop Loss Order Transaction
	TransactionFilterTrailingStopLossOrder = TransactionFilter("TRAILING_STOP_LOSS_ORDER")
	// TransactionFilterTrailingStopLossOrderReject filters trailing Stop Loss Order Reject Transaction
	TransactionFilterTrailingStopLossOrderReject = TransactionFilter("TRAILING_STOP_LOSS_ORDER_REJECT")
	// TransactionFilterOneCancelsAllOrder filters one Cancels All Order Transaction
	TransactionFilterOneCancelsAllOrder = TransactionFilter("ONE_CANCELS_ALL_ORDER")
	// TransactionFilterOneCancelsAllOrderReject filters one Cancels All Order Reject Transaction
	TransactionFilterOneCancelsAllOrderReject = TransactionFilter("ONE_CANCELS_ALL_ORDER_REJECT")
	// TransactionFilterOneCancelsAllOrderTriggered filters one Cancels All Order Trigger Transaction
	TransactionFilterOneCancelsAllOrderTriggered = TransactionFilter("ONE_CANCELS_ALL_ORDER_TRIGGERED")
	// TransactionFilterOrderFill filters order Fill Transaction
	TransactionFilterOrderFill = TransactionFilter("ORDER_FILL")
	// TransactionFilterOrderCancel filters order Cancel Transaction
	TransactionFilterOrderCancel = TransactionFilter("ORDER_CANCEL")
	// TransactionFilterOrderCancelReject filters order Cancel Reject Transaction
	TransactionFilterOrderCancelReject = TransactionFilter("ORDER_CANCEL_REJECT")
	// TransactionFilterOrderClientExtensionsModify filters order Client Extensions Modify Transaction
	TransactionFilterOrderClientExtensionsModify = TransactionFilter("ORDER_CLIENT_EXTENSIONS_MODIFY")
	// TransactionFilterOrderClientExtensionsModifyReject filters order Client Extensions Modify Reject Transaction
	TransactionFilterOrderClientExtensionsModifyReject = TransactionFilter("ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT")
	// TransactionFilterTradeClientExtensionsModify filters trade Client Extensions Modify Transaction
	TransactionFilterTradeClientExtensionsModify = TransactionFilter("TRADE_CLIENT_EXTENSIONS_MODIFY")
	// TransactionFilterTradeClientExtensionsModifyReject filters trade Client Extensions Modify Reject Transaction
	TransactionFilterTradeClientExtensionsModifyReject = TransactionFilter("TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT")
	// TransactionFilterMarginCallEnter filters margin Call Enter Transaction
	TransactionFilterMarginCallEnter = TransactionFilter("MARGIN_CALL_ENTER")
	// TransactionFilterMarginCallExtend filters margin Call Extend Transaction
	TransactionFilterMarginCallExtend = TransactionFilter("MARGIN_CALL_EXTEND")
	// TransactionFilterMarginCallExit filters margin Call Exit Transaction
	TransactionFilterMarginCallExit = TransactionFilter("MARGIN_CALL_EXIT")
	// TransactionFilterDelayedTradeClosure filters delayed Trade Closure Transaction
	TransactionFilterDelayedTradeClosure = TransactionFilter("DELAYED_TRADE_CLOSURE")
	// TransactionFilterDailyFinancing filters daily Financing Transaction
	TransactionFilterDailyFinancing = TransactionFilter("DAILY_FINANCING")
	// TransactionFilterResetResettablePl filters reset Resettable PL Transaction
	TransactionFilterResetResettablePl = TransactionFilter("RESET_RESETTABLE_PL")
)

// TransactionHeartbeat object is injected into the Transaction stream to ensure that the HTTP connection remains active.
type TransactionHeartbeat struct {
	// The string "HEARTBEAT
	Type string `json:"type"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The date/time when the TransactionHeartbeat was created.
	Time time.Time `json:"time"`
}
