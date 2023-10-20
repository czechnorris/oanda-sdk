package oanda_sdk

type GetAccountsResponse struct {
	// The list of Accounts the client is authorized to access and their associated properties.
	Accounts []AccountProperties `json:"accounts"`
}

type GetAccountResponse struct {
	// The full details of the requested Account.
	Account Account `json:"account"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type GetAccountSummaryResponse struct {
	// The summary of the requested Account.
	Account AccountSummary `json:"account"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type GetAccountInstrumentsResponse struct {
	// The requested list of instruments.
	Instruments []Instrument `json:"instruments"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type SetAccountConfigurationResponse struct {
	// The transaction that configures the Account.
	ClientConfigureTransaction ClientConfigureTransaction `json:"clientConfigureTransaction"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type SetAccountConfigurationErrorResponse struct {
	ClientConfigureRejectTransaction ClientConfigureRejectTransaction `json:"clientConfigureRejectTransaction"`
	LastTransactionID                TransactionID                    `json:"lastTransactionID"`
	ErrorCode                        *string                          `json:"errorCode"`
	ErrorMessage                     string                           `json:"errorMessage"`
}

func (er SetAccountConfigurationErrorResponse) Error() string {
	return er.ErrorMessage
}

type GetAccountChangesResponse struct {
	// The changes to the Account’s Orders, Trades and Positions since the specified Transaction ID. Only provided if
	// the sinceTransactionID is supplied to the poll request.
	Changes AccountChanges `json:"changes"`

	// The Account’s current price-dependent state.
	State AccountChangesState `json:"state"`

	// The ID of the last Transaction created for the Account. This TransactionID should be used for future poll
	// requests, as the client has already observed all changes up to and including it.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type GetInstrumentCandlesResponse struct {
	// The instrument whose Prices are represented by the candlesticks.
	Instrument string `json:"instrument"`

	// The granularity of the candlesticks provided.
	Granularity CandlestickGranularity `json:"granularity"`

	// The list of candlesticks that satisfy the request.
	Candles []Candlestick `json:"candles"`
}

type GetInstrumentOrderBookResponse struct {
	// The instrument’s order book
	OrderBook OrderBook `json:"orderBook"`
}

type GetInstrumentPositionBookResponse struct {
	// The instrument’s position book
	PositionBook PositionBook `json:"positionBook"`
}

type CreateOrderResponse struct {
	// The Transaction that created the Order specified by the request.
	OrderCreateTransaction Transaction `json:"orderCreateTransaction"`

	// The Transaction that filled the newly created Order. Only provided when the Order was immediately filled.
	OrderFillTransaction *OrderFillTransaction `json:"orderFillTransaction"`

	// The Transaction that cancelled the newly created Order. Only provided when the Order was immediately cancelled.
	OrderCancelTransaction *OrderCancelTransaction `json:"orderCancelTransaction"`

	// The Transaction that reissues the Order. Only provided when the Order is configured to be reissued for its
	// remaining units after a partial fill and the reissue was successful.
	OrderReissueTransaction *Transaction `json:"orderReissueTransaction"`

	// The Transaction that rejects the reissue of the Order. Only provided when the Order is configured to be reissued
	// for its remaining units after a partial fill and the reissue was rejected.
	OrderReissueRejectTransaction *Transaction `json:"orderReissueRejectTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type CreateOrderErrorResponse struct {
	// The Transaction that rejected the creation of the Order as requested
	OrderRejectTransaction Transaction `json:"orderRejectTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode *string `json:"errorCode"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

func (er CreateOrderErrorResponse) Error() string {
	return er.ErrorMessage
}

type GetAccountOrdersResponse struct {
	// The list of Order detail objects
	Orders []Order `json:"orders"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type GetAccountOrderResponse struct {
	// The details of the Order requested
	Order Order `json:"order"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type ReplaceAccountOrderResponse struct {
	// The Transaction that cancelled the Order to be replaced.
	OrderCancelTransaction OrderCancelTransaction `json:"orderCancelTransaction"`

	// The Transaction that created the replacing Order as requested.
	OrderCreateTransaction Transaction `json:"orderCreateTransaction"`

	// The Transaction that filled the replacing Order. This is only provided when the replacing Order was immediately filled.
	OrderFillTransaction *OrderFillTransaction `json:"orderFillTransaction"`

	// The Transaction that reissues the replacing Order. Only provided when the replacing Order was partially filled
	// immediately and is configured to be reissued for its remaining units.
	OrderReissueTransaction *Transaction `json:"orderReissueTransaction"`

	// The Transaction that rejects the reissue of the Order. Only provided when the replacing Order was partially
	// filled immediately and was configured to be reissued, however the reissue was rejected.
	OrderReissueRejectTransaction *Transaction `json:"orderReissueRejectTransaction"`

	// The Transaction that cancelled the replacing Order. Only provided when the replacing Order was immediately cancelled.
	ReplacingOrderCancelTransaction *OrderCancelTransaction `json:"replacingOrderCancelTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type ReplaceAccountOrderErrorResponse struct {
	// The Transaction that rejected the cancellation of the Order to be replaced. Only present if the Account exists.
	OrderCancelRejectTransaction *Transaction `json:"orderCancelRejectTransaction"`

	// The IDs of all Transactions that were created while satisfying the request. Only present if the Account exists.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account. Only present if the Account exists.
	LastTransactionID *TransactionID `json:"lastTransactionID"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode *string `json:"errorCode"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

func (er ReplaceAccountOrderErrorResponse) Error() string {
	return er.ErrorMessage
}

type CancelAccountOrderResponse struct {
	// The Transaction that cancelled the Order
	OrderCancelTransaction OrderCancelTransaction `json:"orderCancelTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type UpdateClientExtensionsResponse struct {
	// The Transaction that modified the Client Extensions for the Order
	OrderClientExtensionsModifyTransaction OrderClientExtensionsModifyTransaction `json:"orderClientExtensionsModifyTransaction"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
}

type UpdateClientExtensionsErrorResponse struct {
	// The Transaction that rejected the modification of the Client Extensions for the Order. Only present if the Account exists.
	OrderClientExtensionsModifyRejectTransaction *OrderClientExtensionsModifyRejectTransaction `json:"orderClientExtensionsModifyRejectTransaction"`

	// The ID of the most recent Transaction created for the Account. Only present if the Account exists.
	LastTransactionID *TransactionID `json:"lastTransactionID"`

	// The IDs of all Transactions that were created while satisfying the request. Only present if the Account exists.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode *string `json:"errorCode"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

func (er UpdateClientExtensionsErrorResponse) Error() string {
	return er.ErrorMessage
}

type GetAccountTradesResponse struct {
	// The list of Trade detail objects
	Trades []Trade `json:"trades"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type GetAccountTradeResponse struct {
	// The details of the requested trade
	Trade Trade `json:"trade"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type CloseAccountTradeResponse struct {
	// The MarketOrder Transaction created to close the Trade.
	OrderCreateTransaction MarketOrderTransaction `json:"orderCreateTransaction"`

	// The OrderFill Transaction that fills the Trade-closing MarketOrder and closes the Trade.
	OrderFillTransaction OrderFillTransaction `json:"orderFillTransaction"`

	// The OrderCancel Transaction that immediately cancelled the Trade-closing MarketOrder.
	OrderCancelTransaction OrderCancelTransaction `json:"orderCancelTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type CloseAccountTradeErrorResponse struct {
	// The MarketOrderReject Transaction that rejects the creation of the Trade-closing MarketOrder. Only present if the
	// Account exists.
	OrderRejectTransaction MarketOrderRejectTransaction `json:"orderRejectTransaction"`

	// The ID of the most recent Transaction created for the Account. Only present if the Account exists.
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The IDs of all Transactions that were created while satisfying the request. Only present if the Account exists.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode *string `json:"errorCode"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

func (er CloseAccountTradeErrorResponse) Error() string {
	return er.ErrorMessage
}

type UpdateAccountTradeResponse struct {
	// The Transaction that updates the Trade’s Client Extensions.
	TradeClientExtensionsModifyTransaction TradeClientExtensionsModifyTransaction `json:"tradeClientExtensionsModifyTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type UpdateAccountTradeErrorResponse struct {
	// The Transaction that rejects the modification of the Trade’s Client Extensions.
	TradeClientExtensionsModifyRejectTransaction TradeClientExtensionsModifyRejectTransaction `json:"tradeClientExtensionsModifyRejectTransaction"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode *string `json:"errorCode"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

func (er UpdateAccountTradeErrorResponse) Error() string {
	return er.ErrorMessage
}

type UpdateAccountTradeOrdersResponse struct {
	// The Transaction created that cancels the Trade’s existing TakeProfitOrder.
	TakeProfitOrderCancelTransaction *OrderCancelTransaction `json:"takeProfitOrderCancelTransaction"`

	// The Transaction created that creates a new TakeProfitOrder for the Trade.
	TakeProfitOrderTransaction *TakeProfitOrderTransaction `json:"takeProfitOrderTransaction"`

	// The Transaction created that immediately fills the Trade’s new TakeProfitOrder. Only provided if the new TakeProfitOrder was immediately filled.
	TakeProfitOrderFillTransaction *OrderFillTransaction `json:"takeProfitOrderFillTransaction"`

	// The Transaction created that immediately cancels the Trade’s new TakeProfitOrder. Only provided if the new TakeProfitOrder was immediately cancelled.
	TakeProfitOrderCreatedCancelTransaction *OrderCancelTransaction `json:"takeProfitOrderCreatedCancelTransaction"`

	// The Transaction created that cancels the Trade’s existing StopLossOrder.
	StopLossOrderCancelTransaction *OrderCancelTransaction `json:"stopLossOrderCancelTransaction"`

	// The Transaction created that creates a new StopLossOrder for the Trade.
	StopLossOrderTransaction *StopLossOrderTransaction `json:"stopLossOrderTransaction"`

	// The Transaction created that immediately fills the Trade’s new StopOrder. Only provided if the new StopLossOrder was immediately filled.
	StopLossOrderFillTransaction *OrderFillTransaction `json:"stopLossOrderFillTransaction"`

	// The Transaction created that immediately cancels the Trade’s new StopLossOrder. Only provided if the new StopLossOrder was immediately cancelled.
	StopLossOrderCreatedCancelTransaction *OrderCancelTransaction `json:"stopLossOrderCreatedCancelTransaction"`

	// The Transaction created that cancels the Trade’s existing TrailingStopLossOrder.
	TrailingStopLossOrderCancelTransaction *OrderCancelTransaction `json:"trailingStopLossOrderCancelTransaction"`

	// The Transaction created that creates a new TrailingStopLossOrder for the Trade.
	TrailingStopLossOrderTransaction *TrailingStopLossOrderTransaction `json:"trailingStopLossOrderTransaction"`

	// The Transaction created that cancels the Trade’s existing Guaranteed StopLossOrder.
	GuaranteedStopLossOrderCancelTransaction *OrderCancelTransaction `json:"guaranteedStopLossOrderCancelTransaction"`

	// The Transaction created that creates a new GuaranteedStopLossOrder for the Trade.
	GuaranteedStopLossOrderTransaction *GuaranteedStopLossOrderTransaction `json:"guaranteedStopLossOrderTransaction"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type UpdateAccountTradeOrdersErrorResponse struct {
	// An OrderCancelRejectTransaction represents the rejection of the cancellation of an Order in the client’s Account.
	TakeProfitOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"takeProfitOrderCancelRejectTransaction"`

	// A TakeProfitOrderRejectTransaction represents the rejection of the creation of a TakeProfitOrder.
	TakeProfitOrderRejectTransaction *TakeProfitOrderRejectTransaction `json:"takeProfitOrderRejectTransaction"`

	// An OrderCancelRejectTransaction represents the rejection of the cancellation of an Order in the client’s Account.
	StopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"stopLossOrderCancelRejectTransaction"`

	// A StopLossOrderRejectTransaction represents the rejection of the creation of a StopLossOrder.
	StopLossOrderRejectTransaction *StopLossOrderRejectTransaction `json:"stopLossOrderRejectTransaction"`

	// An OrderCancelRejectTransaction represents the rejection of the cancellation of an Order in the client’s Account.
	TrailingStopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"trailingStopLossOrderCancelRejectTransaction"`

	// A TrailingStopLossOrderRejectTransaction represents the rejection of the creation of a TrailingStopLossOrder.
	TrailingStopLossOrderRejectTransaction *TrailingStopLossOrderRejectTransaction `json:"trailingStopLossOrderRejectTransaction"`

	// An OrderCancelRejectTransaction represents the rejection of the cancellation of an Order in the client’s Account.
	GuaranteedStopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"guaranteedStopLossOrderCancelRejectTransaction"`

	// A GuaranteedStopLossOrderRejectTransaction represents the rejection of the creation of a GuaranteedStopLossOrder.
	GuaranteedStopLossOrderRejectTransaction *GuaranteedStopLossOrderRejectTransaction `json:"guaranteedStopLossOrderRejectTransaction"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode *string `json:"errorCode"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

func (er UpdateAccountTradeOrdersErrorResponse) Error() string {
	return er.ErrorMessage
}
