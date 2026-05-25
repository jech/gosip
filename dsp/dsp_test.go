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

package dsp_test

import (
	"fmt"
	"testing"

	"github.com/jart/gosip/dsp"
)

func testSadd(t *testing.T, n int) {
	x := make([]int16, n)
	y := make([]int16, n)
	for i := 0; i < n; i++ {
		x[i] = int16(i)
		y[i] = int16(666)
	}
	dsp.Sadd(x, y)
	for i := 0; i < n; i++ {
		want := int16(i + 666)
		if x[i] != want {
			t.Errorf("x[%v] = %v (wanted: %v)", i, x[i], want)
			return
		}
		if y[i] != int16(666) {
			t.Errorf("side effect y[%v] = %v", i, y[i])
			return
		}
	}
}

func TestSadd(t *testing.T) {
	for n := range []int{
		0, 3, 7, 8, 16, 32, 33,
		160, 161, 162, 163, 164, 165, 166, 167, 168,
	} {
		t.Run(fmt.Sprintf("%d", n), func(t *testing.T) {
			testSadd(t, n)
		})
	}
}

func TestSaddUnaligned(t *testing.T) {
	x := make([]int16, 161)
	y := make([]int16, 161)
	for i := 0; i < 161; i++ {
		y[i] = int16(i)
	}
	dsp.Sadd(x[1:], y[1:])
	if x[0] != 0 || x[1] != 1 || x[160] != 160 {
		t.Errorf("omg")
	}
	dsp.Sadd(x[1:], y)
	if x[0] != 0 || x[1] != 1 || x[160] != 160+159 {
		t.Errorf("omg")
	}
	dsp.Sadd(x, y[1:])
	if x[0] != 1 || x[1] != 3 || x[160] != 160+159 {
		t.Errorf("omg %v", x[160])
	}
}

func BenchmarkSadd(b *testing.B) {
	x := make([]int16, 160)
	y := make([]int16, 160)
	for i := 0; i < 160; i++ {
		x[i] = int16(i)
		y[i] = int16(666)
	}

	b.ResetTimer()
	b.SetBytes(160 * 2)

	for i := 0; i < b.N; i++ {
		dsp.Sadd(x, y)
	}
	if b.N > 100 && (x[0] != 0x7FFF || x[159] != 0x7FFF) {
		b.Errorf("omg")
	}
}

var ulawTable = []int16{
	-32124, -31100, -30076, -29052, -28028, -27004, -25980, -24956, -23932,
	-22908, -21884, -20860, -19836, -18812, -17788, -16764, -15996, -15484,
	-14972, -14460, -13948, -13436, -12924, -12412, -11900, -11388, -10876,
	-10364, -9852, -9340, -8828, -8316, -7932, -7676, -7420, -7164,
	-6908, -6652, -6396, -6140, -5884, -5628, -5372, -5116, -4860,
	-4604, -4348, -4092, -3900, -3772, -3644, -3516, -3388, -3260,
	-3132, -3004, -2876, -2748, -2620, -2492, -2364, -2236, -2108,
	-1980, -1884, -1820, -1756, -1692, -1628, -1564, -1500, -1436,
	-1372, -1308, -1244, -1180, -1116, -1052, -988, -924, -876,
	-844, -812, -780, -748, -716, -684, -652, -620, -588,
	-556, -524, -492, -460, -428, -396, -372, -356, -340,
	-324, -308, -292, -276, -260, -244, -228, -212, -196,
	-180, -164, -148, -132, -120, -112, -104, -96, -88,
	-80, -72, -64, -56, -48, -40, -32, -24, -16,
	-8, 0, 32124, 31100, 30076, 29052, 28028, 27004, 25980,
	24956, 23932, 22908, 21884, 20860, 19836, 18812, 17788, 16764,
	15996, 15484, 14972, 14460, 13948, 13436, 12924, 12412, 11900,
	11388, 10876, 10364, 9852, 9340, 8828, 8316, 7932, 7676,
	7420, 7164, 6908, 6652, 6396, 6140, 5884, 5628, 5372,
	5116, 4860, 4604, 4348, 4092, 3900, 3772, 3644, 3516,
	3388, 3260, 3132, 3004, 2876, 2748, 2620, 2492, 2364,
	2236, 2108, 1980, 1884, 1820, 1756, 1692, 1628, 1564,
	1500, 1436, 1372, 1308, 1244, 1180, 1116, 1052, 988,
	924, 876, 844, 812, 780, 748, 716, 684, 652,
	620, 588, 556, 524, 492, 460, 428, 396, 372,
	356, 340, 324, 308, 292, 276, 260, 244, 228,
	212, 196, 180, 164, 148, 132, 120, 112, 104,
	96, 88, 80, 72, 64, 56, 48, 40, 32, 24, 16, 8, 0,
}

var alawTable = []int16{
	-5504, -5248, -6016, -5760, -4480, -4224, -4992, -4736,
	-7552, -7296, -8064, -7808, -6528, -6272, -7040, -6784,
	-2752, -2624, -3008, -2880, -2240, -2112, -2496, -2368,
	-3776, -3648, -4032, -3904, -3264, -3136, -3520, -3392,
	-22016, -20992, -24064, -23040, -17920, -16896, -19968, -18944,
	-30208, -29184, -32256, -31232, -26112, -25088, -28160, -27136,
	-11008, -10496, -12032, -11520, -8960, -8448, -9984, -9472,
	-15104, -14592, -16128, -15616, -13056, -12544, -14080, -13568,
	-344, -328, -376, -360, -280, -264, -312, -296,
	-472, -456, -504, -488, -408, -392, -440, -424,
	-88, -72, -120, -104, -24, -8, -56, -40,
	-216, -200, -248, -232, -152, -136, -184, -168,
	-1376, -1312, -1504, -1440, -1120, -1056, -1248, -1184,
	-1888, -1824, -2016, -1952, -1632, -1568, -1760, -1696,
	-688, -656, -752, -720, -560, -528, -624, -592,
	-944, -912, -1008, -976, -816, -784, -880, -848,
	5504, 5248, 6016, 5760, 4480, 4224, 4992, 4736,
	7552, 7296, 8064, 7808, 6528, 6272, 7040, 6784,
	2752, 2624, 3008, 2880, 2240, 2112, 2496, 2368,
	3776, 3648, 4032, 3904, 3264, 3136, 3520, 3392,
	22016, 20992, 24064, 23040, 17920, 16896, 19968, 18944,
	30208, 29184, 32256, 31232, 26112, 25088, 28160, 27136,
	11008, 10496, 12032, 11520, 8960, 8448, 9984, 9472,
	15104, 14592, 16128, 15616, 13056, 12544, 14080, 13568,
	344, 328, 376, 360, 280, 264, 312, 296,
	472, 456, 504, 488, 408, 392, 440, 424,
	88, 72, 120, 104, 24, 8, 56, 40,
	216, 200, 248, 232, 152, 136, 184, 168,
	1376, 1312, 1504, 1440, 1120, 1056, 1248, 1184,
	1888, 1824, 2016, 1952, 1632, 1568, 1760, 1696,
	688, 656, 752, 720, 560, 528, 624, 592,
	944, 912, 1008, 976, 816, 784, 880, 848,
}

func TestLinearToUlaw(t *testing.T) {
	if dsp.LinearToUlaw(0) != 255 {
		t.Error("omg")
	}
	if dsp.LinearToUlaw(-100) != 114 {
		t.Error("omg")
	}
}

func TestUlawToLinear(t *testing.T) {
	for n := 0; n <= 255; n++ {
		if dsp.UlawToLinear(byte(n)) != ulawTable[n] {
			t.Error("omg")
			return
		}
	}
}

func TestAlawToLinear(t *testing.T) {
	for n := 0; n <= 255; n++ {
		if dsp.AlawToLinear(byte(n)) != alawTable[n] {
			t.Error("omg")
			return
		}
	}
}

func TestLinearToUlawToLinear(t *testing.T) {
	for n := 0; n <= 255; n++ {
		if dsp.UlawToLinear(dsp.LinearToUlaw(ulawTable[n])) != ulawTable[n] {
			t.Error("omg")
			return
		}
	}
}

func TestLinearToAlawToLinear(t *testing.T) {
	for n := 0; n <= 255; n++ {
		if dsp.AlawToLinear(dsp.LinearToAlaw(alawTable[n])) != alawTable[n] {
			t.Error("omg")
			return
		}
	}
}

func BenchmarkUlawToLinear(b *testing.B) {
	u := make([]byte, 4096)
	l := make([]int16, 4096)
	for i := 0; i < 4096; i++ {
		u[i] = byte(i)
	}
	b.ResetTimer()
	b.SetBytes(4096)
	for i := 0; i <= b.N; i++ {
		for j := 0; j < 4096; j++ {
			l[j] = dsp.UlawToLinear(u[j])
		}
		if l[0]+l[4095] != -32124 {
			b.Errorf("omg")
		}
	}
}

func BenchmarkLinearToUlaw(b *testing.B) {
	l := make([]int16, 4096)
	u := make([]byte, 4096)
	for i := 0; i < 4096; i++ {
		l[i] = int16(i)
	}
	b.ResetTimer()
	b.SetBytes(2 * 4096)
	for i := 0; i <= b.N; i++ {
		for j := 0; j < 4096; j++ {
			u[j] = dsp.LinearToUlaw(l[j])
		}
		if u[0]+u[4095] != 174 {
			b.Errorf("omg")
		}
	}
}

func BenchmarkAlawToLinear(b *testing.B) {
	u := make([]byte, 4096)
	l := make([]int16, 4096)
	for i := 0; i < 4096; i++ {
		u[i] = byte(i)
	}
	b.ResetTimer()
	b.SetBytes(4096)
	for i := 0; i <= b.N; i++ {
		for j := 0; j < 4096; j++ {
			l[j] = dsp.AlawToLinear(u[j])
		}
		if l[0]+l[4095] != -4656 {
			b.Errorf("omg")
		}
	}
}

func BenchmarkLinearToAlaw(b *testing.B) {
	l := make([]int16, 4096)
	a := make([]byte, 4096)
	for i := 0; i < 4096; i++ {
		l[i] = int16(i)
	}
	b.ResetTimer()
	b.SetBytes(2 * 4096)
	for i := 0; i <= b.N; i++ {
		for j := 0; j < 4096; j++ {
			a[j] = dsp.LinearToAlaw(l[j])
		}
		if a[0]+a[4095] != 111 {
			b.Errorf("omg")
		}
	}
}
