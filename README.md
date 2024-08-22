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


== older stuff ==


In examples/ I'm trying to convert C ebpf code to non-working Go code to get a feel on how to the Go
API should work; completely uncertain if this is going to work.

## Requirements

* Normal Go code, it should compile
* Like working with any package in this case "ebpf"
* Seperate "compiler" that compiles to ebpf ELF
    - Look at avo?
* libbpf is defacto standard, that manages the loading, unloading, etc?
* compile ebpf helpers to assembly and use that? Just like the .S file from Go but
    then for the ebpf VM?


How is this all different than writing a new compiler backend for the Go compiler?

## TinyGo

Maybe this should be a backend for tinygo, that still uses llvm to generate binary, but that may
actually be a good thing, as llvm is the official(?) ebpf compiler. This [was even suggested a
while](https://github.com/tinygo-org/tinygo/issues/1015) a go.
