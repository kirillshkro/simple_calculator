package main

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

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
			go func() {
				convey.Convey(tt.name, t, func() {
					result := make(chan int)
					defer close(result)
					go Calculate(tt.line, result)
					for res := range result {
						convey.So(tt.expected, convey.ShouldEqual, res)
					}
				})

			}()
		})
	}
}

func Test_trim(t *testing.T) {
	convey.Convey("Test_trim", t, func() {
		mockData := "I'm where, i'm there, i'm everyhere"
		expected := "I'mwhere,i'mthere,i'meveryhere"

		result := trim(mockData)
		convey.So(expected, convey.ShouldEqual, result)
	})

}

func Test_parse(t *testing.T) {
	convey.Convey("Test parse", t, func() {
		mockData := " 2 * 2"
		expectedToken := Token{
			operand:  "2",
			operand2: "2",
			opKind:   '*',
		}
		result, _ := parse(mockData)
		convey.So(expectedToken.opKind, convey.ShouldEqual, result.opKind)
		convey.So(expectedToken.operand, convey.ShouldEqual, result.operand)
		convey.So(expectedToken.operand2, convey.ShouldEqual, result.operand2)
	})
}
