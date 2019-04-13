// +build !linux

package main

import (
	"fmt"
	"runtime"
)

func setSubreaper() error {
	return fmt.Errorf("-r not supported on %s-%s", runtime.GOOS, runtime.GOARCH)
}
