
.PHONY: gen
gen:
	go run github.com/99designs/gqlgen generate


.PHONY: run
run:
	go build && ./gqlgen-playground
