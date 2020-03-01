package main

// START OMIT

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

// END OMIT

// func TestAdd(t *testing.T) {
// 	result := add(36, 6)
// 	if result != 42 {
// 		t.Errorf("%d is not the meaning of life", result)
// 	}
// }

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestAdd", F: TestAdd})
	testing.Verbose()
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
