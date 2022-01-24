package prose_test

import (
	"fmt"
	"go-playground/examples/tests/prose"
	"testing"
)

func TestOneElement(t *testing.T) {
	list := []string{
		"apple",
	}
	want := "apple"
	got := prose.JoinWithCommas(list)
	if want != got {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{
		"apple",
		"orange",
	}
	want := "apple and orange"
	got := prose.JoinWithCommas(list)
	if want != got {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{
		"apple",
		"orange",
		"pear",
	}
	want := "apple, orange, and pear"
	got := prose.JoinWithCommas(list)
	if want != got {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\",want \"%s\"", list, got, want)
}

type testData struct {
	list []string
	want string
}

//表驱动测试,推荐
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{
			list: []string{},
			want: "",
		},
		{
			list: []string{"apple"},
			want: "apple",
		}, {
			list: []string{"apple", "orange"},
			want: "apple and orange",
		}, {
			list: []string{"apple", "orange", "pear"},
			want: "apple, orange, and pear",
		},
	}
	for _, test := range tests {
		got := prose.JoinWithCommas(test.list)
		want := test.want
		if got != want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\",want \"%s\"", test.list, got, want)
		}
	}
}
