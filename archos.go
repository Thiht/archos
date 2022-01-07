package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var format string

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "archos prints the current OS and architecture using the Golang namings.")
		fmt.Fprintln(os.Stderr, "It can be useful to download released Go binaries with OS and architecture present in the URL.")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
	}
	flag.StringVar(&format, "format", "{os}-{arch}", "Format of the output (allowed variables: {os}, {arch})")
	flag.Parse()

	operatingSystem := runtime.GOOS
	architecture := runtime.GOARCH

	replacer := strings.NewReplacer(
		"{os}", operatingSystem,
		"{arch}", architecture,
	)
	fmt.Println(replacer.Replace(format))
}
