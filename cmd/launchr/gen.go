//go:build ignore

package main

import (
	"github.com/launchrctl/launchr"

	_ "github.com/launchrctl/update"
)

func main() {
	launchr.GenAndExit()
}
