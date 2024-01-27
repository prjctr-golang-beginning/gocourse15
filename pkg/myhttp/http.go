package myhttp

import (
	"context"
	"github.com/samber/do"
	"net/http"
	"time"
)

// HTTP is the http server
type HTTP struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

// NewHTTP creates a new server
func NewHTTP(srv *http.Server, st time.Duration) *HTTP {
	return &HTTP{srv: srv, shutdownTimeout: st}
}

// Start starts the server
func (s HTTP) Start() error {
	return s.srv.ListenAndServe()
}

// Stop gracefully stops the server
// https://golang.org/pkg/net/http/#Server.Shutdown
func (s HTTP) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s HTTP) Shutdown() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancelFunc()
	return s.Stop(ctx)
}

// ServerProvider provides a new http server
func ServerProvider(_ *do.Injector) (*HTTP, error) {
	serv := &http.Server{}

	return NewHTTP(serv, 1*time.Second), nil
}
