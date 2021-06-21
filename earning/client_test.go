package earning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEarningHistory(t *testing.T) {
	response, err := GetEarningHistory("RELIANCE.NS")
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
