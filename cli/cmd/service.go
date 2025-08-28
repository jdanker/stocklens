package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func serviceBaseURL() string {
	if v := os.Getenv("STOCKLENS_SERVICE_URL"); v != "" {
		return v
	}
	return "http://localhost:8000"
}

func fetchJSON(path string, out interface{}) error {
	endpoint := fmt.Sprintf("%s%s", serviceBaseURL(), path)
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service error: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func priceURL(ticker string) string {
	return fmt.Sprintf("/price?ticker=%s", url.QueryEscape(ticker))
}

func analyzeURL(ticker string) string {
	return fmt.Sprintf("/analyze?ticker=%s", url.QueryEscape(ticker))
}
