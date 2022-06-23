BIN=bin/opt/payment-service-emulator/payment-service-emulator
DOCKER_IMG=emulator
RELEASE=develop

PACKAGE_PATH := "github.com/Feride3d/payment-service-emulator"

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/...

run: build
	$(BIN) -config ./config/config.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG):$(RELEASE) \
		-f build/Dockerfile .

run-img: build-img
	docker rm $(DOCKER_IMG) || true
	docker run \
	--name $(DOCKER_IMG) \
	-p 8080:8080 \
	$(DOCKER_IMG):$(RELEASE)

version: build
	$(BIN) version

test:
	go test -race -count 100 ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build run build-img run-img version test lint