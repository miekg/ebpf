package main

import "github.com/miekg/ebpf"

// copied from https://eunomia.dev/tutorials/4-opensnoop/#capturing-the-system-call-collection-of-process-opening-files-in-ebpf

func SySEnterWriteHandle(ctx *ebpf.Context) int {
	pid := ebpf.CurrentPidTgid() >> 32

	ebpf.Printk("BPF triggered sys_enter_write from PID %d.\n", pid)

	return 0
}

func main() {
	m := ebpf.Meta{
		License:     "GPL",
		Hook:        "tracepoint/syscalls/sys_enter_write",
		HandlerFunc: SySEnterWriteHandle,
	}

	m.Handle()
}
