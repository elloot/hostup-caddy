Hostup module for Caddy
=======================

[![Go Reference](https://pkg.go.dev/badge/github.com/elloot/hostup-caddy.svg)](https://pkg.go.dev/github.com/elloot/hostup-caddy)

This package contains a DNS provider module for Caddy, `dns.providers.hostup`. It wraps the [libdns](https://github.com/libdns/libdns) provider at [`github.com/elloot/hostup-libdns`](https://github.com/elloot/hostup-libdns), which manages DNS records through the public [Hostup API v2](https://developer.hostup.se/) (`https://cloud.hostup.se/api/v2`).

It is used to solve the ACME DNS-01 challenge: Caddy creates and removes the `_acme-challenge` TXT record via this module.

## Caddy module name

```
dns.providers.hostup
```

## Building Caddy with this module

```sh
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
xcaddy build --with github.com/elloot/hostup-caddy
```

With Docker:

```dockerfile
FROM caddy:builder AS builder
RUN xcaddy build --with github.com/elloot/hostup-caddy

FROM caddy:latest
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
```

## Configuring

Create an API key at **cloud.hostup.se → API management** with the `read:dns` and `write:dns` scopes, and expose it as `HOSTUP_API_KEY`.

Caddyfile, globally:

```caddy
{
	acme_dns hostup {
		api_token {$HOSTUP_API_KEY}
	}
}

example.com {
	respond "Hello"
}
```

Caddyfile, for one site:

```caddy
example.com {
	tls {
		dns hostup {
			api_token {$HOSTUP_API_KEY}
		}
	}
	respond "Hello"
}
```

JSON, in the ACME issuer:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "hostup",
				"api_token": "{env.HOSTUP_API_KEY}"
			}
		}
	}
}
```
