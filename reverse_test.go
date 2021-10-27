package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		text string
		want string
	}{
		{"Hello, World.", ".dlroW ,olleH"},
		{"Привет, мир.", ".рим ,тевирП"},
		{"你好, 世界。", "。界世 ,好你"},
		{"\U0001f60a", "\U0001f60a"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("New string is %s", tc.text), func(t *testing.T) {
			text := Reverse(tc.text)
			if got := text; got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}
