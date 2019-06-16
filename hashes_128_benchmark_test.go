package go_checksum_benchmarks

import "testing"

/*
We measure the Initialization, Write and Sum of algorithms.
*/

func Benchmark_text1_FNV1_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewFNV1_128()
		_, err := h.Write(text1)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text1_FNV1_128a(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewFNV1_128a()
		_, err := h.Write(text1)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text1_murmur3_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewMurmur3_128()
		_, err := h.Write(text1)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text1_jenkins_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewJenkins_128()
		_, err := h.Write(text1)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text2_FNV1_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewFNV1_128()
		_, err := h.Write(text2)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text2_FNV1_128a(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewFNV1_128a()
		_, err := h.Write(text2)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text2_murmur3_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewMurmur3_128()
		_, err := h.Write(text2)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text2_jenkins_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewJenkins_128()
		_, err := h.Write(text2)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text3_FNV1_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewFNV1_128()
		_, err := h.Write(text3)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text3_FNV1_128a(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewFNV1_128a()
		_, err := h.Write(text3)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text3_murmur3_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewMurmur3_128()
		_, err := h.Write(text3)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}

func Benchmark_text3_jenkins_128(b *testing.B) {
	var h Hash128
	for n := 0; n < b.N; n++ {
		h = NewJenkins_128()
		_, err := h.Write(text3)
		if err != nil {
			b.Fatal(err)
		}
		h.Sum128Bytes()
	}
}
