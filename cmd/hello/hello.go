package main

import (
	"github.com/miekg/ebpf"
)

func HelloWorld(bctx *ebpf.Context) int {
	ebpf.TracePrintk("Hello world!\n")
	return 0
}
