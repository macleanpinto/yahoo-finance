package eps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEpsHistory(t *testing.T) {
	response, err := GetEpsHistory("RELIANCE.NS")
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
func TestGetEpsHistoryBadQuote(t *testing.T) {
	response, err := GetEpsHistory("CRAP.NS")
	assert.Nil(t, err)
	assert.Nil(t, response)
}
