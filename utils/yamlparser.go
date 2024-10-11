package utils

import (
	"bytes"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/lang"
	"gopkg.in/yaml.v2"
)

func YamlToJson(data []byte) ([]byte, error) {
	var val any
	if err := yaml.Unmarshal(data, &val); err != nil {
		return nil, err
	}

	return encodeToJSON(toStringKeyMap(val))
}

// encodeToJSON encodes the given value into its JSON representation.
func encodeToJSON(val any) ([]byte, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// convertKeyToString ensures all keys of the map are of type string.
func convertKeyToString(in map[any]any) map[string]any {
	res := make(map[string]any)
	for k, v := range in {
		res[lang.Repr(k)] = toStringKeyMap(v)
	}
	return res
}

func toStringKeyMap(v any) any {
	switch v := v.(type) {
	case []any:
		return convertSlice(v)
	case map[any]any:
		return convertKeyToString(v)
	case bool, string:
		return v
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		return convertNumberToJsonNumber(v)
	default:
		return lang.Repr(v)
	}
}

func convertSlice(in []any) []any {
	res := make([]any, len(in))
	for i, v := range in {
		res[i] = toStringKeyMap(v)
	}
	return res
}

// convertNumberToJsonNumber converts numbers into json.Number type for compatibility.
func convertNumberToJsonNumber(in any) json.Number {
	return json.Number(lang.Repr(in))
}
