package array

import (
	"fmt"
	"math/rand"
	"testing"
)

// Examples
func ExampleIntersection() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersections := Intersection(a, b).([]int)
	fmt.Println(intersections)
	// Output:
	// [2 3]
}

func ExampleExcept_a() {
	a := []int{1, 2, 5}
	b := []int{2, 3, 4}

	except := Except(a, b).([]int)
	fmt.Println(except)
	// Output:
	// [1 5]
}

func ExampleExcept_b() {
	a := []int{1, 2, 5}
	b := []int{2, 3, 4}

	except := Except(b, a).([]int)
	fmt.Println(except)
	// Output:
	// [3 4]
}

func ExampleExcept_string() {
	a := []string{"who", "what", "when"}
	b := []string{"when", "where", "why"}

	except := Except(a, b).([]string)
	fmt.Println(except)
	// Output:
	// [who what]
}

// Benchmark
func genTwoIntArr(len int) (a []int, b []int) {
	for i := 0; i < len; i++ {
		a = append(a, rand.Intn(100))
		b = append(b, rand.Intn(100))
	}
	return
}

func benchIntersection(len int, b *testing.B) {
	one, two := genTwoIntArr(len)
	for n := 0; n < b.N; n++ {
		Intersection(one, two)
	}
}

func BenchmarkIntersection100(b *testing.B)    { benchIntersection(100, b) }
func BenchmarkIntersection1000(b *testing.B)   { benchIntersection(1000, b) }
func BenchmarkIntersection10000(b *testing.B)  { benchIntersection(10000, b) }
func BenchmarkIntersection100000(b *testing.B) { benchIntersection(100000, b) }

// Tests
// Intersection

func hasInteger(i int, arr []int) bool {
	for _, item := range arr {
		if item == i {
			return true
		}
	}
	return false
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersects := Intersection(a, b).([]int)

	if len(intersects) != 2 {
		fmt.Println(intersects)
		t.Fatal("Result array has incorrect terms")
	}

	if !hasInteger(2, intersects) {
		t.Fatal("Number two not found")
	}

	if !hasInteger(3, intersects) {
		t.Fatal("Number three not found")
	}

}

// Except
func TestExcept_int(t *testing.T) {
	a := []int{1, 2, 5}
	b := []int{2, 3, 4}

	except := Except(a, b).([]int)

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

	except := Except(a, b).([]string)

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
