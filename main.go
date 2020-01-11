package main

import (
	"fmt"
	"github.com/kasika-technologies/gpx2linestring/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
