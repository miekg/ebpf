package main

import "github.com/miekg/ebpf/examples/pkg/plus"

func main() {
	c := plus.Add(5, 6)

	c = c
}
