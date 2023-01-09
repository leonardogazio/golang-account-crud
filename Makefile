PKGS ?= $(shell go list ./server/routes/handlers/testing/...)

postgres:
	docker-compose up -d psql
	docker exec crudtest-psql sh /tmp/health.sh

migrate_up:
	migrate -database "postgresql://postgres:123@localhost/crudtest?sslmode=disable" -path db/migrations up

seed:
	docker exec crudtest-psql psql -d crudtest -f /tmp/seed.sql

unit-test:
	go test -ldflags="-extldflags=-Wl,--allow-multiple-definition" -v ${PKGS}

build:
	go build -o ./bin/app

build-x:
	docker run --rm -v $(shell pwd):/go/src/app -w /go/src/app leonardogazio/go:1.18 gox -osarch="linux/amd64" -output="app"

run:
	go run .