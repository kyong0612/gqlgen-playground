
.PHONY: gen
gen:
	go run github.com/99designs/gqlgen generate


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
