package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

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
	var wg sync.WaitGroup
	for _, arg := range args {
		wg.Add(1)
		go func(arg string) {
			defer wg.Done()
			in, err := url.Parse(arg)
			if err != nil {
				fmt.Println(arg, "=>", fmt.Sprint(err))
				return
			}

			dst, err := mem.Mementos(context.TODO(), in)
			if err != nil {
				fmt.Println(arg, "=>", fmt.Sprint(err))
				return
			}
			fmt.Println(arg, "=>", dst)
		}(arg)
		// nice for memento service
		time.Sleep(time.Second)
	}
	wg.Wait()
}
