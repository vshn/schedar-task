# Set Shell to bash, otherwise some targets fail with dash/zsh etc.
SHELL := /bin/bash

# Disable built-in rules
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.SUFFIXES:
.SECONDARY:
.DEFAULT_GOAL := help

# General variables
include Makefile.vars.mk
# KIND module
include kind/kind.mk
# STACKGRES module
include stackgres/stackgres.mk

.PHONY: install
install: kind stackgres ## Install kind with stackgres

.PHONY: generate
generate: ## Generates code for APIs and Custom Resource Definitions.
	go run sigs.k8s.io/controller-tools/cmd/controller-gen paths=./api/... object crd:crdVersions=v1,allowDangerousTypes=true output:artifacts:config=./api/generated

.PHONY: deploy
deploy: export KUBECONFIG = $(KIND_KUBECONFIG)
deploy: kind-load-image ## Run your operator directly in the kind cluster
	kubectl apply -f api/generated/schedar.task.io_simplepostgresqls.yaml
	kubectl apply -f config/namespace.yaml
	kubectl apply -R -f config/

.PHONY: clean
clean: kind-clean ## Install kind
	rm -rf schedar-task

.PHONY: build
build: export CGO_ENABLED = 0
build: generate  ## Build manager binary.
	@echo "GOOS=$$(go env GOOS) GOARCH=$$(go env GOARCH)"
	go build -o $(BIN_FILENAME)

.PHONY: docker-build
docker-build: build
	env CGO_ENABLED=0 GOOS=$(DOCKER_IMAGE_GOOS) GOARCH=$(DOCKER_IMAGE_GOARCH) \
		go build -o ${BIN_FILENAME}
	docker build --platform $(DOCKER_IMAGE_GOOS)/$(DOCKER_IMAGE_GOARCH) -t ${GHCR_IMG} .
