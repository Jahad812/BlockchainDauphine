package httptest

import (
	"context"
	"net"
	"net/http"
	"time"

	pkgerrors "github.com/pkg/errors"
)

// NewTestHTTPClient returns a real HTTP client that may only make requests to
// localhost
func NewTestLocalOnlyHTTPClient() *http.Client {
	// Don't use the default transport, we want zero limits and zero timeouts
	tr := &http.Transport{}
	tr.DialContext = testDialContext
	tr.DisableCompression = true
	return &http.Client{Transport: tr}
}

func testDialContext(ctx context.Context, network, address string) (net.Conn, error) {
	con, err := (&net.Dialer{
		// Defaults from GoLang standard http package
		// https://golang.org/pkg/net/http/#RoundTripper
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext(ctx, network, address)
	if err != nil {
		return con, err
	}
	a := con.RemoteAddr().(*net.TCPAddr)
	if a != nil && !a.IP.IsLoopback() {
		return nil, pkgerrors.Errorf("Test HTTP client may only dial localhost, got address: %v", a.String())
	}
	return con, err
}
