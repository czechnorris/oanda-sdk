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
