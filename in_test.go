package array

import (
	"fmt"
	"testing"
)

type T struct {
	test1 string
	test2 int
}

func ExampleIn_int() {
	a := []int{1, 2, 3, 4}

	fmt.Println(In(2, a))
	// Output:
	// true
}

func TestIn_intequal(t *testing.T) {
	a := []int{1, 2, 3, 4}

	if !In(2, a) {
		t.Fatalf("Incorrect answer")
	}
}

func TestIn_intnotequal(t *testing.T) {
	a := []int{1, 2, 3, 4}

	if In(0, a) {
		t.Fatalf("Incorrect answer")
	}
}

func TestIn_stringequal(t *testing.T) {
	a := []string{"dam", "test", "memes"}

	if !In("test", a) {
		t.Fatalf("Incorrect answer")
	}
}

func TestIn_stringnotequal(t *testing.T) {
	a := []string{"test", "memes", "dam"}

	if In("wtf", a) {
		t.Fatalf("Incorrect answer")
	}
}
