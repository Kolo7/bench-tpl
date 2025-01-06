package utils

import (
	"strings"
	"unicode"

	"github.com/samber/lo"
)

func ToUpperCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, "")
}

func ToLowerCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = strings.Title(strings.ToLower(word))
		}
	}
	return strings.Join(words, "")
}

func ToSnakeCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "_")
}

func parseOrder(order string, fieldNames []string) string {
	order = strings.ToLower(order)
	orders := strings.Split(order, ",")
	result := make([]string, 0)
	for _, order := range orders {
		result = append(result, parseOrderSingle(strings.TrimSpace(order), fieldNames))
	}
	result = lo.Filter(result, func(s string, _ int) bool { return s != "" })
	if len(result) == 0 {
		return "id desc"
	}
	return strings.Join(result, ",")
}

func parseOrderSingle(order string, fieldNames []string) string {
	result := ""
	for _, col := range fieldNames {
		if strings.Contains(order, col) {
			result = col
			break
		}
	}
	if result == "" {
		return ""
	}
	if strings.Contains(order, "desc") {
		result = result + " desc"
	} else if strings.Contains(order, "asc") {
		result = result + " asc"
	}

	return result
}
