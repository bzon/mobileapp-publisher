package cmd

import (
	"fmt"
	"runtime"
)

const version = "v0.1.6"

func getVersion() string {
	return fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
}
