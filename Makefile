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

.PHONY: clean
clean: kind-clean ## Install kind
