SHELL := /bin/bash
MAKEFLAGS += --warn-undefined-variables
.DEFAULT_GOAL := build
.PHONY: *

VERSION := $(shell jq .version package.json | xargs)

install-dep:
	@echo "==> Installing dep"
	@go get -u github.com/golang/dep/cmd/dep

install-gometalinter:
	@echo "==> Installing gometalinter"
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

build:
	@echo "==> Building terraform-provider-netbox"
	@go build -mod=vendor -o terraform-provider-netbox .

build-all: clean
	@echo "==> Building terraform-provider-netbox"
	@for GOOS in darwin linux windows; do for GOARCH in 386 amd64; do env GOOS=$${GOOS} GOARCH=$${GOARCH} env VERSION=${VERSION} go build -mod=vendor -o terraform-provider-netbox_v${VERSION} . && tar zcf terraform-provider-netbox_v${VERSION}_$${GOOS}_$${GOARCH}.tar.gz terraform-provider-netbox_v${VERSION} && rm terraform-provider-netbox_v${VERSION}; done; done
	@for file in *.tar.gz; do echo "$$(sha256sum $${file})" >> sha256sums.txt; done

localinstall:
	@echo "==> Creating folder ~/.terraform.d/plugins/registry.terraform.io/smutel/netbox/0.0.1/linux_amd64"
	@mkdir -p ~/.terraform.d/plugins/registry.terraform.io/smutel/netbox/0.0.1/linux_amd64
	@echo "==> Installing provider in this folder"
	@cp terraform-provider-netbox ~/.terraform.d/plugins/registry.terraform.io/smutel/netbox/0.0.1/linux_amd64/terraform-provider-netbox_v0.0.1

check:
	@echo "==> Checking terraform-provider-netbox"
	@golangci-lint run

clean:
	@echo "==> Cleaning terraform-provider-netbox"
	@rm -f terraform-provider-netbox
	@rm -rf *.tar.gz
	@rm -f sha256sums.txt
