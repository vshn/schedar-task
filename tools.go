//go:build tools
// +build tools

package tools

// Package tools is a place to put any tooling dependencies as imports.
// Go modules will be forced to download and install them.
import (
	// controller-gen
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
