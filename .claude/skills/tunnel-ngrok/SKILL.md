---
name: tunnel-ngrok
description: Expose a local togo app via ngrok — set TUNNEL_DRIVER=ngrok + NGROK_AUTHTOKEN and call tunnel.Start
---

# togo tunnel-ngrok

ngrok driver for the togo `tunnel` subsystem (ngrok-go SDK — no separate binary
needed).

## Setup

```bash
togo install togo-framework/tunnel
togo install togo-framework/tunnel-ngrok
```

Get an authtoken from the [ngrok dashboard](https://dashboard.ngrok.com), then
`.env`:

```bash
TUNNEL_DRIVER=ngrok
NGROK_AUTHTOKEN=2abc...
# NGROK_DOMAIN=myapp.ngrok.app   # optional reserved domain
```

## Use

```go
import (
	_ "github.com/togo-framework/tunnel"
	_ "github.com/togo-framework/tunnel-ngrok"
	"github.com/togo-framework/tunnel"
)

if tn, ok := tunnel.FromKernel(k); ok {
	url, _ := tn.Start(ctx, "8080") // https://<...>.ngrok-free.app
	defer tn.Stop(ctx)
}
```

## Notes
- Uses the ngrok-go SDK — no `ngrok` binary required.
- Set `NGROK_DOMAIN` to use a reserved/custom domain instead of a random one.
- Keep `NGROK_AUTHTOKEN` in `.env`; never commit it.
