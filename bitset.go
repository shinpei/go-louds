package main

import (
	"errors"
)

type bitSet struct {
	len  int
	body []uint64
}

type BitSet interface {
	Set(pos int) error
	Test(pos int) bool // x = 001, x.Test(0) = false
	Len() int
	String()
}

func newBitSet(len int) BitSet {
	return &bitSet{len, make([]uint64, (len+63)/64)}
}
func (b *bitSet) Len() int {
	return b.len
}

func (b *bitSet) Set(pos int) (e error) {
	if pos < b.len {
		e = errors.New("pos=" + string(pos) + " is bigger than len=" + string(b.len))
	}
	bodyIdx := pos / 64
	setPos := uint64(pos - bodyIdx*64)
	value := b.body[bodyIdx] | setPos
	b.body[bodyIdx] = value
	return
}

func (b *bitSet) Test(pos int) bool {
	//	bodyIdx := pos / 64
	return true
}

func (b *bitSet) String() {
	for _, x := range b.body {
		println(x)
	}
}
