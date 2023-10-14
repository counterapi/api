export PATH := $(abspath ./vendor/bin):$(PATH)

K8S_DEPLOYMENT_FILE  = kubernetes/application/deployment.yaml
IMAGE_NAME  		 = ghcr.io/counterapi/counterapi
BASE_PACKAGE_NAME    = github.com/counterapi/api
GIT_VERSION          = $(shell git describe --tags --always 2> /dev/null || echo 0.0.0)
LDFLAGS              = -ldflags "-X $(BASE_PACKAGE_NAME)/pkg/info.Version=$(GIT_VERSION)"
BUFFER               := $(shell mktemp)
REPORT_DIR           = dist/report
COVER_PROFILE        = $(REPORT_DIR)/coverage.out
BINARY_NAME   		 = dist/counterapi

.PHONY: build
build:
	CGO_ENABLED=0 GOOS="$(TARGETOS)" GOARCH="$(TARGETARCH)" go build $(LDFLAGS) -a -installsuffix cgo -o $(BINARY_NAME) main.go

.PHONY: lint
lint:
	@echo "Checking code style"
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)
	go vet ./...
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1
	@golangci-lint --version
	golangci-lint run

.PHONY: test
test:
	@echo "Running unit tests"
	mkdir -p $(REPORT_DIR)
	go test -covermode=count -coverprofile=$(COVER_PROFILE) -tags test -failfast ./...
	go tool cover -html=$(COVER_PROFILE) -o $(REPORT_DIR)/coverage.html

.PHONY: cut-tag
cut-tag:
	@echo "Cutting $(version)"
	git tag $(version)
	git push origin $(version)

.PHONY: release
release: build-for-container
	@echo "Releasing $(GIT_VERSION)"
	docker build -t counter .
	docker tag counter:latest counterapi/counterapi:$(GIT_VERSION)
	docker push counterapi/counter:$(GIT_VERSION)

.PHONY: dev
dev:
	gin --appPort 80 --port 8000 -i

.PHONY: dev-db
dev-db:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=root -d postgres

.PHONY: set-db-variables
set-db-variables:
	export DB_HOST=localhost
	export DB_PORT=5432
	export DB_USER=postgres
	export DB_NAME=counter_api
	export DB_PASSWORD=root
