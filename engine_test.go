package main

import (
	"testing"
)

func TestFog(t *testing.T) {
	if !(FOG == "?") {
		t.Fatalf(`FOG fail`)
	}
}

func TestSplitShortAlgebraic(t *testing.T) {
	x, y := SplitShortAlgebraic("a1")
	if x != "a" || y != "1" {
		t.Fatalf(`SplitShortAlgebraic a1 failed, was %s, %s`, x, y)
	}
}

func TestReplaceCharAt(t *testing.T) {
	s := ReplaceCharAt("abcd", 2, "z")
	if s != "abzd" {
		t.Fatalf(`ReplaceCharAt failed, was %s`, s)
	}
}

func TestToShortAlgebraic(t *testing.T) {
	s := ToShortAlgebraic([]string{"a", "b"}, []string{"1", "2"}, 2)
	if s != "b1" {
		t.Fatalf(`ToShortAlgebraic failed, was %s`, s)
	}
}
