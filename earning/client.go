package earning

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/iter"
)

// Iter is an iterator for a list of quotes.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*iter.Iter
}

// Quote returns the most recent Quote
// visited by a call to Next.
func (i *Iter) EarningHistory() *[]History {
	return i.Current().(*[]History)
}

// Returns Earning History for a symbol.

func GetEarningHistory(symbol string) ([]*History, error) {
	resp, err := http.Get(finance.YFinURL + "/v10/finance/quoteSummary/" + symbol + "?modules=earningsHistory")
	if err != nil {
		err = finance.CreateRemoteError(err)
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
