package murmur

import (
	"testing"

	"github.com/vcaesar/tt"
)

var (
	testText = "The quick brown fox jumps over the lazy dog"
	text2    = "The quick brown fox jumps over the lazy cog"
)

func Benchmark_Sum32(b *testing.B) {
	fn := func() {
		Sum32(testText)
	}

	tt.BM(b, fn)

}

func Benchmark_Murmur3(b *testing.B) {
	fn := func() {
		Murmur3([]byte(testText))
	}

	tt.BM(b, fn)
}

func TestMurmur3(t *testing.T) {
	text1 := []byte(testText)
	expectedHash1 := uint32(0x78e69e27)
	actualHash1 := Murmur3(text1)
	tt.Equal(t, expectedHash1, actualHash1)

	expectedHash2 := uint32(0xd5ece287)
	actualHash2 := Murmur3([]byte(text2), 1)
	tt.Equal(t, expectedHash2, actualHash2)
}

func TestSum32(t *testing.T) {
	expectedHash1 := uint32(0x78e69e27)
	actualHash1 := Sum32(testText)
	tt.Equal(t, expectedHash1, actualHash1)

	expectedHash2 := uint32(0xd5ece287)
	actualHash2 := Sum32(text2, 1)
	tt.Equal(t, expectedHash2, actualHash2)
}
