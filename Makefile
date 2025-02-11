.PHONY: run-tests start stop lint

lint:
	golangci-lint run -c ./configure/.golangci.yml

start:
	docker-compose up -d

stop:
	docker-compose stop

run-tests:
	go test ./...

