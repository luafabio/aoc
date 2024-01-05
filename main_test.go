package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	filePath := "C:\\Users\\Luis\\go\\aoc\\aoc.txt"

	b, _ := os.ReadFile(filePath)

	tests := []struct {
		expected int
		input    []byte
	}{
		{
			expected: 281,
			input:    b,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, partone(b))
	}
}
