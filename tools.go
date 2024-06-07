//go:build tools
// +build tools

package main

import (
	_ "github.com/codegangsta/gin"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/x-motemen/gore/cmd/gore"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/gopls"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
