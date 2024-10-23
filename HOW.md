This document is a (running) text on how the compiler is implemented, and why certain choices have
been made.

## RO Data

Read only data, like literal strings, etc., `ast.BasicLit` in Go, used to be stored in the 512B eBPF
stack, but as that space is limited, it is now stored in a map. This map is named after the Go file
name with `.rodata` appended.

This map is filled as-we-go, and duplicate literals are _not_ detected. Short literal are just used
with an immediate load.
