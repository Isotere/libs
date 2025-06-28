SHELL   := /bin/zsh -euo pipefail
TIMEOUT := 10s
BINPATH = $(PWD)/bin

.PHONY: deps
deps:
	@go mod tidy && go mod verify


.PHONY: bench
bench: bench_errors

.PHONY: bench_errors
bench_errors:
	cd ./errors && go test -bench=. -benchmem -benchtime=10s -count=5