package oanda_sdk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"time"
)

type Client struct {
	baseUrl     string
	accessToken string
	conn        *http.Client
}

func NewClient(baseUrl, accessToken string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		baseUrl:     baseUrl,
		accessToken: accessToken,
		conn:        client,
	}
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Oanda SDK for GO")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Datetime-Format", "RFC3339")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
}

// GetAccounts returns a list of all Accounts authorized for the provided token
func (c *Client) GetAccounts() (*GetAccountsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+"/v3/accounts", nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accounts GetAccountsResponse
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		return nil, err
	}
	return &accounts, nil
}

// GetAccount gets the full details for a single Account that a client has access to. Full pending Order, open Trade
// and open Position representations are provided.
func (c *Client) GetAccount(accountID AccountID) (*GetAccountResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountResponse GetAccountResponse
	err = json.NewDecoder(resp.Body).Decode(&accountResponse)
	if err != nil {
		return nil, err
	}
	return &accountResponse, nil
}

// GetAccountSummary gets a summary for a single Account that a client has access to.
func (c *Client) GetAccountSummary(accountID AccountID) (*GetAccountSummaryResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/summary", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountSummaryResponse GetAccountSummaryResponse
	err = json.NewDecoder(resp.Body).Decode(&accountSummaryResponse)
	if err != nil {
		return nil, err
	}
	return &accountSummaryResponse, nil
}

// GetAccountInstruments gets the list of tradeable instruments for the given Account. The list of tradeable instruments
// is dependent on the regulatory division that the Account is located in, thus should be the same for all Accounts
// owned by a single user.
func (c *Client) GetAccountInstruments(accountID AccountID, instruments []string) (*GetAccountInstrumentsResponse, error) {
	urlQuery, err := query.Values(struct {
		Instruments []string `url:"instruments,comma,omitempty"`
	}{
		Instruments: instruments,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/instruments?%s", c.baseUrl, accountID, urlQuery.Encode()), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountInstrumentsResponse GetAccountInstrumentsResponse
	err = json.NewDecoder(resp.Body).Decode(&accountInstrumentsResponse)
	if err != nil {
		return nil, err
	}
	return &accountInstrumentsResponse, nil
}

// SetAccountConfiguration sets the client-configurable portions of the Account.
func (c *Client) SetAccountConfiguration(accountID AccountID, requestBody SetAccountConfigurationRequest) (*SetAccountConfigurationResponse, error) {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/v3/accounts/%s/configuration", c.baseUrl, accountID), &buffer)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var setAccountConfigurationResponse SetAccountConfigurationResponse
		err = json.NewDecoder(resp.Body).Decode(&setAccountConfigurationResponse)
		if err != nil {
			return nil, err
		}
		return &setAccountConfigurationResponse, nil
	case http.StatusBadRequest, http.StatusForbidden:
		var setAccountConfigurationErrorResponse SetAccountConfigurationErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&setAccountConfigurationErrorResponse)
		if err != nil {
			return nil, fmt.Errorf("received an HTTP %d response, err: %v", resp.StatusCode, err)
		}
		return nil, setAccountConfigurationErrorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// GetAccountChanges is used to poll an Account for its current state and changes since a specified TransactionID
func (c *Client) GetAccountChanges(accountID AccountID, sinceTransactionID TransactionID) (*GetAccountChangesResponse, error) {
	urlQuery, err := query.Values(struct {
		SinceTransactionID TransactionID `url:"sinceTransactionID"`
	}{
		SinceTransactionID: sinceTransactionID,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/changes?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountChangesResponse GetAccountChangesResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountChangesResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountChangesResponse, nil
}

// GetInstrumentCandles fetches candlestick data for an instrument
func (c *Client) GetInstrumentCandles(instrument string, request GetInstrumentCandlesRequest) (*GetInstrumentCandlesResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/instruments/%s/candles?%s", c.baseUrl, instrument, urlQuery.Encode()), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var instrumentCandlesResponse GetInstrumentCandlesResponse
	err = json.NewDecoder(resp.Body).Decode(&instrumentCandlesResponse)
	if err != nil {
		return nil, err
	}
	return &instrumentCandlesResponse, nil
}

// GetInstrumentOrderBook fetches an order book for an instrument
func (c *Client) GetInstrumentOrderBook(instrument string, snapshotTime *time.Time) (*GetInstrumentOrderBookResponse, error) {
	urlQuery, err := query.Values(struct {
		Time *time.Time `url:"time,omitempty"`
	}{
		Time: snapshotTime,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/instruments/%s/orderBook?%s", c.baseUrl, instrument, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var instrumentOrderBookResponse GetInstrumentOrderBookResponse
	err = json.NewDecoder(resp.Body).Decode(&instrumentOrderBookResponse)
	if err != nil {
		return nil, err
	}
	return &instrumentOrderBookResponse, nil
}

// GetInstrumentPositionBook fetches a position book for an instrument
func (c *Client) GetInstrumentPositionBook(instrument string, snapshotTime *time.Time) (*GetInstrumentPositionBookResponse, error) {
	urlQuery, err := query.Values(struct {
		Time *time.Time `url:"time,omitempty"`
	}{
		Time: snapshotTime,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/instruments/%s/positionBook?%s", c.baseUrl, instrument, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var instrumentPositionBookResponse GetInstrumentPositionBookResponse
	err = json.NewDecoder(resp.Body).Decode(&instrumentPositionBookResponse)
	if err != nil {
		return nil, err
	}
	return &instrumentPositionBookResponse, nil
}

// CreateOrder creates an Order for an Account
func (c *Client) CreateOrder(accountID AccountID, orderRequest OrderRequest) (*CreateOrderResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(orderRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v3/accounts/%s/orders", c.baseUrl, accountID), &buf)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusCreated:
		var createdOrderResponse CreateOrderResponse
		err = json.NewDecoder(resp.Body).Decode(&createdOrderResponse)
		if err != nil {
			return nil, err
		}
		return &createdOrderResponse, nil
	case http.StatusBadRequest, http.StatusNotFound:
		var errorResponse CreateOrderErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// GetAccountOrders gets a list of Orders for an Account
func (c *Client) GetAccountOrders(accountID AccountID, request GetAccountOrdersRequest) (*GetAccountOrdersResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/orders?%s", c.baseUrl, accountID, urlQuery.Encode()), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountOrdersResponse GetAccountOrdersResponse
	err = json.NewDecoder(resp.Body).Decode(&accountOrdersResponse)
	if err != nil {
		return nil, err
	}
	return &accountOrdersResponse, err
}

// GetAccountPendingOrders lists all pending Orders in an Account
func (c *Client) GetAccountPendingOrders(accountID AccountID) (*GetAccountOrdersResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/pendingOrders", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountOrdersResponse GetAccountOrdersResponse
	err = json.NewDecoder(resp.Body).Decode(&accountOrdersResponse)
	if err != nil {
		return nil, err
	}
	return &accountOrdersResponse, nil
}

// GetAccountOrder gets details for a single Order in an Account
func (c *Client) GetAccountOrder(accountID AccountID, orderSpecifier OrderSpecifier) (*GetAccountOrderResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/orders/%s", c.baseUrl, accountID, orderSpecifier), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountOrderResponse GetAccountOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&accountOrderResponse)
	if err != nil {
		return nil, err
	}
	return &accountOrderResponse, nil
}

// ReplaceAccountOrder replaces an Order in an Account by simultaneously cancelling it and creating a replacement Order
func (c *Client) ReplaceAccountOrder(accountID AccountID, orderSpecifier OrderSpecifier, orderRequest OrderRequest) (*ReplaceAccountOrderResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(orderRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/orders/%s", c.baseUrl, accountID, orderSpecifier), &buf)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusCreated:
		var replaceAccountOrderResponse ReplaceAccountOrderResponse
		err = json.NewDecoder(resp.Body).Decode(&replaceAccountOrderResponse)
		if err != nil {
			return nil, err
		}
		return &replaceAccountOrderResponse, nil
	case http.StatusBadRequest:
		var errorResponse CreateOrderErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	case http.StatusNotFound:
		var errorResponse ReplaceAccountOrderErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// CancelAccountOrder cancels a pending Order in an Account
func (c *Client) CancelAccountOrder(accountID AccountID, orderSpecifier OrderSpecifier) (*CancelAccountOrderResponse, error) {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/orders/%s/cancel", c.baseUrl, accountID, orderSpecifier), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var cancelAccountOrderResponse CancelAccountOrderResponse
		err = json.NewDecoder(resp.Body).Decode(&cancelAccountOrderResponse)
		if err != nil {
			return nil, err
		}
		return &cancelAccountOrderResponse, nil
	case http.StatusNotFound:
		var errorResponse ReplaceAccountOrderErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

func (c *Client) UpdateAccountOrderClientExtensions(accountID AccountID, orderSpecifier OrderSpecifier, updateClientExtensionsRequest UpdateClientExtensionsRequest) (*UpdateClientExtensionsResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updateClientExtensionsRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/orders/%s/clientExtensions", c.baseUrl, accountID, orderSpecifier), &buf)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var okResponse UpdateClientExtensionsResponse
		err = json.NewDecoder(resp.Body).Decode(&okResponse)
		if err != nil {
			return nil, err
		}
		return &okResponse, nil
	case http.StatusBadRequest:
		var errorResponse UpdateClientExtensionsErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// GetAccountTrades gets a list of Trades for an Account
func (c *Client) GetAccountTrades(accountID AccountID, request GetAccountTradesRequest) (*GetAccountTradesResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/trades?%s", c.baseUrl, accountID, urlQuery.Encode()), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountTradesResponse GetAccountTradesResponse
	err = json.NewDecoder(resp.Body).Decode(&accountTradesResponse)
	if err != nil {
		return nil, err
	}
	return &accountTradesResponse, nil
}

// GetAccountOpenTrades gets the list of open Trades for an Account
func (c *Client) GetAccountOpenTrades(accountID AccountID) (*GetAccountTradesResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/openTrades", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountOpenTradesResponse GetAccountTradesResponse
	err = json.NewDecoder(resp.Body).Decode(&accountOpenTradesResponse)
	if err != nil {
		return nil, err
	}
	return &accountOpenTradesResponse, nil
}

// GetAccountTrade gets the details of a specific Trade in an Account
func (c *Client) GetAccountTrade(accountID AccountID, tradeSpecifier TradeSpecifier) (*GetAccountTradeResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/trades/%s", c.baseUrl, accountID, tradeSpecifier), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var accountTradeResponse GetAccountTradeResponse
	err = json.NewDecoder(resp.Body).Decode(&accountTradeResponse)
	if err != nil {
		return nil, err
	}
	return &accountTradeResponse, nil
}

// CloseAccountTrade closes (partially or fully) a specific open Trade in an Account
func (c *Client) CloseAccountTrade(accountID AccountID, tradeSpecifier TradeSpecifier) (*CloseAccountTradeResponse, error) {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/trades/%s/close", c.baseUrl, accountID, tradeSpecifier), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var closeAccountTradeResponse CloseAccountTradeResponse
		err = json.NewDecoder(resp.Body).Decode(&closeAccountTradeResponse)
		if err != nil {
			return nil, err
		}
		return &closeAccountTradeResponse, nil
	case http.StatusBadRequest:
		fallthrough
	case http.StatusNotFound:
		var errorResponse CloseAccountTradeErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// UpdateAccountTradeClientExtensions updates the ClientExtensions for a Trade. Do not add, update or delete the
// ClientExtensions if your account is associated with MT4.
func (c *Client) UpdateAccountTradeClientExtensions(accountID AccountID, tradeSpecifier TradeSpecifier) (*UpdateAccountTradeResponse, error) {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/trades/%s/clientExtensions", c.baseUrl, accountID, tradeSpecifier), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var updateAccountTradeResponse UpdateAccountTradeResponse
		err = json.NewDecoder(resp.Body).Decode(&updateAccountTradeResponse)
		if err != nil {
			return nil, err
		}
		return &updateAccountTradeResponse, nil
	case http.StatusBadRequest:
		fallthrough
	case http.StatusNotFound:
		var errorResponse UpdateAccountTradeErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// UpdateAccountTradeOrders creates, replaces and cancels a Trade's dependent Orders (TakeProfit, StopLoss and
// TrailingStopLoss) through the Trade itself
func (c *Client) UpdateAccountTradeOrders(accountID AccountID, tradeSpecifier TradeSpecifier, updateAccountTradeOrdersRequest UpdateAccountTradeOrdersRequest) (*UpdateAccountTradeOrdersResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(updateAccountTradeOrdersRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/trades/%s/orders", c.baseUrl, accountID, tradeSpecifier), &buf)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var updateAccountTradeOrdersResponse UpdateAccountTradeOrdersResponse
		err = json.NewDecoder(resp.Body).Decode(&updateAccountTradeOrdersResponse)
		if err != nil {
			return nil, err
		}
		return &updateAccountTradeOrdersResponse, nil
	case http.StatusBadRequest:
		var errorResponse UpdateAccountTradeOrdersErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// GetAccountPositions lists all Positions for an Account. The Positions returned are for every instrument that has had
// a position during the lifetime of the Account.
func (c *Client) GetAccountPositions(accountID AccountID) (*GetAccountPositionsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/positions", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountPositionsResponse GetAccountPositionsResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountPositionsResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountPositionsResponse, nil
}

// GetAccountOpenPositions lists all open Positions for an Account. An open Position is a Position in an Account that
// currently has a Trade opened for it.
func (c *Client) GetAccountOpenPositions(accountID AccountID) (*GetAccountPositionsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/openPositions", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountOpenPositionsResponse GetAccountPositionsResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountOpenPositionsResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountOpenPositionsResponse, nil
}

// GetAccountInstrumentPosition gets the details of a single Instrument's Position in an Account. The Position may be
// open or not.
func (c *Client) GetAccountInstrumentPosition(accountID AccountID, instrument string) (*GetAccountInstrumentPositionResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/positions/%s", c.baseUrl, accountID, instrument), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d repsonse", resp.StatusCode)
	}
	var getAccountInstrumentPositionResponse GetAccountInstrumentPositionResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountInstrumentPositionResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountInstrumentPositionResponse, nil
}

// CloseAccountInstrumentPosition closeouts the opan Position for a specific Instrument in an Account.
func (c *Client) CloseAccountInstrumentPosition(accountID AccountID, instrument string, request CloseAccountInstrumentPositionRequest) (*CloseAccountInstrumentPositionResponse, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v3/accounts/%s/positions/%s/close", c.baseUrl, accountID, instrument), &buf)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		var closeAccountInstrumentPositionResponse CloseAccountInstrumentPositionResponse
		err = json.NewDecoder(resp.Body).Decode(&closeAccountInstrumentPositionResponse)
		if err != nil {
			return nil, err
		}
		return &closeAccountInstrumentPositionResponse, nil
	case http.StatusBadRequest:
		fallthrough
	case http.StatusNotFound:
		var errorResponse CloseAccountInstrumentPositionErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}
	return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
}

// GetAccountTransactions gets a list of Transaction pages that satisfy a time-based Transaction query.
func (c *Client) GetAccountTransactions(accountID AccountID, request GetAccountTransactionsRequest) (*GetAccountTransactionsResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/transactions?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountTransactionsResponse GetAccountTransactionsResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountTransactionsResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountTransactionsResponse, nil
}

// GetAccountTransaction gets the details of a single Account Transaction
func (c *Client) GetAccountTransaction(accountID AccountID, transactionID TransactionID) (*GetAccountTransactionResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/transactions/%s", c.baseUrl, accountID, transactionID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountTransactionResponse GetAccountTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountTransactionResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountTransactionResponse, nil
}

// GetAccountTransactionsByIdRange gets a range of Transactions for an Account based on the TransactionIDs.
func (c *Client) GetAccountTransactionsByIdRange(accountID AccountID, request GetAccountTransactionsByIdRangeRequest) (*GetAccountTransactionsRangeResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/transactions/idrange?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d reponse", resp.StatusCode)
	}
	var getAccountTransactionsResponse GetAccountTransactionsRangeResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountTransactionsResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountTransactionsResponse, nil
}

// GetAccountTransactionsSinceId gets a range of Transactions for an Account starting at a provided TransactionID
func (c *Client) GetAccountTransactionsSinceId(accountID AccountID, request GetAccountTransactionsSinceIdRequest) (*GetAccountTransactionsRangeResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/transactions/idrange?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d reponse", resp.StatusCode)
	}
	var getAccountTransactionsResponse GetAccountTransactionsRangeResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountTransactionsResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountTransactionsResponse, nil
}

func (c *Client) GetAccountTransactionsStream(accountID AccountID) (<-chan Transaction, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/transactions/stream", c.baseUrl, accountID), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	transactions := make(chan Transaction)
	go func() {
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			type message struct {
				Type string `json:"type"`
			}
			var msg message
			err = json.NewDecoder(bytes.NewReader(line)).Decode(&msg)
			switch msg.Type {
			case "HEARTBEAT":
				continue
			default:
				var transaction Transaction
				err = json.NewDecoder(bytes.NewReader(line)).Decode(&transaction)
				if err != nil {
					continue
				}
				transactions <- transaction
			}
		}
	}()

	return transactions, nil
}

// GetAccountLatestCandles get dancing bears and most recently completed candles within an Account for specified
// combinations of instrument, granularity and price component.
func (c *Client) GetAccountLatestCandles(accountID AccountID, request GetAccountLatestCandlesRequest) (*GetAccountLatestCandlesResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/candles/latest?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountLatestCandlesResponse GetAccountLatestCandlesResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountLatestCandlesResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountLatestCandlesResponse, nil
}

// GetAccountPricing gets pricing information for a specified list of Instruments within an Account
func (c *Client) GetAccountPricing(accountID AccountID, request GetAccountPricingRequest) (*GetAccountPricingResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/pricing?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountPricingResponse GetAccountPricingResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountPricingResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountPricingResponse, nil
}

// GetAccountInstrumentCandles fetches candlestick data for an Instrument
func (c *Client) GetAccountInstrumentCandles(accountID AccountID, instrument string, request GetAccountInstrumentCandlesRequest) (*GetAccountInstrumentCandlesResponse, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/instruments/%s/candles?%s", c.baseUrl, accountID, instrument, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	var getAccountInstrumentCandlesResponse GetAccountInstrumentCandlesResponse
	err = json.NewDecoder(resp.Body).Decode(&getAccountInstrumentCandlesResponse)
	if err != nil {
		return nil, err
	}
	return &getAccountInstrumentCandlesResponse, nil
}

func (c *Client) GetAccountPricingStream(accountID AccountID, request GetAccountPricingStreamRequest) (<-chan ClientPrice, error) {
	urlQuery, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/pricing/stream?%s", c.baseUrl, accountID, urlQuery), nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	resp, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received an HTTP %d response", resp.StatusCode)
	}
	prices := make(chan ClientPrice)
	go func() {
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			type message struct {
				Type string `json:"type"`
			}
			var msg message
			err = json.NewDecoder(bytes.NewReader(line)).Decode(&msg)
			switch msg.Type {
			case "HEARTBEAT":
				continue
			case "PRICE":
				var price ClientPrice
				err = json.NewDecoder(bytes.NewReader(line)).Decode(&price)
				if err != nil {
					continue
				}
				prices <- price
			}
		}
	}()

	return prices, nil
}
