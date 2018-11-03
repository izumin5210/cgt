package main

import "testing"

func TestA(t *testing.T) {
	// no-op
}

func TestB(t *testing.T) {
	t.Error("error!")
	t.Fatal("failed")
}

func TestC(t *testing.T) {
	t.Skip("skipped")
	t.Log("hello")
}

func TestD(t *testing.T) {
	t.Parallel()
	t.Log("hello")
}

func TestE(t *testing.T) {
	t.Parallel()
	t.Run("1", func(t *testing.T) {
		// no-op
	})
	t.Run("2", func(t *testing.T) {
		t.Error("error!")
	})
	t.Run("3", func(t *testing.T) {
		t.Skip("skipped")
	})
	t.Run("4", func(t *testing.T) {
		t.Error("error!")
	})
}

func TestF(t *testing.T) {
	t.Parallel()
	t.Error("error!")
}
