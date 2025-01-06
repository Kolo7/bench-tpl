package utils

import (
	"testing"
)

func TestParseOrder(t *testing.T) {
	tests := []struct {
		order      string
		fieldNames []string
		expected   string
	}{
		{"name", []string{"name", "age"}, "name"},
		{"AGE", []string{"name", "age"}, "age"},
		{"name DESC", []string{"name", "age"}, "name desc"},
		{"AGE ASC", []string{"name", "age"}, "age asc"},
		{"unknown", []string{"name", "age"}, "id desc"},
		{"", []string{"name", "age"}, "id desc"},
		{"name asc", []string{"name", "age", "unknown"}, "name asc"},
		{"age DESC", []string{"name", "age", "unknown"}, "age desc"},
		{"unknown asc", []string{"name", "age"}, "id desc"},
		{"name DESC age", []string{"name", "age"}, "name desc"},
		{"age name asc", []string{"name", "age"}, "name asc"},
		{"desc", []string{}, "id desc"},
		{"asc", []string{}, "id desc"},
		{"name age", []string{"name", "age"}, "name"},
		{"name, age", []string{"name", "age"}, "name,age"},
		{"name, age DESC", []string{"name", "age"}, "name,age desc"},
		{"name, age DESC, unknown", []string{"name", "age"}, "name,age desc"},
		{"unknown ASC, name, age DESC", []string{"name", "age"}, "name,age desc"},
		{"sdnasuidnas", []string{"name", "age"}, "id desc"},
		{"412,,312sda,12", []string{"name", "age"}, "id desc"},
		{"sda,name,s2,,w21,,213,desc", []string{"name", "age"}, "name"},
		{"desc desc asc", []string{"name", "age"}, "id desc"},
	}

	for _, test := range tests {
		actual := parseOrder(test.order, test.fieldNames)
		if actual != test.expected {
			t.Errorf("ParseOrder(%q, %v) = %q, expected %q", test.order, test.fieldNames, actual, test.expected)
		}
	}
}
