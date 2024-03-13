APP=pack-svc
APP_VERSION:=0.1
APP_COMMIT:=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./out/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

CONFIG_FILE="./.env"
HTTP_SERVE_COMMAND="http-serve"

setup: copy-config

compile:
	mkdir -p out/
	go build -ldflags "-X main.version=$(APP_VERSION) -X main.commit=$(APP_COMMIT)" -o $(APP_EXECUTABLE) cmd/*.go

build: deps compile

http-serve: build
	$(APP_EXECUTABLE) -configFile=$(configFile) $(HTTP_SERVE_COMMAND)

app:
	docker-compose -f build/docker-compose.pack-svc.yml -f build/docker-compose.network.yml up -d --build
	docker logs -f pack-svc-go

http-local-serve: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE) $(HTTP_SERVE_COMMAND)

copy-config:
	cp .env.sample .env

tidy:
	go mod tidy

deps:
	go mod download

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

clean:
	rm -rf out/


build-test-container:
	docker build -t go-build go-build

test: build-test-container
	docker run --rm -v $$(pwd):/project -w /project go-build -c "go clean -testcache ; go test ./pkg/..."

functional-test:
	docker exec -w /project -ti functional-test-go  sh -c "go clean -testcache ; go test ./functionaltest/..."
test-cover-html:
	go clean -testcache
	mkdir -p out/
	go test ./... -coverprofile=out/coverage.out
	go tool cover -html=out/coverage.out

ci-test: test

lint:
	golangci-lint run cmd/... pkg/...

