package apis

import (
	"bytes"
	"io"
	"net/http"
)

func JsonPost(url string, jsonByte []byte) ([]byte, error) {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, err
}
