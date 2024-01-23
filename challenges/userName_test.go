package main

import "testing"

func TestAValidUserName(t *testing.T) {
	want := CodelandUsernameValidation("fabricio_alessio")
	if want != "true" {
		t.Fatal("fabricio_alessio is a valid username")
	}
}

func TestAInvalidUserName(t *testing.T) {
	want := CodelandUsernameValidation("fabricio_alessio_")
	if want != "false" {
		t.Fatal("fabricio_alessio_ is a valid username")
	}
}
