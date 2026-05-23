#include "textflag.h"

// func sadd(dst, src *int16, n int)
TEXT ·sadd(SB),NOSPLIT,$0
	MOVD	dst+0(FP), R0
	MOVD	src+8(FP), R1
	MOVD	n+16(FP), R2

more:
	CMP	$8, R2
	BLT	tail
	VLD1	(R1), [V0.H8]
	VLD1	(R0), [V1.H8]
	WORD	$0x4E600C21  // SQADD V0.H8, V1.H8, V1.H8
	VST1	[V1.H8], (R0)
	ADD	$16, R0
	ADD	$16, R1
	SUB	$8, R2
	B	more

tail:
	MOVD	$32767, R3
	MOVD	$-32768, R4

evenmore:
	CBZ	R2, done
	MOVH	(R0), R5
	MOVH	(R1), R6
	ADD	R6, R5
	CMP	R3, R5
	CSEL	GT, R3, R5, R5
	CMP	R4, R5
	CSEL	LT, R4, R5, R5
	MOVH	R5, (R0)
	ADD	$2, R0
	ADD	$2, R1
	SUB	$1, R2
	B	evenmore

done:
	RET
