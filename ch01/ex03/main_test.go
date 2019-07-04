package main

import (
	"strings"
	"testing"
)

func TestJoinString(t *testing.T) {
	testCases := []struct {
		arg  []string
		want string
	}{
		{
			arg:  []string{"hoge", "foo", "bar"},
			want: "hoge foo bar",
		},
		{
			arg:  []string{"one-string"},
			want: "one-string",
		},
	}

	for _, tc := range testCases {
		got := JoinString(tc.arg)
		if got != tc.want {
			t.Errorf("got %s, want %s\n", got, tc.want)
		}
	}
}

func BenchmarkJoinString(b *testing.B) {
	fixture := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		JoinString(fixture)
	}
}

func BenchmarkJoinStringStd(b *testing.B) {
	fixture := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		strings.Join(fixture, " ")
	}
}
