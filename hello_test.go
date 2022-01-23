package main

import "testing"

/**
编写测试
*/
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
