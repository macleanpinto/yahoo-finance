package eps

import (
	"testing"

	finance "github.com/piquette/finance-go"
	"github.com/stretchr/testify/assert"
)

func TestGetEpsHistory(t *testing.T) {
	response, err := getEpsHistory("RELIANCE.NS")
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
func TestGetEpsHistoryBadQuote(t *testing.T) {
	response, err := getEpsHistory("CRAP.NS")
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.Equal(t, finance.CreateRemoteErrorS("error response recieved from upstream api"), err)
}
