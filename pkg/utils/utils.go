package utils

import (
	"encoding/json"
	"net/http"
)

// ParseBody parses the request body and returns a map.
func ParseBody(r *http.Request) (map[string]interface{}, error) {
    var body map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

// MarshalJSON marshals a Go data structure to JSON.
func MarshalJSON(data interface{}) ([]byte, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    return jsonData, nil
}
