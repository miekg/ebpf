package ebpf

// Context is a eBPF context, every eBPF program will have a (pointer) to a context
// loaded in r1 on startup.
type Context struct{}
