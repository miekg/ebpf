package gobpf

// Needs: internal/sys from bpf-go
// gopbf.BPF_F_INDEX_MASK

func BPF() {
	// context
	// gobpf.PerfEventOutput(bpf.Context, events.FD(), bpf.BPF_F_INDEX_MASK, 123 /* value */, 4 /* size*/)
	gobpf.PerfEventOutput(bpf.Context, events.FD(), gobpf.BPF_F_INDEX_MASK, 123)
}
