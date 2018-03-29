package array

import (
	"fmt"
	"testing"
)

// Examples
func ExampleIntersection() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersectionsI, _ := Intersection(a, b)
	intersections := intersectionsI.([]int)
	fmt.Println(intersections)
	// Output:
	// [2 3]
}

func ExampleExcept_a() {
	a := []int{1, 2, 5}
	b := []int{2, 3, 4}

	exceptI, _ := Except(a, b)
	except := exceptI.([]int)
	fmt.Println(except)
	// Output:
	// [1 5]
}

func ExampleExcept_b() {
	a := []int{1, 2, 5}
	b := []int{2, 3, 4}

	exceptI, _ := Except(b, a)
	except := exceptI.([]int)
	fmt.Println(except)
	// Output:
	// [3 4]
}

func ExampleExcept_string() {
	a := []string{"who", "what", "when"}
	b := []string{"when", "where", "why"}

	exceptI, _ := Except(a, b)
	except := exceptI.([]string)
	fmt.Println(except)
	// Output:
	// [who what]
}

// Benchmark
var result []int

func BenchmarkIntersection(bench *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{12, 4, 3, 16, 100, 9, 1, 18, 124}
	var r []int
	for i := 0; i < bench.N; i++ {
		rI, _ := Intersection(a, b)
		r = rI.([]int)
	}
	result = r
}

// Tests
// Intersection
func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersectsI, err := Intersection(a, b)
	if err != nil {
		t.Fatal(err)
	}
	intersects := intersectsI.([]int)

	if len(intersects) != 2 || intersects[0] != 2 || intersects[1] != 3 {
		t.Fatal("Result array has incorrect terms")
	}
}

// Except
func TestExcept_int(t *testing.T) {
	a := []int{1, 2, 5}
	b := []int{2, 3, 4}

	exceptI, err := Except(a, b)
	if err != nil {
		t.Fatal(err)
	}
	except := exceptI.([]int)

	if len(except) != 2 {
		t.Fatal("Two items should be returned")
	}

	foundOne := false
	for _, item := range except {
		if item == 1 {
			foundOne = true
			break
		}
	}
	if !foundOne {
		t.Fatal("Value 1 was not found")
	}

	foundFive := false
	for _, item := range except {
		if item == 5 {
			foundFive = true
			break
		}
	}
	if !foundFive {
		t.Fatal("Value 5 was not found")
	}
}

func TestExcept_string(t *testing.T) {
	a := []string{"who", "what", "when"}
	b := []string{"when", "where", "why"}

	exceptI, err := Except(a, b)
	if err != nil {
		t.Fatal(err)
	}
	except := exceptI.([]string)

	if len(except) != 2 {
		t.Fatal("Two items should be returned")
	}

	foundWho := false
	for _, item := range except {
		if item == "who" {
			foundWho = true
			break
		}
	}
	if !foundWho {
		t.Fatal("Value 'who' was not found")
	}

	foundWhat := false
	for _, item := range except {
		if item == "what" {
			foundWhat = true
			break
		}
	}
	if !foundWhat {
		t.Fatal("Value 'what' was not found")
	}
}
