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
