// Copyright 2020 Justine Alexandra Roberts Tunney
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "textflag.h"

// func sadd(dst, src *int16, n int)
TEXT	·sadd(SB),NOSPLIT,$0
	MOVQ	dst+0(FP), AX
	MOVQ	src+8(FP), BX
	MOVQ	n+16(FP), CX
more:
	CMPQ	CX, $8
	JL	tail
	MOVOU	0(BX), X0
	MOVOU   0(AX), X1
	PADDSW	X1, X0
	MOVOU	X0, 0(AX)
	ADDQ	$16, AX
	ADDQ	$16, BX
	SUBQ	$8, CX
	JMP	more

tail:
	MOVL	$32767, R8
	MOVL	$-32768, R9

evenmore:
	TESTQ	CX, CX
	JZ	done
	MOVWLSX	0(AX), DX
	MOVWLSX	0(BX), SI
	ADDL	SI, DX
	CMPL	DX, $32767
	CMOVLGE	R8, DX
	CMPL	DX, $-32768
	CMOVLLT	R9, DX
	MOVW	DX, 0(AX)
	ADDQ	$2, AX
	ADDQ	$2, BX
	DECQ	CX
	JMP	evenmore

done:           
	RET
