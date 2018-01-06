package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var loc = time.Local

func Test_GreaterThan(t *testing.T) {
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
		assert.Equal(test.expect, (Time{test.t}).GreaterThan(test.other), i)
		assert.Equal(test.expect, (Time{test.t}).GT(test.other), i)
	}
}

func Test_GreaterEqual(t *testing.T) {
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
		assert.Equal(test.expect, (Time{test.t}).GreaterEqual(test.other), i)
		assert.Equal(test.expect, (Time{test.t}).GE(test.other), i)
	}
}

func Test_LessThan(t *testing.T) {
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
		assert.Equal(test.expect, (Time{test.t}).LessThan(test.other), i)
		assert.Equal(test.expect, (Time{test.t}).LT(test.other), i)
	}
}

func Test_LessEqual(t *testing.T) {
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
		assert.Equal(test.expect, (Time{test.t}).LessEqual(test.other), fmt.Sprintf("No. %d", i+1))
		assert.Equal(test.expect, (Time{test.t}).LE(test.other), fmt.Sprintf("No. %d", i+1))
	}
}

func Test_Equal(t *testing.T) {
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
		assert.Equal(test.expect, (Time{test.t}).Equal(test.other), fmt.Sprintf("No. %d", i+1))
		assert.Equal(test.expect, (Time{test.t}).EQ(test.other), fmt.Sprintf("No. %d", i+1))
	}
}

func Test_Between(t *testing.T) {
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
		assert.Equal(test.expect, (Time{test.t}).Between(test.t1, test.t2), fmt.Sprintf("No. %d", i+1))
	}
}
