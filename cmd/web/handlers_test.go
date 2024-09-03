package main

import (
	"io"
	"log/slog"
	"net/http"
	"testing"

	"snippetbox.mzwallow.dev/internal/assert"
)

func TestPing(t *testing.T) {
	app := &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
