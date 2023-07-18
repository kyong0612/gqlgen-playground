
.PHONY: generate.gqlgen
generate.gqlgen:
	go run github.com/99designs/gqlgen generate
	go mod tidy

.PHONY: run
run:
	go build && ./gqlgen-playground

.PHONY: compose.up
compose.up:
	@docker compose up

.PHONY: compose.down
compose.down:
	@docker compose down

.PHONY: compose.reset
compose.reset: 
	rm -rf postgres/db-data
	make compose.down 
	make compose.up

.PHONY: generate.sqlboiler
generate.sqlboiler:
	go run github.com/volatiletech/sqlboiler/v4 --config=sqlboiler.toml psql
	go mod tidy
