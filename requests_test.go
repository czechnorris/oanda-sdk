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
