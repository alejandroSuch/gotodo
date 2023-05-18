package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Method      string
	Url         string
	RequestBody io.Reader
	ApiResponse any
	AuthToken   string
}

func DoRequest(r Request) error {
	req, err := http.NewRequest(r.Method, r.Url, r.RequestBody)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if r.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+r.AuthToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error when making the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server response error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, r.ApiResponse)
	if err != nil {
		return fmt.Errorf("error while processing response: %s", resp.Status)
	}

	return nil
}
