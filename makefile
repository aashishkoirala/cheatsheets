SHELL := /bin/bash

all: generateHtml

.PHONY: generateHtml
generateHtml:
	@go run generator.go > docs/index.html
