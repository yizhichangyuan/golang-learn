package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}

	for _, test := range tests {
		out = new(bytes.Buffer)
		desc := fmt.Sprintf("echo(%v, %q, %q)",
			test.newline, test.sep, test.args)

		if err := echo(test.newline, test.args, test.sep); err != nil {
			t.Errorf("%s failed: %v", desc, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", desc, got, test.want)
		}
	}
}
