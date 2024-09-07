package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLength(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"varchar(255)", "varchar"},
		{"int(11)", "int"},
		{"decimal(10,2)", "decimal"},
		{"text", "text"},
		{"", ""},
		{"varchar", "varchar"},
		{"int", "int"},
		{"decimal", "decimal"},
		{"varchar(255)int(11)", "varcharint"},
		{"varchar(255)int(11)decimal(10,2)", "varcharintdecimal"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := removeLength(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
