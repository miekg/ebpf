return0.o:
	clang -target bpf -Wall -O2 -c return0.c -o return0.o

return0_obj.o:
	clang -O2 -emit-llvm -c return0.c -o - | llc -march=bpf -mcpu=probe -filetype=obj -o return0.o

return0.s: return0.c
	clang -target bpf -S -o return0.s return0.c

dump: return0_obj.o
	 readelf -x .text return0.o

objdupmp:
	llvm-objdump -d return0.o

# go build -gcflags=-S return0.go
#
ssadump:
	ssadump -build=F return0.go
