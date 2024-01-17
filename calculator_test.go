package main

import "testing"

func Test_Calculate(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		result   int
		expected int
	}{
		// TODO: Add test cases.
		{
			name:     "Add",
			line:     "2+2",
			result:   4,
			expected: 4,
		},
		{
			name:     "Substract",
			line:     "10 - 5",
			result:   5,
			expected: 5,
		},
		{
			name:     "Multiply",
			line:     "6*9",
			result:   54,
			expected: 54,
		},
		{
			name:     "Divide",
			line:     "9 / 3",
			result:   3,
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Calculate(tt.line, tt.result)
			if tt.expected != tt.result {
				t.Errorf("In %s %d != %d\n", tt.name, tt.expected, tt.result)
			}
		})
	}
}

func Test_trim(t *testing.T) {
	mockData := "I'm where, i'm there, i'm everyhere"
	expected := "I'mwhere,i'mthere,i'meveryhere"

	result := trim(mockData)
	if result != expected {
		t.Errorf("%s not compare with %s\n", result, expected)
	}
}

func Test_parse(t *testing.T) {
	mockData := " 2 * 2"
	expectedToken := Token{
		operand:  "2",
		operand2: "2",
		opKind:   '*',
	}
	result, _ := parse(mockData)

	if result.opKind != expectedToken.opKind || result.operand != expectedToken.operand || result.operand2 != expectedToken.operand2 {
		t.Error("Test fallen")
	}
}
