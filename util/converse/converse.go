package converse

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// return Key=Value
func ParseToStringKeyValue(key string, value interface{}) string {
	return fmt.Sprintf(" %v=%+v", key, value)
}

// return any type to string Value
func ParseToString(value interface{}) string {
	return fmt.Sprintf("%+v", value)
}

// return string to int Value
func ParseStringToInt(value string) int {
	valueInt, _ := strconv.Atoi(value)
	return valueInt
}

func JsonMarshalIndent(data interface{}) string {
	jsonMarshalIndent, _ := json.MarshalIndent(data, "\t", "\t")
	return fmt.Sprintf("%v", string(jsonMarshalIndent))
}

func JsonMarshal(data interface{}) string {
	jsonMarshal, _ := json.Marshal(data)
	return fmt.Sprintf("%v", string(jsonMarshal))
}

func JsonMarshalToByte(data interface{}) []byte {
	jsonMarshal, _ := json.Marshal(data)
	return jsonMarshal
}

func JsonUnmarshal(data string) interface{} {
	var response interface{}
	json.Unmarshal([]byte(data), &response)
	return response
}

func ParsVersionToInt64(v string) (int64, error) {
	sections := strings.Split(v, ".")
	mergeSection := func(v string, n int) string {
		if n < len(sections) {
			return fmt.Sprintf("%04s", sections[n])
		} else {
			return "0000"
		}
	}
	s := ""
	for i := 0; i < 4; i++ {
		s += mergeSection(v, i)
	}
	return strconv.ParseInt(s, 10, 64)
}
