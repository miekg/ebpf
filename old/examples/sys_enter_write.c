/* SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause) */
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

char LICENSE[] SEC("license") = "Dual BSD/GPL";

SEC("tp/syscalls/sys_enter_write")
int handle_tp(void *ctx)
{
	int pid = bpf_get_current_pid_tgid() >> 32;
	bpf_printk("PID %d.\n", pid);
	return 0;
}
