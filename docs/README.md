# tunnel-ngrok — docs

**ngrok.** Pure-Go tunnel via the ngrok-go SDK — no binary needed.

## Install

```bash
togo install togo-framework/tunnel-ngrok
```

Registers on the [`tunnel`](https://github.com/togo-framework/tunnel) base; select it with **tunnel.provider in togo.yaml (or TUNNEL_DRIVER)**, then use **`togo tunnel`**.

## Interface

`Tunnel` — `Start(ctx, addr) -> publicURL`, `Stop`, `Status`.

## Configuration

| Env var | Description |
|---|---|
| `NGROK_AUTHTOKEN` | ngrok auth token (required). |
| `NGROK_DOMAIN` | Reserved ngrok domain to bind. Optional. |

## Usage & notes

Uses `ngrok.ListenAndForward`; returns the public `.ngrok` URL (or your reserved `NGROK_DOMAIN`).

## Example

```bash
togo tunnel:start --provider ngrok
```

## Links

- [ngrok-go](https://github.com/ngrok/ngrok-go)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/tunnel-ngrok)
