package oanda_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
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
func (c *Client) GetAccount(accountID string) (*GetAccountResponse, error) {
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
func (c *Client) GetAccountSummary(accountID string) (*GetAccountSummaryResponse, error) {
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
func (c *Client) GetAccountInstruments(accountID, instruments string) (*GetAccountInstrumentsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/instruments?instruments=%s", c.baseUrl, accountID, instruments), nil)
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
func (c *Client) SetAccountConfiguration(accountID string, requestBody SetAccountConfigurationRequest) (*SetAccountConfigurationResponse, error) {
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
func (c *Client) GetAccountChanges(accountID string, sinceTransactionID TransactionID) (*GetAccountChangesResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/accounts/%s/changes?sinceTransactionID=%s", c.baseUrl, accountID, sinceTransactionID), nil)
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
func (c *Client) GetInstrumentOrderBook(instrument string, time *string) (*GetInstrumentOrderBookResponse, error) {
	queryString := ""
	if time != nil {
		queryString = "time=" + *time
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/instruments/%s/orderBook?%s", c.baseUrl, instrument, queryString), nil)
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
func (c *Client) GetInstrumentPositionBook(instrument string, time *string) (*GetInstrumentPositionBookResponse, error) {
	queryString := ""
	if time != nil {
		queryString = "time=" + *time
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v3/instruments/%s/positionBook?%s", c.baseUrl, instrument, queryString), nil)
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
