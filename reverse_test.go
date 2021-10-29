package strrvrs

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testReverseCases := []struct {
		name string
		text string
		want string
	}{
		{"Simple string", "Hello, World.", ".dlroW ,olleH"},
		{"Russian string", "Привет, мир.", ".рим ,тевирП"},
		{"Chinese string", "你好, 世界。", "。界世 ,好你"},
		{"Lots of emojis", "\U0001f60a \U0001f914 \U0001f644 \U0001f925 \U0001f634 \U0001f635", "\U0001f635 \U0001f634 \U0001f925 \U0001f644 \U0001f914 \U0001f60a"},
		{"One emoji", "\U0001f60a", "\U0001f60a"},
		{"Empty string", "", ""},
		{"String with tabs and spaces", "  		", "		  "},
		{"Very long string", "Ob0ZoEDYxoPEYLXtVXGCm6fldgpucVSJxjwP5ae3BWiRgVoH4B5s", "s5B4HoVgRiWB3ea5PwjxJSVcupgdlf6mCGXVtXLYEPoxYDEoZ0bO"},
	}
	for _, tc := range testReverseCases {
		t.Run(fmt.Sprintf("Couldnt reverse string: %s in case : %s", tc.text, tc.name), func(t *testing.T) {
			if got := Reverse(tc.text); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestCalculator(t *testing.T) {
	testCalculatorCases := []struct {
		name string
		text string
		want int
	}{
		{"Simple string", "Example", 7},
		{"Russian string", "Россия!!!", 9},
		{"Chinese string", "你好, 世界。", 7},
		{"Lots of emojis", "\U0001f60a \U0001f616 \U0001f644 \U0001f925 \U0001f634 \U0001f635 \U0001f480", 13},
		{"One emoji", "\U0001f60a", 1},
		{"Empty string", "", 0},
		{"String with tabs and spaces", "  		", 4},
		{"Very long string", "oy4iGAPpxkz1NfJd4wak3wAIZ4xfSUFOjlTpLoO1OEQd26vXaBCS23cSnCPPFSOBI0H44T", 70},
	}
	for _, tc := range testCalculatorCases {
		t.Run(fmt.Sprintf("Couldnt count chars in string: %s in case : %s", tc.text, tc.name), func(t *testing.T) {
			if got := Calculator(tc.text); got != tc.want {
				t.Errorf("got %d\n; want %d\n", got, tc.want)
			}
		})
	}
}
