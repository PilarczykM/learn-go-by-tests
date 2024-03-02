package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectAnswer(got, want, t)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectAnswer(want, got, t)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"

		assertCorrectAnswer(got, want, t)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"

		assertCorrectAnswer(got, want, t)
	})
}

func ExampleHello() {
	output := Hello("Chris", "Spanish")
	fmt.Println(output)
	// Output: Hola, Chris
}

func assertCorrectAnswer(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
