package timex

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"fmt"
)

func Test_TimexGreaterThan(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      time.Time
		other  time.Time
		expect bool
	}{
		{
			t:      time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: false,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			expect: false,
		},
	}

	for i, test := range tests {
		assert.Equal(test.expect, A(test.t).GreaterThan().B(test.other), fmt.Sprintf("No. %d", i+1))
	}
}

func Test_TimexGreaterEqual(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      time.Time
		other  time.Time
		expect bool
	}{
		{
			t:      time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			expect: false,
		},
	}

	for i, test := range tests {
		assert.Equal(test.expect, A(test.t).GreaterEqual().B(test.other), fmt.Sprintf("No. %d", i+1))
	}
}

func Test_TimexLessThan(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      time.Time
		other  time.Time
		expect bool
	}{
		{
			t:      time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: false,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: false,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			expect: true,
		},
	}

	for i, test := range tests {
		assert.Equal(test.expect, A(test.t).LessThan().B(test.other), fmt.Sprintf("No. %d", i+1))
	}
}

func Test_TimexLessEqual(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      time.Time
		other  time.Time
		expect bool
	}{
		{
			t:      time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: false,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			expect: true,
		},
	}

	for i, test := range tests {
		assert.Equal(test.expect, A(test.t).LessEqual().B(test.other), fmt.Sprintf("No. %d", i+1))
	}
}


func Test_TimexEqual(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      time.Time
		other  time.Time
		expect bool
	}{
		{
			t:      time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: false,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			other:  time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			expect: true,
		},
	}

	for i, test := range tests {
		assert.Equal(test.expect, A(test.t).Equal().B(test.other), fmt.Sprintf("No. %d", i+1))
	}
}
func Test_TimexBetween(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		t      time.Time
		t1     time.Time
		t2     time.Time
		expect bool
	}{
		{
			t:      time.Date(2017, 12, 18, 0, 0, 0, 0, loc),
			t1:     time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			t2:     time.Date(2017, 12, 18, 0, 0, 3, 0, loc),
			expect: false,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			t1:     time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			t2:     time.Date(2017, 12, 18, 0, 0, 3, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 2, 0, loc),
			t1:     time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			t2:     time.Date(2017, 12, 18, 0, 0, 3, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 2, 0, loc),
			t1:     time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			t2:     time.Date(2017, 12, 18, 0, 0, 3, 0, loc),
			expect: true,
		},
		{
			t:      time.Date(2017, 12, 18, 0, 0, 4, 0, loc),
			t1:     time.Date(2017, 12, 18, 0, 0, 1, 0, loc),
			t2:     time.Date(2017, 12, 18, 0, 0, 3, 0, loc),
			expect: false,
		},
	}

	for i, test := range tests {
		assert.Equal(test.expect, X(test.t).Between().Y(test.t1).And().Z(test.t2), fmt.Sprintf("No. %d", i+1))
	}
}