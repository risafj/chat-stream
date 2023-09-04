package main

import "flag"

// flags
// if --block, use block responses
var blockFlag = flag.Bool("block", false, "Use non-streaming block responses")

func main() {
	flag.Parse()
	if *blockFlag {
		block()
	} else {
		stream()
	}
}
