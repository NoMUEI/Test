package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got'%s' want '%s'",got, want)
		}
	}

	t.Run("saying hello to people" ,func(t *testing.T){
		got := Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
		if got != want {
			t.Errorf("got'%s' want '%s'",got, want)
		}
	})

	t.Run("say hello world when an empty",func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		if got != want {
			t.Errorf("got'%s' want '%s'",got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie","Spanish")
		want := "Hello, Elodie"
		assertCorrectMessage(t, got, want)
	})




}
