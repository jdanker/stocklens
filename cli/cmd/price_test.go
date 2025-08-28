package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// Test that the price command calls the service and prints the price.
func TestPriceCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/price" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"ticker":"AAPL","price":123.45}`)
	}))
	defer ts.Close()

	os.Setenv("STOCKLENS_SERVICE_URL", ts.URL)
	defer os.Unsetenv("STOCKLENS_SERVICE_URL")

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"price", "AAPL"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "123.45") {
		t.Fatalf("expected output to contain price, got %s", out)
	}
}
