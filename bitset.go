package main

import (
	"bytes"
	"errors"
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
	String() string
}

func newBitSet(len int) BitSet {
	return &bitSet{len, make([]uint64, (len+63)/64)}
}
func (b *bitSet) Len() int {
	return b.len
}

func calcUpDigit64(pos int) uint64 {
	if 0 < pos && 64 < pos {
		log.Fatalf("Cannot set more than 64, you set ", pos)
	}
	return 1 << uint(pos)

}
func (b *bitSet) Set(pos int) (e error) {
	if pos < b.len {
		e = errors.New("pos=" + string(pos) + " is bigger than len=" + string(b.len))
	}
	bodyIdx := pos / 64
	shifts := pos % 64
	value := b.body[bodyIdx] | calcUpDigit64(shifts)
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

func (b *bitSet) String() string {
	len := b.len
	println(len)
	var buf bytes.Buffer
	for _, x := range b.body {
		if len < 64 {
			buf.WriteString(uint64ToBinary(x, len))

		} else {
			buf.WriteString(uint64ToBinary(x, 64))
			buf.WriteString(" ")
		}
		len -= 64
	}
	return buf.String()
}
