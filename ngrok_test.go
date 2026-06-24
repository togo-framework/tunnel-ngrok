package ngrok

import (
	"testing"

	"github.com/togo-framework/tunnel"
)

func TestUpstreamURL(t *testing.T) {
	cases := map[string]string{
		"8080":           "http://127.0.0.1:8080",
		":3000":          "http://127.0.0.1:3000",
		"localhost:5173": "http://localhost:5173",
	}
	for in, want := range cases {
		u, err := upstreamURL(in)
		if err != nil {
			t.Fatalf("upstreamURL(%q): %v", in, err)
		}
		if u.String() != want {
			t.Errorf("upstreamURL(%q) = %q, want %q", in, u.String(), want)
		}
	}
}

func TestDriverRegistered(t *testing.T) {
	found := false
	for _, n := range tunnel.Drivers() {
		if n == "ngrok" {
			found = true
		}
	}
	if !found {
		t.Fatal("ngrok driver not registered on tunnel base")
	}
}
