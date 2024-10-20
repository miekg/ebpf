# eBPF in *pure* Go

> eBPF in C? What am I, a farmer?

The goal here is: have a normal Go environment with a bunch a packages that compiles into a elf
binary that can be loaded into the kernel with `bpftool`.

This Go program can _also_ be compiled to a native Go (elf) binary, i.e. with `go build`, but that
problably doesn't do much.

The benefit of this approach is that you can use _all_ the Go development tooling for writing a eBPF
program.

How this will actually look in practice is uncertain, ideally eBPF should be "just" a compiler
backend for the Go compiler. Don't know how feasible that is, given the limitation of eBPF.

## Hello World

An [eBPF hello
world](https://github.com/eunomia-bpf/bpf-developer-tutorial/blob/main/src/1-helloworld/README_en.md#hello-world---minimal-ebpf-program),
look like this:

~~~ c
/* SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause) */
#define BPF_NO_GLOBAL_DATA
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

typedef unsigned int u32;
typedef int pid_t;
const pid_t pid_filter = 0;

char LICENSE[] SEC("license") = "Dual BSD/GPL";

SEC("tp/syscalls/sys_enter_write")
int handle_tp(void *ctx)
{
 pid_t pid = bpf_get_current_pid_tgid() >> 32;
 if (pid_filter && pid != pid_filter)
  return 0;
 bpf_printk("BPF triggered sys_enter_write from PID %d.\n", pid);
 return 0;
}
~~~

Would become something like

~~~ go
import "github.com/miekg/ebpf"

func Syscalls_SysEnterWrite(ctx ebpf.Context) int {
    pid := ebpf.GetCurrentPidTgid() >> 32
    if pid != 0 {
        epbf.Printk("BPF triggered sys_enter_write from PID %d.\n", pid)
    }
    return 0
}
~~~

## TinyGo

Maybe this should be a backend for tinygo, that still uses llvm to generate binary, but that may
actually be a good thing, as llvm is the official(?) ebpf compiler. This [was even suggested a
while](https://github.com/tinygo-org/tinygo/issues/1015) a go.

## Stuff of interest

* go-delve/delve@v1.22.1/pkg/proc/internal/ebpf/helpers.go
* github.com:miekg/ebpfcat (forked)
* https://github.com/DQNEO/babygo

## Reading list

- https://qmonnet.github.io/whirl-offload/2020/04/12/llvm-ebpf-asm/
- https://pkg.go.dev/golang.org/x/tools/cmd/ssadump
- https://go.googlesource.com/tools/+/master/go/ssa/example_test.go
