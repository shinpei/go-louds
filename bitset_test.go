package main

import (
	"testing"
)

func TestAllocateZero(t *testing.T) {
	b := newBitSet(0)
	if b.Len() != 0 {
		t.Errorf("Empty bitset should have capacity 0, not %d", b.Len())
	}
}

func TestAllocateVariable(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("No further trace")
		}
	}()
	var b bitSet
	if b.Len() != 0 {
		t.Errorf("Empty bitset should have capacity 0, not %d", b.Len())
	}
}

func TestStringBitset(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	b := newBitSet(123)
	b.String()
}
