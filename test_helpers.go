package main

import "testing"

func assertBoolEqual(t *testing.T, a bool, b bool) {
	if a != b {
		t.Fatalf("Boolean is not equal")
	}
}

func assertIntEqual(t *testing.T, a int, b int) {
	if a != b {
		t.Fatalf("Integers is not equal %d != %d", a, b)
	}
}

func assertNodeEqual(t *testing.T, a *Node, b *Node) {
	if a != b {
		t.Fatalf("Node is not equal %d != %d", a, b)
	}
}