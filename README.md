# Go Fiber

[![Deploy on Railway](https://railway.com/button.svg)](https://railway.com/deploy/go-fiber?referralCode=dPr4mc)

A minimal Go web API built with [Fiber](https://gofiber.io), ready to deploy on Railway, with structured logging ([zap](https://github.com/uber-go/zap)) and env config ([viper](https://github.com/spf13/viper)).

## Endpoints

| Method | Path      | Description                                  |
|--------|-----------|----------------------------------------------|
| GET    | `/`       | Returns a JSON status message                |
| GET    | `/health` | Health check (used by Railway's healthcheck) |

## Environment variables

| Variable      | Default       | Description                                                   |
|---------------|---------------|---------------------------------------------------------------|
| `PORT`        | `8080`        | Port to listen on. Railway sets this automatically.           |
| `LOG_LEVEL`   | `info`        | `debug`, `info`, `warn` or `error`                            |
| `ENVIRONMENT` | `development` | `production` switches to JSON logs (also the case on Railway) |

## Run locally

Requires Go 1.27.4+ (the version in `go.mod`; Go 1.21+ toolchains fetch it automatically).

```sh
go run .
curl localhost:8080/health
```

## Deploy

Railway builds it with [Railpack](https://railpack.com) using `railway.toml`, checks `/health` before switching traffic to a new deploy, and restarts the service if it crashes.
