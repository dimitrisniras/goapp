package util

import (
	"regexp"
	"testing"
)

func TestGenerateHexString(t *testing.T) {
	// Test string length
	hexStr := RandHexString(16)
	if len(hexStr) != 16 {
		t.Errorf("Expected length 16, got %d", len(hexStr))
	}

	// Test hex format
	match, _ := regexp.MatchString("^[0-9A-F]+$", hexStr)
	if !match {
		t.Errorf("Expected hex string, got %s", hexStr)
	}
}

func BenchmarkGenerateHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandHexString(16)
	}
}
