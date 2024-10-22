package ebpf

// Context is the ebpf context that is used as the first paramaters in every eBPF function that
// gets called from kernel context. TODO(miek): no idea if true
type Context struct{}

type HandlerFunc func(ctx *Context) int

// Meta contains meta data for an eBPF function, the varias fields have their own documentation.
type Meta struct {
	// License is the license. TODO(miek): make it a type?
	License string

	// Hook .... also make this a type, so ebpf/stracepoint.Syscall.Sys_enter_write is a thing?
	Hook string

	HandlerFunc
}

func (m Meta) Register() int {
	// do something, or nothing
	return 0
}
