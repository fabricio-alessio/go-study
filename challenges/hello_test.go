package main

import "testing"

func TestHello(t *testing.T) {
	want := hello("size")
	if want != 4 {
		t.Fatal("Word size must have size 4")
	}
}
