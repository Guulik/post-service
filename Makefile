.PHONY: run-tests start stop down lint

lint:
	golangci-lint run -c ./configure/.golangci.yml

start:
	docker-compose up -d

down:
	docker-compose down

stop:
	docker-compose stop

run-tests:
	go test ./...

