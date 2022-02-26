package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_mainHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mainHandler))
	defer ts.Close()
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	result, ioErr := io.ReadAll(r.Body)
	if ioErr != nil {
		t.Fatal(ioErr)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if !strings.Contains(string(result), "container") {
		t.Logf("%s", result)
		t.Fatal("not in container")
	}
}
