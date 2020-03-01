package main

import "testing"

func setup() {

}
func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	defer teardown()

	m.Run()
}

func TestSomething(t *testing.T) {}

func main() {}
