BUILD_DIR ?= .build
PROJECT_NAME ?= $(shell basename $(dir $(abspath $(firstword $(MAKEFILE_LIST)))))

include scripts/makefiles/third_party/pasdam/makefiles/docker.mk
include scripts/makefiles/third_party/pasdam/makefiles/go.mk
include scripts/makefiles/third_party/pasdam/makefiles/go.mod.mk
include scripts/makefiles/third_party/pasdam/makefiles/help.mk

.DEFAULT_GOAL := help

GO_MAIN_DIR := ./cmd/go-scaffold

## build: Build all artifacts (binary and docker image)
.PHONY: build
build: | go-build docker-build

## clean: Remove all artifacts (binary and docker image)
.PHONY: clean
clean: | go-clean docker-clean

## install: Install all artifacts
.PHONY: install
install: | go-install
