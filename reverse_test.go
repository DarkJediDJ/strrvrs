package strrvrs

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testReverseCases := []struct {
		name string
		text string
		want string
	}{
		{"English reverse", "Hello, World.", ".dlroW ,olleH"},
		{"Russian reverse", "Привет, мир.", ".рим ,тевирП"},
		{"Chinese reverse", "你好, 世界。", "。界世 ,好你"},
		{"Emojis reverse", "\U0001f60a \U0001f914 \U0001f644 \U0001f925 \U0001f634 \U0001f635", "\U0001f635 \U0001f634 \U0001f925 \U0001f644 \U0001f914 \U0001f60a"},
		{"Emoji reverse", "\U0001f60a", "\U0001f60a"},
		{"Empty reverse", "", ""},
		{"Tabs and spaces reverse", "  		", "		  "},
		{"Long string reverse", "Ob0ZoEDYxoPEYLXtVXGCm6fldgpucVSJxjwP5ae3BWiRgVoH4B5s", "s5B4HoVgRiWB3ea5PwjxJSVcupgdlf6mCGXVtXLYEPoxYDEoZ0bO"},
	}
	for _, tc := range testReverseCases {
		t.Run(tc.name, func(t *testing.T) {
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
		{"English calculator", "Example", 7},
		{"Russian calculator", "Пример строки", 13},
		{"Chinese calculator", "你好, 世界。", 7},
		{"Emojis calculator", "\U0001f60a \U0001f616 \U0001f644 \U0001f925 \U0001f634 \U0001f635 \U0001f480", 13},
		{"Emoji calculator", "\U0001f60a", 1},
		{"Empty calculator", "", 0},
		{"Tabs and spaces calculator", "  		", 4},
		{"Long string calculator", "oy4iGAPpxkz1NfJd4wak3wAIZ4xfSUFOjlTpLoO1OEQd26vXaBCS23cSnCPPFSOBI0H44T", 70},
	}
	for _, tc := range testCalculatorCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Calculator(tc.text); got != tc.want {
				t.Errorf("got %d\n; want %d\n", got, tc.want)
			}
		})
	}
}
