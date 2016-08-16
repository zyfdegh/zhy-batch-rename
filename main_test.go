package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestStringfy(t *testing.T) {
	var cases = []struct {
		in     string
		expect string
	}{
		{"", ""},
		{"HELLO", "Hello"},
		{"R0001D1D5", "R0001d1d5"},
		{"C0005793A.png", "C0005793a.png"},
	}

	for _, c := range cases {
		//call
		got := stringfy(c.in)
		//assert
		assert.Equal(t, c.expect, got)
	}
}
