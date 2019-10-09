TEXT ·asm_neg(SB), $0-8
	MOVQ 	a+0(FP), AX
	NEGQ 	AX
	MOVQ 	AX, ret+8(FP)
	RET
