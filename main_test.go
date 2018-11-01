package main

import "testing"

func TestSuccess(t *testing.T) {
	// no-op
	t.Log("hoge")
}

func TestFail(t *testing.T) {
	t.Error("error")
}
