// Code generated by github.com/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
	_ "github.com/google/addlicense"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go run internal/cmd/convert-swagger/main.go
//go:generate go build -v -o=./bin/addlicense github.com/google/addlicense
