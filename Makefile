SHELL   := /bin/zsh -euo pipefail
TIMEOUT := 10s
BINPATH = $(PWD)/bin

.PHONY: deps
deps:
	@go mod tidy && go mod verify