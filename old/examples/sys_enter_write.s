	.text
	.file	"sys_enter_write.c"
	.section	"tp/syscalls/sys_enter_write","ax",@progbits
	.globl	handle_tp                       # -- Begin function handle_tp
	.p2align	3
	.type	handle_tp,@function
handle_tp:                              # @handle_tp
# %bb.0:
	*(u64 *)(r10 - 8) = r1
	r1 = bpf_get_current_pid_tgid ll
	r1 = *(u64 *)(r1 + 0)
	callx r1
	r0 >>= 32
	*(u32 *)(r10 - 12) = r0
	r1 = bpf_trace_printk ll
	r4 = *(u64 *)(r1 + 0)
	r3 = *(u32 *)(r10 - 12)
	r1 = handle_tp.____fmt ll
	r2 = 9
	callx r4
	*(u64 *)(r10 - 24) = r0
	r0 = 0
	exit
.Lfunc_end0:
	.size	handle_tp, .Lfunc_end0-handle_tp
                                        # -- End function
	.type	LICENSE,@object                 # @LICENSE
	.section	license,"aw",@progbits
	.globl	LICENSE
LICENSE:
	.asciz	"Dual BSD/GPL"
	.size	LICENSE, 13

	.type	bpf_get_current_pid_tgid,@object # @bpf_get_current_pid_tgid
	.data
	.p2align	3, 0x0
bpf_get_current_pid_tgid:
	.quad	14
	.size	bpf_get_current_pid_tgid, 8

	.type	handle_tp.____fmt,@object       # @handle_tp.____fmt
	.section	.rodata,"a",@progbits
handle_tp.____fmt:
	.asciz	"AAA %d.\n"
	.size	handle_tp.____fmt, 9

	.type	bpf_trace_printk,@object        # @bpf_trace_printk
	.data
	.p2align	3, 0x0
bpf_trace_printk:
	.quad	6
	.size	bpf_trace_printk, 8

	.addrsig
	.addrsig_sym handle_tp
	.addrsig_sym LICENSE
	.addrsig_sym bpf_get_current_pid_tgid
	.addrsig_sym handle_tp.____fmt
	.addrsig_sym bpf_trace_printk
