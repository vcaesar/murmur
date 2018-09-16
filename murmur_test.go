package murmur

import (
	"testing"

	"github.com/vcaesar/tt"
)

func TestMurmur3(t *testing.T) {
	text1 := []byte("The quick brown fox jumps over the lazy dog")
	expectedHash1 := uint32(0x78e69e27)
	actualHash1 := Murmur3(text1)
	tt.Equal(t, expectedHash1, actualHash1)

	text2 := []byte("The quick brown fox jumps over the lazy cog")
	expectedHash2 := uint32(0xd5ece287)
	actualHash2 := Murmur3(text2, 1)
	tt.Equal(t, expectedHash2, actualHash2)
}

func TestSum32(t *testing.T) {
	text1 := "The quick brown fox jumps over the lazy dog"
	expectedHash1 := uint32(0x78e69e27)
	actualHash1 := Sum32(text1)
	tt.Equal(t, expectedHash1, actualHash1)

	text2 := "The quick brown fox jumps over the lazy cog"
	expectedHash2 := uint32(0xd5ece287)
	actualHash2 := Sum32(text2, 1)
	tt.Equal(t, expectedHash2, actualHash2)
}
