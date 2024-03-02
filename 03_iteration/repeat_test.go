package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b")
	}
}

// Go test -bench ./repeat_test.go
func ExampleRepeat() {
	result := Repeat("a")
	fmt.Println(result)
	// Output aaaaa
}
