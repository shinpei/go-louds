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

func TestUint64ToBinary(t *testing.T) {
	println(uint64ToBinary(0, 1))
	println(uint64ToBinary(1, 1))
	println(uint64ToBinary(2, 2))
	println(uint64ToBinary(3, 3))
	println(uint64ToBinary(5, 3))
	println(uint64ToBinary(5, 64))
}

/*
func TestStringBitset(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	b := newBitSet(123)
	b.Set(3)
	b.String()
}
*/
