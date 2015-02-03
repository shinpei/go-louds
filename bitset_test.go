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
	if uint64ToBinary(0, 1) != "0" {
		t.Errorf("should be '0'")
	}
	if uint64ToBinary(0, 2) != "00" {
		t.Errorf("should be '00'")
	}
	if uint64ToBinary(1, 1) != "1" {
		t.Errorf("should be '1'")
	}
	if uint64ToBinary(2, 2) != "10" {
		t.Errorf("should be '10'")
	}
	if uint64ToBinary(3, 2) != "11" {
		t.Errorf("should be '11'")
	}
	if uint64ToBinary(3, 3) != "011" {
		t.Errorf("should be '011'")
	}
	if uint64ToBinary(5, 3) != "101" {
		t.Errorf("should be '101'")
	}
	if uint64ToBinary(5, 64) != "0000000000000000000000000000000000000000000000000000000000000101" {
		t.Errorf("should be '000000000000000000000000000000000000000000000000000000000000101'")
	}
}

func TestStringBitset(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	b := newBitSet(123)
	b.Set(3)
	b.String()
}
