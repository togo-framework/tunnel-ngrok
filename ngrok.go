// Package ngrok is a togo tunnel driver for ngrok, using the pure-Go ngrok-go
// SDK (no external binary). It opens an HTTPS ingress and forwards it to a local
// address.
//
// Install: `togo install togo-framework/tunnel-ngrok`, set TUNNEL_DRIVER=ngrok
// and NGROK_AUTHTOKEN (from https://dashboard.ngrok.com).
package ngrok

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"sync"

	ngrok "golang.ngrok.com/ngrok"
	ngrokconfig "golang.ngrok.com/ngrok/config"

	"github.com/togo-framework/togo"
	"github.com/togo-framework/tunnel"
)

func init() {
	tunnel.RegisterDriver("ngrok", func(k *togo.Kernel) (tunnel.Tunnel, error) {
		if os.Getenv("NGROK_AUTHTOKEN") == "" {
			return nil, fmt.Errorf("tunnel-ngrok: NGROK_AUTHTOKEN not set")
		}
		return &driver{domain: os.Getenv("NGROK_DOMAIN")}, nil
	})
}

// upstreamURL builds the local upstream URL an ngrok ingress forwards to.
func upstreamURL(addr string) (*url.URL, error) {
	return url.Parse("http://" + tunnel.NormalizeAddr(addr))
}

type driver struct {
	domain string

	mu  sync.Mutex
	fwd ngrok.Forwarder
	url string
}

func (d *driver) Start(ctx context.Context, addr string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.fwd != nil {
		return d.url, nil
	}
	upstream, err := upstreamURL(addr)
	if err != nil {
		return "", err
	}

	endpoint := ngrokconfig.HTTPEndpoint()
	if d.domain != "" {
		endpoint = ngrokconfig.HTTPEndpoint(ngrokconfig.WithDomain(d.domain))
	}

	fwd, err := ngrok.ListenAndForward(
		ctx,
		upstream,
		endpoint,
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return "", fmt.Errorf("tunnel-ngrok: %w", err)
	}
	d.fwd = fwd
	d.url = fwd.URL()
	return d.url, nil
}

func (d *driver) Stop(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.fwd == nil {
		return nil
	}
	err := d.fwd.CloseWithContext(ctx)
	d.fwd = nil
	d.url = ""
	return err
}

func (d *driver) Status(context.Context) (tunnel.Status, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return tunnel.Status{Running: d.fwd != nil, URL: d.url}, nil
}
