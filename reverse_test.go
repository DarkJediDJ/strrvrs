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
