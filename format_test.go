package timeutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParseDay(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      string
		expect int64
	}{
		{
			t:      "1d",
			expect: int64(time.Hour * 24),
		},
		{
			t:      "1w",
			expect: int64(time.Hour * 24 * 7),
		},
	}

	for i, test := range tests {
		actual, err := ParseDuration(test.t)
		assert.NoError(err, fmt.Sprintf("No. %d", i+1))
		assert.Equal(test.expect, int64(actual))
	}
}
