TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')
PKG_NAME=cloudflare
VERSION?=$(shell git describe --tags --always)
DEV_VERSION=99.0.0
CLOUDFLARE_GO_VERSION?=master
INCLUDE_VERSION_IN_FILENAME?=false

default: build

install: vet fmtcheck
	go install -ldflags="-X github.com/cloudflare/terraform-provider-cloudflare/main.version=$(VERSION)"

build: vet fmtcheck
	@if $(INCLUDE_VERSION_IN_FILENAME); then \
	    go build -ldflags="-X github.com/cloudflare/terraform-provider-cloudflare/main.version=$(VERSION)" -o terraform-provider-cloudflare_$(VERSION); \
		echo "==> Successfully built terraform-provider-cloudflare_$(VERSION)"; \
	else \
		go build -ldflags="-X github.com/cloudflare/terraform-provider-cloudflare/main.version=$(VERSION)" -o terraform-provider-cloudflare; \
		echo "==> Successfully built terraform-provider-cloudflare"; \
	fi

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(TEST) -v -sweep=$(SWEEP) $(SWEEPARGS)

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4 -race

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -parallel 1

lint: tools terraform-provider-lint golangci-lint

terraform-provider-lint: tools
	$$(go env GOPATH)/bin/tfproviderlintx \
	 -R001=false \
	 -R003=false \
	 -R012=false \
	 -S006=false \
	 -S014=false \
	 -S020=false \
	 -S022=false \
	 -S023=false \
	 -AT001=false \
	 -AT002=false \
	 -AT003=false \
	 -AT006=false \
	 -AT012=false \
	 -R013=false \
	 -XAT001=false \
	 -XR001=false \
	 -XR003=false \
	 -XR004=false \
	 -XS001=false \
	 -XS002=false \
	 ./...

vet:
	@echo "==> Running go vet ."
	@go vet ./... ; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

clean-dev:
	@echo "==> Removing development version ($(DEV_VERSION))"
	@rm -f terraform-provider-cloudflare_$(DEV_VERSION)

build-dev: clean-dev
	@echo "==> Building development version ($(DEV_VERSION))"
	go build -gcflags="all=-N -l" -o terraform-provider-cloudflare_$(DEV_VERSION)

generate-changelog:
	@echo "==> Generating changelog..."
	@sh -c "'$(CURDIR)/scripts/generate-changelog.sh'"

golangci-lint:
	@golangci-lint run ./internal/... --config .golintci.yml

tools:
	@echo "==> Installing development tooling..."
	go generate -tags tools tools/tools.go

update-go-client:
	@echo "==> Updating the cloudflare-go client to $(CLOUDFLARE_GO_VERSION)"
	go get github.com/cloudflare/cloudflare-go@$(CLOUDFLARE_GO_VERSION)
	go mod tidy

docs: tools
	@sh -c "'$(CURDIR)/scripts/generate-docs.sh'"

.PHONY: build install test sweep testacc lint terraform-provider-lint vet fmt fmtcheck errcheck test-compilebuild-dev clean-dev generate-changelog golangci-lint tools update-go-client docs
