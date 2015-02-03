package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
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

func uint64ToBinary(value uint64, l int) string {
	var buf bytes.Buffer

	// Count digits first
	tmp := value
	digits := 1
	for 1 < tmp {
		tmp /= 2
		digits++
	}
	println("digits=", digits)
	if l < digits {
		log.Fatalf("Specified digits=" + string(l) + " are not enough.")
	}
	for 1 < value {
		buf.WriteString(strconv.Itoa(int(value % 2)))
		value /= 2
	}

	buf.WriteString(strconv.Itoa(int(value)))

	s := buf.String()
	// reverse
	buf.Reset()
	// put heading zeros
	for x := 0; x < l-digits; x++ {
		buf.WriteString("0")
	}
	for idx := len(s); idx > 0; idx-- {
		buf.WriteString(s[idx-1 : idx])
	}
	return buf.String()
}

func (b *bitSet) String() {
	println("length:", b.len)
	for _, x := range b.body {
		fmt.Print(uint64ToBinary(x, 64))
		fmt.Print(" ")
	}
	fmt.Println("")
}
