package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Web II")
	want := "Olá, Web II"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
