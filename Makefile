.PHONY: lint

lint:
	golangci-lint run -c ./configure/.golangci.yml

