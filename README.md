<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/tunnel-ngrok</h1>
  <p>ngrok driver for togo tunnel — pure-Go SDK, no external binary.</p>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/tunnel-ngrok"><img src="https://pkg.go.dev/badge/github.com/togo-framework/tunnel-ngrok.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/tunnel-ngrok
```
<!-- /togo-header -->

**ngrok** driver for togo's [`tunnel`](https://github.com/togo-framework/tunnel)
subsystem, built on the pure-Go [ngrok-go](https://golang.ngrok.com/ngrok) SDK — no
external binary. Opens an HTTPS ingress and forwards it to your local app.

## Config

| Env | Meaning |
|-----|---------|
| `TUNNEL_DRIVER` | set to `ngrok` |
| `NGROK_AUTHTOKEN` | your ngrok authtoken (required) |
| `NGROK_DOMAIN` | a reserved ngrok domain (optional; otherwise a random `*.ngrok` URL) |

```go
svc, _ := tunnel.FromKernel(k)
url, _ := svc.Start(ctx, "8080")   // → https://<name>.ngrok-free.app
defer svc.Stop(ctx)
```

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- togo-sponsors -->
