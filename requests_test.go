package oanda_sdk

import (
	"github.com/google/go-querystring/query"
	"testing"
	"time"
)

func TestGetInstrumentCandlesRequestIntoQueryString(t *testing.T) {
	price := PricingComponent("M")
	granularity := M15
	count := 1000
	from := time.Date(2022, 12, 9, 12, 0, 0, 0, time.UTC)
	to := time.Date(2022, 12, 10, 12, 0, 0, 0, time.UTC)
	smooth := false
	includeFirst := false
	dailyAlignment := 5
	alignmentTimezone := "Europe/Prague"
	weeklyAlignment := Friday
	request := GetInstrumentCandlesRequest{
		Price:             &price,
		Granularity:       &granularity,
		Count:             &count,
		From:              &from,
		To:                &to,
		Smooth:            &smooth,
		IncludeFirst:      &includeFirst,
		DailyAlignment:    &dailyAlignment,
		AlignmentTimezone: &alignmentTimezone,
		WeeklyAlignment:   &weeklyAlignment,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "alignmentTimezone=Europe%2FPrague&count=1000&dailyAlignment=5&from=2022-12-09T12%3A00%3A00Z&granularity=M15&includeFirst=false&price=M&smooth=false&to=2022-12-10T12%3A00%3A00Z&weeklyAlignment=Friday" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetInstrumentCandlesRequestIntoQueryStringSimple(t *testing.T) {
	granularity := H1
	request := GetInstrumentCandlesRequest{
		Granularity: &granularity,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "granularity=H1" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountOrdersRequestIntoQueryString(t *testing.T) {
	ids := []OrderID{"123", "456", "789"}
	state := OrderStateFilterFilled
	instrument := "EUR_USD"
	count := 100
	beforeID := OrderID("000")
	request := GetAccountOrdersRequest{
		IDs:        ids,
		State:      &state,
		Instrument: &instrument,
		Count:      &count,
		BeforeID:   &beforeID,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "beforeID=000&count=100&ids=123%2C456%2C789&instrument=EUR_USD&state=FILLED" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountOrdersRequestIntoQueryStringSimple(t *testing.T) {
	request := GetAccountOrdersRequest{}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTradesRequestIntoQueryString(t *testing.T) {
	ids := []TradeID{"123", "456", "789"}
	state := TradeStateFilterOpen
	instrument := "EUR_USD"
	count := 100
	beforeID := TradeID("000")
	request := GetAccountTradesRequest{
		IDs:        ids,
		State:      &state,
		Instrument: &instrument,
		Count:      &count,
		BeforeID:   &beforeID,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "beforeID=000&count=100&ids=123%2C456%2C789&instrument=EUR_USD&state=OPEN" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTradesRequestIntoQueryStringSimple(t *testing.T) {
	request := GetAccountTradesRequest{}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTransactionsRequestIntoQueryString(t *testing.T) {
	from := time.Date(2022, 12, 9, 12, 0, 0, 0, time.UTC)
	to := time.Date(2022, 12, 10, 12, 0, 0, 0, time.UTC)
	pageSize := 200
	typ := []TransactionFilter{TransactionFilterTrailingStopLossOrder, TransactionFilterTrailingStopLossOrderReject}
	request := GetAccountTransactionsRequest{
		From:     &from,
		To:       &to,
		PageSize: &pageSize,
		Type:     typ,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "from=2022-12-09T12%3A00%3A00Z&pageSize=200&to=2022-12-10T12%3A00%3A00Z&type=TRAILING_STOP_LOSS_ORDER%2CTRAILING_STOP_LOSS_ORDER_REJECT" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTransactionsRequestIntoQueryStringSimple(t *testing.T) {
	request := GetAccountTransactionsRequest{}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTransactionsByIdRangeRequest(t *testing.T) {
	from := TransactionID("123")
	to := TransactionID("789")
	typ := []TransactionFilter{TransactionFilterTrailingStopLossOrder, TransactionFilterTrailingStopLossOrderReject}
	request := GetAccountTransactionsByIdRangeRequest{
		From: from,
		To:   to,
		Type: typ,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "from=123&to=789&type=TRAILING_STOP_LOSS_ORDER%2CTRAILING_STOP_LOSS_ORDER_REJECT" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTransactionsByIdRangeRequestSimple(t *testing.T) {
	from := TransactionID("123")
	to := TransactionID("789")
	request := GetAccountTransactionsByIdRangeRequest{
		From: from,
		To:   to,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "from=123&to=789" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTransactionsSinceIdRequestIntoQueryString(t *testing.T) {
	id := TransactionID("123")
	typ := []TransactionFilter{TransactionFilterTrailingStopLossOrder, TransactionFilterTrailingStopLossOrderReject}
	request := GetAccountTransactionsSinceIdRequest{
		Id:   id,
		Type: typ,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "id=123&type=TRAILING_STOP_LOSS_ORDER%2CTRAILING_STOP_LOSS_ORDER_REJECT" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountTransactionsSinceIdRequestIntoQueryStringSimple(t *testing.T) {
	id := TransactionID("123")
	request := GetAccountTransactionsSinceIdRequest{
		Id: id,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "id=123" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountLatestCandlesRequestIntoQueryString(t *testing.T) {
	candleSpecifications := []CandleSpecification{"EUR_USD:S10:BM", "EUR_USD:M5:BMA"}
	units := 10.15
	smooth := true
	dailyAlignment := 5
	alignmentTimezone := "Europe/Prague"
	weeklyAlignment := Monday
	request := GetAccountLatestCandlesRequest{
		CandleSpecifications: candleSpecifications,
		Units:                &units,
		Smooth:               &smooth,
		DailyAlignment:       &dailyAlignment,
		AlignmentTimezone:    &alignmentTimezone,
		WeeklyAlignment:      &weeklyAlignment,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "alignmentTimezone=Europe%2FPrague&candleSpecifications=EUR_USD%3AS10%3ABM%2CEUR_USD%3AM5%3ABMA&dailyAlignment=5&decimal=10.15&smooth=true&weeklyAlignment=Monday" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountLatestCandlesRequestIntoQueryStringSimple(t *testing.T) {
	candleSpecifications := []CandleSpecification{"EUR_USD:S10:BM", "EUR_USD:M5:BMA"}
	request := GetAccountLatestCandlesRequest{
		CandleSpecifications: candleSpecifications,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "candleSpecifications=EUR_USD%3AS10%3ABM%2CEUR_USD%3AM5%3ABMA" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountPricingRequestIntoQueryString(t *testing.T) {
	instruments := []string{"EUR_USD", "CAD_USD"}
	since := time.Date(2022, 12, 9, 12, 0, 0, 0, time.UTC)
	includeHomeConversion := true
	request := GetAccountPricingRequest{
		Instruments:           instruments,
		Since:                 &since,
		IncludeHomeConversion: &includeHomeConversion,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "includeHomeConversion=true&instruments=EUR_USD%2CCAD_USD&since=2022-12-09T12%3A00%3A00Z" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountPricingRequestIntoQueryStringSimple(t *testing.T) {
	instruments := []string{"EUR_USD", "CAD_USD"}
	request := GetAccountPricingRequest{
		Instruments: instruments,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "instruments=EUR_USD%2CCAD_USD" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountInstrumentCandlesRequestIntoQueryString(t *testing.T) {
	price := PricingComponent("M")
	granularity := M15
	count := 1000
	from := time.Date(2022, 12, 9, 12, 0, 0, 0, time.UTC)
	to := time.Date(2022, 12, 10, 12, 0, 0, 0, time.UTC)
	smooth := false
	includeFirst := false
	dailyAlignment := 5
	alignmentTimezone := "Europe/Prague"
	weeklyAlignment := Friday
	units := 5.5
	request := GetAccountInstrumentCandlesRequest{
		Price:             &price,
		Granularity:       &granularity,
		Count:             &count,
		From:              &from,
		To:                &to,
		Smooth:            &smooth,
		IncludeFirst:      &includeFirst,
		DailyAlignment:    &dailyAlignment,
		AlignmentTimezone: &alignmentTimezone,
		WeeklyAlignment:   &weeklyAlignment,
		Units:             &units,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "alignmentTimezone=Europe%2FPrague&count=1000&dailyAlignment=5&from=2022-12-09T12%3A00%3A00Z&granularity=M15&includeFirst=false&price=M&smooth=false&to=2022-12-10T12%3A00%3A00Z&units=5.5&weeklyAlignment=Friday" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountInstrumentCandlesRequestIntoQueryStringSimple(t *testing.T) {
	request := GetAccountInstrumentCandlesRequest{}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountPricingStreamRequestIntoQueryString(t *testing.T) {
	instruments := []string{"EUR_USD", "CAD_USD"}
	snapshot := true
	includeHomeConversion := true
	request := GetAccountPricingStreamRequest{
		Instruments:           instruments,
		Snapshot:              &snapshot,
		IncludeHomeConversion: &includeHomeConversion,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "includeHomeConversion=true&instruments=EUR_USD%2CCAD_USD&snapshot=true" {
		t.Error("Got ", urlQuery.Encode())
	}
}

func TestGetAccountPricingStreamRequestIntoQueryStringSimple(t *testing.T) {
	instruments := []string{"EUR_USD", "CAD_USD"}
	request := GetAccountPricingStreamRequest{
		Instruments: instruments,
	}
	urlQuery, err := query.Values(request)
	if err != nil {
		t.Error(err)
	}
	if urlQuery.Encode() != "instruments=EUR_USD%2CCAD_USD" {
		t.Error("Got ", urlQuery.Encode())
	}
}
