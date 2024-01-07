package main

import (
	"flag"
	"fmt"
	"github.com/jyuart/ccb/internal"
)

func main() {
	isPaste := flag.Bool("p", false, "Set -p to paste from a remote clipboard; copies to it by default")
	flag.Parse()

	if *isPaste {
		fmt.Println("Not implemented")
	} else {
		internal.SetCoolClipboard()
	}
}
