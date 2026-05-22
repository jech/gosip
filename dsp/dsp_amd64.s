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

#define ULAW_BIAS $0x84

// func L16MixSat160(dst, src *int16)
// Explanation: http://i.imgur.com/nejgQ41.jpg
TEXT    ·L16MixSat160(SB),4,$0-16
	MOVQ	dst+0(FP), AX
	MOVQ	src+8(FP), BX
	MOVQ	$19, CX
moar:	MOVO	0(BX), X0
	PADDSW  0(AX), X0
	MOVO	X0, 0(AX)
	ADDQ	$16, AX
	ADDQ	$16, BX
	DECQ	CX
	CMPQ	CX, $0
	JGE	moar
	RET
