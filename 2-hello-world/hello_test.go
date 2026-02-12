package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Olá, mundo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
