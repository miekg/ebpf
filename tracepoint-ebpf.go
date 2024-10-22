// +go:build ignore
package main

func bpf() {
	// context
	bpf.PerfEventOutput(bpf.Context, events.FD(), 0xffffffff /*flags*/, 123 /* value */, 4 /* size*/)
}
