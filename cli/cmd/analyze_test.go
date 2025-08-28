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

func TestAnalyzeCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/analyze" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		fmt.Fprint(w, `{"ticker":"AAPL","price":123.45,"pe":20.0}`)
	}))
	defer ts.Close()

	os.Setenv("STOCKLENS_SERVICE_URL", ts.URL)
	defer os.Unsetenv("STOCKLENS_SERVICE_URL")

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"analyze", "AAPL"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "P/E") {
		t.Fatalf("expected output to contain analysis, got %s", out)
	}
}
