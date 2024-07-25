# ebpf

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
