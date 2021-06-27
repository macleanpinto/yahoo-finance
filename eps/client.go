package eps

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	finance "github.com/piquette/finance-go"
)

// Earning history for the most recent four quarters.
func GetEpsHistory(symbol string) ([]*History, error) {
	resp, err := http.Get(finance.YFinURL + "/v10/finance/quoteSummary/" + symbol + "?modules=earningsHistory")
	if err != nil {
		err = finance.CreateRemoteError(err)
		return nil, err
	}

	if resp.StatusCode >= 400 {
		err = finance.CreateRemoteErrorS("error response recieved from upstream api")
		return nil, err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	result := response{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.QuoteSummary.Error != nil {
		err = finance.CreateRemoteError(result.QuoteSummary.Error)
		return nil, err
	}
	return result.QuoteSummary.Result[0].EarningsHistory.History, err
}

// response is a yfin quote response.
type response struct {
	QuoteSummary struct {
		Result []Result           `json:"result"`
		Error  *finance.YfinError `json:"error"`
	} `json:"quoteSummary"`
}

type History struct {
	MaxAge    int `json:"maxAge"`
	EpsActual struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"epsActual"`
	EpsEstimate struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"epsEstimate"`

	EpsDifference struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"epsDifference"`

	SurprisePercent struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"surprisePercent"`

	Quarter struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"quarter"`
}

type Result struct {
	EarningsHistory struct {
		History []*History `json:"history"`
	} `json:"earningsHistory"`
}
