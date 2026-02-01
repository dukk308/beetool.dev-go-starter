# Project Structure

Root: `golang-clean-arc/` (module: `github.com/dukk308/beetool.dev-go-starter`)

```
golang-clean-arc/
├── api-docs/
│   └── swagger/                 # Swagger (swag init): docs.go, swagger.json, swagger.yaml
├── ai-docs/                     # AI/context docs (this folder)
│   └── STRUCTURE.md
├── cmd/                         # CLI entry points
│   ├── root.go                  # Cobra root
│   ├── server.go                # serve — HTTP server
│   └── worker.go                # worker
├── database/
│   ├── migration_register.go    # Goose migration registration
│   └── migrations/              # Goose SQL migrations; Atlas diff target
│       ├── .gitkeep
│       ├── 20260127152652_create-user-table.sql
│       └── atlas.sum
├── deployment/
│   ├── development/
│   │   ├── docker-compose.yml
│   │   └── Dockerfile
│   ├── local/
│   │   ├── alloy/config.alloy
│   │   ├── docker-compose.infras.yml   # Postgres, Valkey, RabbitMQ
│   │   ├── docker-compose.o11y.yml     # Alloy, Prometheus, Tempo, Grafana
│   │   ├── docker-compose.yml
│   │   ├── Dockerfile
│   │   ├── grafana/provisioning/datasources/datasources.yml
│   │   ├── prometheus/prometheus.yml
│   │   └── tempo/tempo.yml
│   └── production/
│       └── Dockerfile
├── doc/
│   ├── clean-architecture.png
│   └── project-structure.png
├── internal/
│   ├── common/
│   │   └── response.go
│   ├── config/
│   │   ├── config.go            # AuthConfig, Config
│   │   ├── flag.go
│   │   └── fx.go
│   ├── modules/
│   │   ├── {module_name}/
│   │   │   ├── application/
│   │   │   ├── domain/
│   │   │   ├── infrastructure/
│   │   │   ├── presentation/
│   │   │   └── fx_module.go
│   │   └── fx_features.go
│   ├── server/
│   │   └── boostrap.go          # FX app, gin + swagger + modules, HTTP server lifecycle
│   └── validation/
│       ├── email_validation.go
│       └── register.go
├── pkgs/
│   ├── base/
│   │   ├── domain_error.go
│   │   └── domain_model.go
│   ├── components/
│   │   ├── cache_comp/          # Redis/Valkey (package redis_component)
│   │   │   ├── cache_service.go
│   │   │   ├── config.go
│   │   │   ├── doc.md
│   │   │   ├── flag.go
│   │   │   ├── fx.go
│   │   │   ├── redis_client.go
│   │   │   ├── redis_service.go
│   │   │   └── type.go
│   │   ├── gin_comp/
│   │   │   ├── config.go
│   │   │   ├── flag.go
│   │   │   ├── fx.go
│   │   │   ├── gin_logger.go
│   │   │   ├── gin_response.go
│   │   │   └── gin.go
│   │   ├── gorm_comp/
│   │   │   ├── audit_hook.go
│   │   │   ├── dialets/         # mssql, mysql, postgres, sqlite
│   │   │   ├── flag.go
│   │   │   ├── fx.go
│   │   │   ├── gorm.go
│   │   │   └── sql_model.go
│   │   ├── otel_comp/           # OpenTelemetry tracing
│   │   │   ├── enum.go
│   │   │   ├── factory.go
│   │   │   ├── flag.go
│   │   │   ├── fx.go
│   │   │   ├── option.go
│   │   │   ├── span_exporter.go
│   │   │   └── span_processor.go
│   │   ├── rabbitmq_comp/       # RabbitMQ (amqp091)
│   │   │   ├── config.go
│   │   │   ├── doc.md
│   │   │   ├── flag.go
│   │   │   ├── fx.go
│   │   │   ├── rabbitmq.go
│   │   │   └── type.go
│   │   ├── single_flight_comp/
│   │   │   ├── fx.go
│   │   │   ├── group.go
│   │   │   ├── redis_cache.go
│   │   │   └── type.go
│   │   └── swagger_comp/
│   │       ├── config.go
│   │       ├── flag.go
│   │       ├── fx.go
│   │       └── swagger.go
│   ├── constants/
│   │   └── context_key.go
│   ├── global_config/
│   │   ├── config.go
│   │   ├── flag.go
│   │   └── fx.go
│   ├── logger/
│   │   ├── config/log_options.go
│   │   ├── fx_event_logger.go
│   │   ├── fx_zap.go
│   │   ├── logger.go
│   │   └── zap_logger.go
│   ├── middlewares/gin/
│   │   ├── authenticate.go
│   │   ├── authorization.go
│   │   ├── correlate_logger.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   └── tracer.go
│   ├── types/
│   │   └── user_authenticated.go
│   └── utils/
│       ├── dotenv.go
│       ├── flag.go
│       └── request_id/request_id.go
├── scripts/
│   ├── bash/
│   │   └── migrate.sh
│   └── go/
│       ├── migrate/migrate.go
│       └── wait_port/wait_port.go
├── .env.example
├── .gitignore
├── .golangci.yml
├── .vscode/
│   ├── launch.json
│   └── tasks.json
├── air.toml
├── air-worker.toml
├── atlas.hcl
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## Layer mapping (Clean Architecture)

| Layer          | Path pattern                          | Role |
|----------------|----------------------------------------|------|
| Domain         | `internal/modules/*/domain/`           | Entities, value objects, repository interfaces |
| Application    | `internal/modules/*/application/`      | Commands, queries, DTOs |
| Infrastructure | `internal/modules/*/infrastructure/`   | Repo implementations, DB, external services |
| Presentation   | `internal/modules/*/presentation/http/`| HTTP handlers, routes |

## Bootstrap (FX)

- `internal/server/boostrap.go`: builds `fx.App` with `global_config`, `logger`, `config`, then `gorm_comp`, `gin_comp`, `swagger_comp`, `modules.FeatureModuleFx`, and `startHttpServer` invoke.
- Optional components (not in default bootstrap): `cache_comp`, `otel_comp`, `rabbitmq_comp`.

## Commands

- `go run main.go serve` — start HTTP server (bootstrap above).
- `go run main.go worker` — worker entry.
- `go run main.go outenv` — print env/flag help.

## Key config / flags

- `.env.example`: `SERVICE_NAME`, `APP_ENV`, `LOG_LEVEL`, `GIN_PORT`, `DB_DRIVER`, `DB_DSN`, `DB_DSN_SHADOW`, `DB_DSN_SLAVES`, etc.
- Components add flags: `-gin-port`, `-db-dsn`, `-redis-addrs`, `-rabbitmq-url`, etc.
