package murmur

import (
	"testing"

	"github.com/vcaesar/tt"
)

var (
	testText = "The quick brown fox jumps over the lazy dog"
	text2    = "The fox jumps cog"
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
	tt.Equal(t, uint32(0x78e69e27), Murmur3([]byte(testText)))
	tt.Equal(t, 3709117227, Murmur3([]byte(text2), 1))
}

func TestSum32(t *testing.T) {
	tt.Equal(t, uint32(0x78e69e27), Sum32(testText))
	tt.Equal(t, 3709117227, Sum32(text2, 1))
}
