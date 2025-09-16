package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func ToJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func FromJSON(jsonStr string, v interface{}) error {
	return json.Unmarshal([]byte(jsonStr), v)
}

func ValidateRequired(fields map[string]interface{}) error {
	for name, value := range fields {
		if value == nil || value == "" {
			return fmt.Errorf("field %s is required", name)
		}
	}
	return nil
}
