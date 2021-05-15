export PATH := $(abspath ./vendor/bin):$(PATH)

BASE_PACKAGE_NAME	= github.com/omegion/go-ddclient
GIT_VERSION			= $(shell git describe --tags --always 2> /dev/null || echo 0.0.0)
LDFLAGS            	= -ldflags "-X $(BASE_PACKAGE_NAME)/pkg/info.Version=$(GIT_VERSION)"
BUFFER     			:= $(shell mktemp)
REPORT_DIR         	= dist/report
COVER_PROFILE      	= $(REPORT_DIR)/coverage.out
ARCH				= "amd64"

.PHONY: build
build:
	CGO_ENABLED=0 go build $(LDFLAGS) -installsuffix cgo -o dist/ddclient main.go

build-for-container:
	@echo $(ARCH)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build $(LDFLAGS) -a -installsuffix cgo -o dist/ddclient-linux main.go

.PHONY: lint
lint:
	@echo "Checking code style"
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)
	go vet ./...
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.1
	@golangci-lint --version
	golangci-lint run
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ./...

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
	docker build -t ddclient .
	docker tag ddclient:latest omegion/ddclient:$(GIT_VERSION)
	docker push omegion/ddclient:$(GIT_VERSION)
