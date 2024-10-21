/*
Package prog provides a dummy eBPF program that can be attached to perf events.
*/
package prog

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
)

// GetEbpfProg returns an eBPF program specification for a perf event that prints pid and tgid.
func GetEbpfProg() *ebpf.ProgramSpec {
	var progSpec = &ebpf.ProgramSpec{
		Name:    "exampleProg",
		Type:    ebpf.PerfEvent,
		License: "Dual MIT/GPL",
	}

	progSpec.Instructions = asm.Instructions{
		// fmt[] = "pid %d  tgid %d"
		asm.LoadImm(asm.R1, 28188318724876148, asm.DWord),
		asm.StoreMem(asm.RFP, -8, asm.R1, asm.DWord),
		asm.LoadImm(asm.R1, 2314960319088454000, asm.DWord),
		asm.StoreMem(asm.RFP, -16, asm.R1, asm.DWord),

		// id = bpf_get_current_pid_tgid();
		asm.FnGetCurrentPidTgid.Call(),
		asm.Mov.Reg(asm.R3, asm.R0),

		// pid = id >> 32;
		asm.RSh.Imm(asm.R3, 32),

		// tgid = id & 0xFFFF
		asm.And.Imm(asm.R0, 0xFFFF),

		// bpf_trace_printk(fmt, sizeof(fmt), pid, tgid)
		asm.Mov.Reg(asm.R1, asm.RFP),
		asm.Add.Imm(asm.R1, -16),
		asm.Mov.Imm(asm.R2, 16),
		asm.Mov.Reg(asm.R4, asm.R0),
		asm.FnTracePrintk.Call(),

		// return 0
		asm.Mov.Imm(asm.R0, 0),
		asm.Return(),
	}
	return progSpec
}
