package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wabarc/memento"
)

func main() {
	var (
		version bool
	)

	const versionHelp = "Show version"

	flag.BoolVar(&version, "version", false, versionHelp)
	flag.BoolVar(&version, "v", false, versionHelp)
	flag.Parse()

	if version {
		fmt.Println(memento.Version)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		e := os.Args[0]
		fmt.Printf("  %s url [url]\n\n", e)
		fmt.Printf("example:\n  %s https://example.com https://example.org\n\n", e)
		os.Exit(1)
	}

	mem := &memento.Memento{}
	archives, _ := mem.Mementos(args)
	for orig, dest := range archives {
		fmt.Println(orig, "=>", dest)
	}
}
