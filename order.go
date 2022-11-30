package oanda_sdk

import (
	"github.com/shopspring/decimal"
	"time"
)

// Order is an interface that defines properties common to all Orders
type Order interface {
	// GetId returns the Order’s identifier, unique within the Order’s Account.
	GetId() string

	// GetCreateTime returns the time when the Order was created.
	GetCreateTime() time.Time

	// GetState returns the current state of the Order
	GetState() OrderState

	// GetClientExtensions returns the client extensions of the Order.
	// Do not set, modify, or delete clientExtensions if your account is associated with MT4
	GetClientExtensions() ClientExtensions

	// GetType returns the type of the Order
	GetType() OrderType
}

// MarketOrder is an order that is filled immediately upon creation using the current market price
type MarketOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “MARKET” for Market Orders.
	Type OrderType `json:"type"`

	// The Market Order’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the Market Order. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The time-in-force requested for the Market Order. Restricted to FOK or IOC for a MarketOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The worst price that the client is willing to have the Market Order filled at.
	PriceBound decimal.Decimal `json:"priceBound"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

	// Details of the Trade requested to be closed, only provided when the Market Order is being used to
	// explicitly close a Trade.
	TradeClose MarketOrderTradeClose `json:"tradeClose"`

	// Details of the long Position requested to be closed out, only provided when a Market Order is being used to
	// explicitly closeout a long Position.
	LongPositionCloseout MarketOrderPositionCloseout `json:"longPositionCloseout"`

	// Details of the short Position requested to be closed out, only provided when a Market Order is being used to
	// explicitly closeout a short Position.
	ShortPositionCloseout MarketOrderPositionCloseout `json:"shortPositionCloseout"`

	// Details of the Margin Closeout that this Market Order was created for
	MarginCloseout MarketOrderMarginCloseout `json:"marginCloseout"`

	// Details of the delayed Trade close that this Market Order was created for
	DelayedTradeClose MarketOrderDelayedTradeClose `json:"delayedTradeClose"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
	// the Trade.
	GuaranteedStopLossOnFill GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`
}

func (mo MarketOrder) GetId() string {
	return mo.Id
}

func (mo MarketOrder) GetCreateTime() time.Time {
	return mo.CreateTime
}

func (mo MarketOrder) GetState() OrderState {
	return mo.State
}

func (mo MarketOrder) GetClientExtensions() ClientExtensions {
	return mo.ClientExtensions
}

func (mo MarketOrder) GetType() OrderType {
	return mo.Type
}

// FixedPriceOrder is an order that is filled immediately upon creation using a fixed price
type FixedPriceOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to "FIXED_PRICE" for Market Orders.
	Type OrderType `json:"type"`

	// The Fixed Price Order’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the FixedPriceOrder. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price specified for the Fixed Price Order. This price is the exact price that the FixedPriceOrder
	// will be filled at.
	Price decimal.Decimal `json:"price"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

	// The state that the trade resulting from the FixedPriceOrder should be set to.
	TradeState string `json:"tradeState"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
	// the Trade.
	GuaranteedStopLossOnFill GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`
}

func (fpo FixedPriceOrder) GetId() string {
	return fpo.Id
}

func (fpo FixedPriceOrder) GetCreateTime() time.Time {
	return fpo.CreateTime
}

func (fpo FixedPriceOrder) GetState() OrderState {
	return fpo.State
}

func (fpo FixedPriceOrder) GetClientExtensions() ClientExtensions {
	return fpo.ClientExtensions
}

func (fpo FixedPriceOrder) GetType() OrderType {
	return fpo.Type
}

type LimitOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “LIMIT” for Market Orders.
	Type OrderType `json:"type"`

	// The Market Order’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the Market Order. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price threshold specified for the LimitOrder. The LimitOrder will
	// only be filled by a market price that is equal to or better than this price.
	Price decimal.Decimal `json:"price"`

	// The time-in-force requested for the LimitOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the Limit Order will be cancelled if its timeInForce is “GTD”.
	GtdTime time.Time `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// Details of the Trade requested to be closed, only provided when the Market Order is being used to
	// explicitly close a Trade.
	TradeClose MarketOrderTradeClose `json:"tradeClose"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
	// the Trade.
	GuaranteedStopLossOnFill GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (lo LimitOrder) GetId() string {
	return lo.Id
}

func (lo LimitOrder) GetCreateTime() time.Time {
	return lo.CreateTime
}

func (lo LimitOrder) GetState() OrderState {
	return lo.State
}

func (lo LimitOrder) GetClientExtensions() ClientExtensions {
	return lo.ClientExtensions
}

func (lo LimitOrder) GetType() OrderType {
	return lo.Type
}

// StopOrder is an order that is created with a price threshold, and will only be filled by a price that is equal to or
// worse than the threshold.
type StopOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “STOP” for StopOrder.
	Type OrderType `json:"type"`

	// The StopOrder’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the StopOrder. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price threshold specified for the StopOrder. The StopOrder will
	// only be filled by a market price that is equal to or worse than this price.
	Price decimal.Decimal `json:"price"`

	// The worst market price that may be used to fill this StopOrder. If the
	// market gaps and crosses through both the price and the priceBound, the
	// StopOrder will be cancelled instead of being filled.
	PriceBound decimal.Decimal `json:"priceBound"`

	// The time-in-force requested for the StopOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the Stop Order will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// TakeProfitDetails specifies the details of a TakeProfitOrder to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// TakeProfitOrder is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a StopLossOrder to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent StopLossOrder
	// is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// GuaranteedStopLossDetails specifies the details of a GuaranteedStopLossOrder
	// to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent GuaranteedStopLossOrder is modified directly through
	// the Trade.
	GuaranteedStopLossOnFill GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a TrailingStopLossOrder
	// to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent TrailingStopLossOrder is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (so StopOrder) GetId() string {
	return so.Id
}

func (so StopOrder) GetCreateTime() time.Time {
	return so.CreateTime
}

func (so StopOrder) GetState() OrderState {
	return so.State
}

func (so StopOrder) GetClientExtensions() ClientExtensions {
	return so.ClientExtensions
}

func (so StopOrder) GetType() OrderType {
	return so.Type
}

// MarketIfTouchedOrder is an order that is created with a price threshold, and will only be filled by a market price
// that touches or crosses the threshold.
type MarketIfTouchedOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “MARKET_IF_TOUCHED” for MarketIfTouchedOrder.
	Type OrderType `json:"type"`

	// The MarketIfTouchedOrder’s Instrument.
	Instrument string `json:"instrument"`

	// The quantity requested to be filled by the MarketIfTouchedOrder. A positive number of units results in a long Order,
	// and a negative number of units results in a short Order.
	Units decimal.Decimal `json:"units"`

	// The price threshold specified for the MarketIfTouchedOrder. The
	// MarketIfTouchedOrder will only be filled by a market price that crosses
	// this price from the direction of the market price at the time when the
	// Order was created (the initialMarketPrice). Depending on the value of the
	// Order’s price and initialMarketPrice, the MarketIfTouchedOrder will
	// behave like a LimitOrder or a StopOrder.
	Price decimal.Decimal `json:"price"`

	// The worst market price that may be used to fill this MarketIfTouchedOrder.
	PriceBound decimal.Decimal `json:"priceBound"`

	// The time-in-force requested for the MarketIfTouched Order. Restricted to “GTC”, “GFD” and “GTD” for MarketIfTouchedOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

	// The date/time when the MarketIfTouched Order will be cancelled if its timeInForce is “GTD”.
	GtdTime *time.Time `json:"gtdTime"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The Market price at the time when the MarketIfTouchedOrder was created.
	InitialMarketPrice decimal.Decimal `json:"initialMarketPrice"`

	// TakeProfitDetails specifies the details of a TakeProfitOrder to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// TakeProfitOrder is modified directly through the Trade.
	TakeProfitOnFill TakeProfitDetails `json:"takeProfitOnFill"`

	// StopLossDetails specifies the details of a StopLossOrder to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent StopLossOrder
	// is modified directly through the Trade.
	StopLossOnFill StopLossDetails `json:"stopLossOnFill"`

	// GuaranteedStopLossDetails specifies the details of a GuaranteedStopLossOrder
	// to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent GuaranteedStopLossOrder is modified directly through
	// the Trade.
	GuaranteedStopLossOnFill GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`

	// TrailingStopLossDetails specifies the details of a TrailingStopLossOrder
	// to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent TrailingStopLossOrder is modified directly through
	// the Trade.
	TrailingStopLossOnFill TrailingStopLossDetails `json:"trailingStopLossOnFill"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions ClientExtensions `json:"tradeClientExtensions"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (mito MarketIfTouchedOrder) GetId() string {
	return mito.Id
}

func (mito MarketIfTouchedOrder) GetCreateTime() time.Time {
	return mito.CreateTime
}

func (mito MarketIfTouchedOrder) GetState() OrderState {
	return mito.State
}

func (mito MarketIfTouchedOrder) GetClientExtensions() ClientExtensions {
	return mito.ClientExtensions
}

func (mito MarketIfTouchedOrder) GetType() OrderType {
	return mito.Type
}

// TakeProfitOrder is an order that is linked to an open Trade and created with a price threshold. The Order will be
// filled (closing the Trade) by the first price that is equal to or better than the threshold. A TakeProfitOrder
// cannot be used to open a new Position.
type TakeProfitOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “TAKE_PROFIT” for TakeProfitOrder.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the TakeProfitOrder. The associated
	// Trade will be closed by a market price that is equal to or better than this threshold.
	Price decimal.Decimal `json:"price"`

	// The time-in-force requested for the TakeProfitOrder. Restricted to “GTC”, “GFD” and “GTD” for TakeProfitOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (tpo TakeProfitOrder) GetId() string {
	return tpo.Id
}

func (tpo TakeProfitOrder) GetCreateTime() time.Time {
	return tpo.CreateTime
}

func (tpo TakeProfitOrder) GetState() OrderState {
	return tpo.State
}

func (tpo TakeProfitOrder) GetClientExtensions() ClientExtensions {
	return tpo.ClientExtensions
}

func (tpo TakeProfitOrder) GetType() OrderType {
	return tpo.Type
}

// StopLossOrder is an order that is linked to an open Trade and created with a price threshold. The Order will be
// filled (closing the Trade) by the first price that is equal to or worse than the threshold. A StopLossOrder
// cannot be used to open a new Position.
type StopLossOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “STOP_LOSS” for StopLossOrder.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the StopLossOrder. The associated
	// Trade will be closed by a market price that is equal to or worse than this threshold.
	Price decimal.Decimal `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the StopLossOrder price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance decimal.Decimal `json:"distance"`

	// The time-in-force requested for the StopLossOrder. Restricted to “GTC”, “GFD” and “GTD” for StopLossOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (slo StopLossOrder) GetId() string {
	return slo.Id
}

func (slo StopLossOrder) GetCreateTime() time.Time {
	return slo.CreateTime
}

func (slo StopLossOrder) GetState() OrderState {
	return slo.State
}

func (slo StopLossOrder) GetClientExtensions() ClientExtensions {
	return slo.ClientExtensions
}

func (slo StopLossOrder) GetType() OrderType {
	return slo.Type
}

// GuaranteedStopLossOrder is an order that is linked to an open Trade and created with a price threshold which is
// guaranteed against slippage that may occur as the market crosses the price set for that order. The Order will be
// filled (closing the Trade) by the first price that is equal to or worse than the threshold. The price level specified
// for the GuaranteedStopLossOrder must be at least the configured minimum distance (in price units) away from the entry
// price for the traded instrument. A GuaranteedStopLossOrder cannot be used to open a new Position.
type GuaranteedStopLossOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “GUARANTEED_STOP_LOSS” for GuaranteedStopLossOrder.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price threshold specified for the GuaranteedStopLossOrder. The associated
	// Trade will be closed by a market price that is equal to or worse than this threshold.
	Price decimal.Decimal `json:"price"`

	// Specifies the distance (in price units) from the Account’s current price
	// to use as the GuaranteedStopLossOrder price. If the Trade is short the
	// Instrument’s bid price is used, and for long Trades the ask is used.
	Distance decimal.Decimal `json:"distance"`

	// The time-in-force requested for the GuaranteedStopLossOrder. Restricted to “GTC”, “GFD” and “GTD” for
	// GuaranteedStopLossOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (gslo GuaranteedStopLossOrder) GetId() string {
	return gslo.Id
}

func (gslo GuaranteedStopLossOrder) GetCreateTime() time.Time {
	return gslo.CreateTime
}

func (gslo GuaranteedStopLossOrder) GetState() OrderState {
	return gslo.State
}

func (gslo GuaranteedStopLossOrder) GetClientExtensions() ClientExtensions {
	return gslo.ClientExtensions
}

func (gslo GuaranteedStopLossOrder) GetType() OrderType {
	return gslo.Type
}

// TrailingStopLossOrder is an order that is linked to an open Trade and created with a price distance. The price
// distance is used to calculate a trailing stop value for the order that is in the losing direction from the market
// price at the time of the order's creation. The trailing stop value will follow the market price as it moves in the
// winning direction, and the order will be filled (closing the Trade) by the first price that is equal to or worse than
// the trailing stop value. A TrailingStopLossOrder cannot be used to open a new Position.
type TrailingStopLossOrder struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id string `json:"id"`

	// The time when the Order was created.
	CreateTime time.Time `json:"createTime"`

	// The current state of the Order.
	State OrderState `json:"state"`

	// The client extensions of the Order. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	ClientExtensions ClientExtensions `json:"clientExtensions"`

	// The type of the Order. Always set to “TRAILING_STOP_LOSS” for TrailingStopLossOrder.
	Type OrderType `json:"type"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`

	// The price distance (in price units) specified for the TrailingStopLossOrder.
	Distance decimal.Decimal `json:"distance"`

	// The time-in-force requested for the TrailingStopLossOrder. Restricted to “GTC”, “GFD” and “GTD” for
	// TrailingStopLossOrder.
	TimeInForce TimeInForce `json:"timeInForce"`

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
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`

	// The trigger price for the TrailingStopLossOrder. The trailing stop
	// value will trail (follow) the market price by the TSL order’s configured
	// “distance” as the market price moves in the winning direction. If the
	// market price moves to a level that is equal to or worse than the trailing
	// stop value, the order will be filled and the Trade will be closed.
	TrailingStopValue decimal.Decimal `json:"trailingStopValue"`

	// ID of the Transaction that filled this Order (only provided when the Order’s state is FILLED)
	FillingTransactionID *TransactionID `json:"fillingTransactionID"`

	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime *time.Time `json:"filledTime"`

	// Trade ID of Trade opened when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID *TradeID `json:"tradeOpenedID"`

	// Trade ID of Trade reduced when the Order was filled (only provided when
	// the Order’s state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID *TradeID `json:"tradeReducedID"`

	// Trade IDs of Trades closed when the Order was filled (only provided when
	// the Order’s state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`

	// ID of the Transaction that cancelled the Order (only provided when the Order’s state is CANCELLED)
	CancellingTransactionID *TransactionID `json:"cancellingTransactionID"`

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime time.Time `json:"cancelledTime"`

	// The ID of the Order that was replaced by this Order (only provided if
	//  this Order was created as part of a cancel/replace).
	ReplacesOrderID *OrderID `json:"replacesOrderID"`

	// The ID of the Order that replaced this Order (only provided if this Order
	// was cancelled as part of a cancel/replace).
	ReplacedByOrderID *OrderID `json:"replacedByOrderID"`
}

func (tslo TrailingStopLossOrder) GetId() string {
	return tslo.Id
}

func (tslo TrailingStopLossOrder) GetCreateTime() time.Time {
	return tslo.CreateTime
}

func (tslo TrailingStopLossOrder) GetState() OrderState {
	return tslo.State
}

func (tslo TrailingStopLossOrder) GetClientExtensions() ClientExtensions {
	return tslo.ClientExtensions
}

func (tslo TrailingStopLossOrder) GetType() OrderType {
	return tslo.Type
}

// OrderID is a string representation of the OANDA-assigned OrderID. OANDA-assigned OrderIDs are positive integers, and
// are derived from the TransactionID of the Transaction that created the Order.
type OrderID string

// OrderType represents the type of the Order
type OrderType string

const (
	Market             = OrderType("MARKET")
	Limit              = OrderType("LIMIT")
	Stop               = OrderType("STOP")
	MarketIfTouched    = OrderType("MARKET_IF_TOUCHED")
	TakeProfit         = OrderType("TAKE_PROFIT")
	StopLoss           = OrderType("STOP_LOSS")
	GuaranteedStopLoss = OrderType("GUARANTEED_STOP_LOSS")
	TrailingStopLoss   = OrderType("TRAILING_STOP_LOSS")
	FixedPrice         = OrderType("FIXED_PRICE")
)

// CancellableOrderType represents the type of the Order
type CancellableOrderType string

const (
	CancellableOrderTypeLimit              = CancellableOrderType("LIMIT")
	CancellableOrderTypeStop               = CancellableOrderType("STOP")
	CancellableOrderTypeMarketIfTouched    = CancellableOrderType("MARKET_IF_TOUCHED")
	CancellableOrderTypeTakeProfit         = CancellableOrderType("TAKE_PROFIT")
	CancellableOrderTypeStopLoss           = CancellableOrderType("STOP_LOSS")
	CancellableOrderTypeGuaranteedStopLoss = CancellableOrderType("GUARANTEED_STOP_LOSS")
	CancellableOrderTypeTrailingStopLoss   = CancellableOrderType("TRAILING_STOP_LOSS")
)

// OrderState represents the current state of the Order
type OrderState string

const (
	// Pending represents the state where the Order is currently pending execution
	Pending = OrderState("PENDING")

	// Filled represents the state where the Order has been filled
	Filled = OrderState("FILLED")

	// Triggered represents the state where the Order has been triggered
	Triggered = OrderState("TRIGGERED")

	// Cancelled represents the state where the Order has been cancelled
	Cancelled = OrderState("CANCELLED")
)

// OrderStateFilter represents the state to filter the requested Orders by
type OrderStateFilter string

const (
	// OrderStateFilterPending represents the state where the Order is currently pending execution
	OrderStateFilterPending = OrderStateFilter("PENDING")

	// OrderStateFilterFilled represents the state where the Order has been filled
	OrderStateFilterFilled = OrderStateFilter("FILLED")

	// OrderStateFilterTriggered represents the state where the Order has been triggered
	OrderStateFilterTriggered = OrderStateFilter("TRIGGERED")

	// OrderStateFilterCancelled represents the state where the Order has been cancelled
	OrderStateFilterCancelled = OrderStateFilter("CANCELLED")

	// OrderStateFilterAll represents any of the possible states listed above
	OrderStateFilterAll = OrderStateFilter("ALL")
)

// OrderIdentifier is used to refer to an Order, and contains both the OrderID and the ClientOrderID
type OrderIdentifier struct {
	// The OANDA-assigned Order ID
	OrderID OrderID `json:"orderID"`

	// The client-provided client Order ID
	ClientOrderID ClientID `json:"clientOrderID"`
}

// OrderSpecifier is a specification of an Order as referred to by clients
type OrderSpecifier string

type TimeInForce string

const (
	// GTC means the Order is "Good unTil Cancelled"
	GTC = TimeInForce("GTC")

	// GTD means the Order is "Good unTil Date" and will be cancelled at the provided time
	GTD = TimeInForce("GTD")

	// GFD means the Order is "Good For Day" and will be cancelled at 5pm New York time
	GFD = TimeInForce("GFD")

	// FOK means the Order must be immediately "Filled Or Killed"
	FOK = TimeInForce("FOK")

	// IOC means the Order must be "Immediately partially filled Or Cancelled"
	IOC = TimeInForce("IOK")
)

// OrderPositionFill is a specification of how Positions in the Account are modified when the Order is filled
type OrderPositionFill string

const (
	// OrderPositionFillOpenOnly means when the Order is filled, only allow Positions to be opened or extended.
	OrderPositionFillOpenOnly = OrderPositionFill("OPEN_ONLY")
	// OrderPositionFillReduceFirst means when the Order is filled, always fully reduce an existing Position before
	// opening a new Position.
	OrderPositionFillReduceFirst = OrderPositionFill("REDUCE_FIRST")
	// OrderPositionFillReduceOnly means when the Order is filled, only reduce an existing Position.
	OrderPositionFillReduceOnly = OrderPositionFill("REDUCE_ONLY")
	// OrderPositionFillDefault means when the Order is filled, use REDUCE_FIRST behaviour for non-client hedging
	// Accounts, and OPEN_ONLY behaviour for client hedging Accounts.
	OrderPositionFillDefault = OrderPositionFill("DEFAULT")
)

// OrderTriggerCondition is a specification of which price component should be used when determining if an Order should
// be triggered and filled. This allows Orders to be triggered based on the bid, ask, mid, default (ask for buy, bid for
// sell) or inverse (ask for sell, bid for buy) price depending on the desired behaviour. Orders are always filled using
// their default price component. This feature is only provided through the REST API. Clients who choose to specify a
// non-default trigger condition will not see it reflected in any of OANDA’s proprietary or partner trading platforms,
// their transaction history or their account statements. OANDA platforms always assume that an Order’s trigger
// condition is set to the default value when indicating the distance from an Order’s trigger price, and will always
// provide the default trigger condition when creating or modifying an Order. A special restriction applies when
// creating a GuaranteedStopLossOrder. In this case the OrderTriggerCondition value must either be “DEFAULT”, or the
// “natural” trigger side “DEFAULT” results in. So for a GuaranteedStopLossOrder for a long trade valid values are
// “DEFAULT” and “BID”, and for short trades “DEFAULT” and “ASK” are valid.
type OrderTriggerCondition string

const (
	// OrderTriggerConditionDefault triggers an Order the "natural" way: compare its price to the ask for long Orders
	// and bid for short Orders.
	OrderTriggerConditionDefault = OrderTriggerCondition("DEFAULT")

	// OrderTriggerConditionInverse triggers an Order the opposite of the “natural” way: compare its price the bid for
	// long Orders and ask for short Orders.
	OrderTriggerConditionInverse = OrderTriggerCondition("INVERSE")

	// OrderTriggerConditionBid triggers an Order by comparing its price to the bid regardless of whether it is long or
	// short.
	OrderTriggerConditionBid = OrderTriggerCondition("BID")

	// OrderTriggerConditionAsk triggers an Order by comparing its price to the ask regardless of whether it is long or
	// short.
	OrderTriggerConditionAsk = OrderTriggerCondition("ASK")

	// OrderTriggerConditionMid triggers an Order by comparing its price to the midpoint regardless of whether it is
	// long or short.
	OrderTriggerConditionMid = OrderTriggerCondition("MID")
)

// DynamicOrderState is the dynamic state of an Order. This is only relevant to TrailingStopLossOrder s, as no other
// type has dynamic state.
type DynamicOrderState struct {
	// The Order's ID.
	ID OrderID `json:"id"`

	// The Order’s calculated trailing stop value.
	TrailingStopValue decimal.Decimal `json:"trailingStopValue"`

	// The distance between the TrailingStopLossOrder’s trailingStopValue and
	// the current Market Price. This represents the distance (in price units)
	// of the Order from a triggering price. If the distance could not be
	// determined, this value will not be set.
	TriggerDistance *decimal.Decimal `json:"triggerDistance"`

	// True if an exact trigger distance could be calculated. If false, it means
	// the provided trigger distance is a best estimate. If the distance could
	// not be determined, this value will not be set.
	IsTriggerDistanceExact *bool `json:"isTriggerDistanceExact"`
}

// UnitsAvailableDetails as a representation of how many units of an Instrument are available to be traded for both long
// and short Orders.
type UnitsAvailableDetails struct {
	// The units available for long Orders.
	Long decimal.Decimal `json:"long"`

	// The units available for short Orders.
	Short decimal.Decimal `json:"short"`
}

// UnitsAvailable is a representation of how many units of an Instrument are available to be traded by an Order
// depending on its positionFill option.
type UnitsAvailable struct {
	// The number of units that are available to be traded using an Order with a
	// positionFill option of “DEFAULT”. For an Account with hedging enabled,
	// this value will be the same as the “OPEN_ONLY” value. For an Account
	// without hedging enabled, this value will be the same as the
	// “REDUCE_FIRST” value.
	Default UnitsAvailableDetails `json:"default"`

	// The number of units that may are available to be traded with an Order with a positionFill option of
	// “REDUCE_FIRST”.
	ReduceFirst UnitsAvailableDetails `json:"reduceFirst"`

	// The number of units that may are available to be traded with an Order with a positionFill option of “REDUCE_ONLY”
	ReduceOnly UnitsAvailableDetails `json:"reduceOnly"`

	// The number of units that may are available to be traded with an Order with a positionFill option of “OPEN_ONLY”.
	OpenOnly UnitsAvailableDetails `json:"openOnly"`
}

// GuaranteedStopLossOrderEntryData carries details required by clients creating a GuaranteedStopLossOrder
type GuaranteedStopLossOrderEntryData struct {
	// The minimum distance allowed between the Trade’s fill price and the
	// configured price for GuaranteedStopLossOrders created for this
	// instrument. Specified in price units.
	MinimumDistance decimal.Decimal `json:"minimumDistance"`

	// The amount that is charged to the account if a GuaranteedStopLossOrder
	// is triggered and filled. The value is in price units and is charged for
	// each unit of the Trade.
	Premium decimal.Decimal `json:"premium"`

	// The GuaranteedStopLossOrderLevelRestriction for this instrument.
	LevelRestriction GuaranteedStopLossOrderLevelRestriction `json:"levelRestriction"`
}
