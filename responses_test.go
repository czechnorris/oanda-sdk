package oanda_sdk

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetAccountResponseUnmarshalling(t *testing.T) {
	file, err := os.Open("test/accountResponse.json")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	var accountResponse GetAccountResponse
	err = json.NewDecoder(file).Decode(&accountResponse)
	if err != nil {
		t.Error(err)
	}
	if accountResponse.Account.ResettablePLTime.Time != nil {
		t.Errorf("ResettablePLTime should be nil")
	}
}
