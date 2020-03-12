BINARY=url_shortener_api

test:
	go test -v -cover -race -covermode=atomic -timeout 30s ./...

build:
	go build -v -o ${BINARY} ./cmd/${BINARY}

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose up -d

stop:
	docker-compose down

run:
	go run ./cmd/${BINARY}/main.go

.DEFAULT_GOAL := build

.PHONY: clean unittest build docker run stop