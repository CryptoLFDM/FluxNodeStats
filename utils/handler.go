package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// DoRequest Execute a request. If token provided, it's added as Authorization.
func DoRequest(method, url string, data interface{}, token ...string) (*http.Response, error) {
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(method, url, body)
	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token[0])
	}
	client := http.Client{}
	return client.Do(req)
}
