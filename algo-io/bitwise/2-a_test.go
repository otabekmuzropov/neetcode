package bitwise_test

import (
	"neetcode/algo-io/bitwise"
	"testing"
)

func TestSetThirdPositionTo1(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Zero input",
			input:    0, // Binary: 0000
			expected: 2, // Binary: 0010
		},
		{
			name:     "Third bit already 1",
			input:    2, // Binary: 0010
			expected: 2, // Binary: 0010
		},
		{
			name:     "All bits 1",
			input:    255, // Binary: 11111111
			expected: 255, // Binary: 11111111
		},
		{
			name:     "Negative number",
			input:    -4, // Binary: ...1100
			expected: -2, // Binary: ...1110
		},
		{
			name:     "Other bits set",
			input:    5, // Binary: 0101
			expected: 7, // Binary: 0111
		},
		{
			name:     "Large number",
			input:    1024, // Binary: 10000000000
			expected: 1026, // Binary: 10000000010
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bitwise.SetPosition21(tt.input, 3)
			if result != tt.expected {
				t.Errorf("SetThirdPositionTo1(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
