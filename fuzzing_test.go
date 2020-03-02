package main

// START OMIT

import (
	"testing"

	fuzz "github.com/google/gofuzz"
)

func Add(a, b int) int {
	return a / b
}

func TestAdd(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var a, b int

		fuzz.New().NumElements(0, 10).Fuzz(&a)
		fuzz.New().NumElements(0, 10).Fuzz(&b)

		t.Logf("Add %d and %d", a, b)
		Add(a, b)
	}
}

// END OMIT

func TestMain(m *testing.M) {
	m.Run()
}

// func main() {
// 	var tests []testing.InternalTest
// 	tests = append(tests, testing.InternalTest{Name: "TestFuzzAdd", F: TestAdd})
// 	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
// }
