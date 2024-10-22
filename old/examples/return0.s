	.text
	.file	"return0.c"
	.globl	func                            # -- Begin function func
	.p2align	3
	.type	func,@function
func:                                   # @func
# %bb.0:
	r0 = 0
	exit
.Lfunc_end0:
	.size	func, .Lfunc_end0-func
                                        # -- End function
	.addrsig
