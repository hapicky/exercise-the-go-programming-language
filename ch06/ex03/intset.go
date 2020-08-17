package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	var min_len int
	if len(s.words) < len(t.words) {
		min_len = len(s.words)
	} else {
		min_len = len(t.words)
	}

	var i IntSet
	for n := 0; n < min_len; n++ {
		i.words = append(i.words, s.words[n]&t.words[n])
	}
	return &i
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var d IntSet
	for i := 0; i < len(s.words); i++ {
		var t_word uint64
		if len(t.words) > i {
			t_word = t.words[i]
		}
		d.words = append(d.words, (s.words[i] & ^(t_word)))
	}
	return &d
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		l += bits.OnesCount64(word)
	}
	return l
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.UnionWith(s)
	return &c
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(x.Len())

	x.Remove(192)
	x.Remove(9)
	fmt.Println(x.String())

	x.Clear()
	fmt.Println(x.String())

	z := y.Copy()
	fmt.Println(z.String())

	z.AddAll(1, 2, 3)
	fmt.Println(z.String())

	var i1, i2 IntSet
	i1.AddAll(1, 2, 3)
	i2.AddAll(2, 3, 4)
	fmt.Println(i1.IntersectWith(&i2))

	var d1, d2 IntSet
	d1.AddAll(1, 2, 3, 64)
	d2.AddAll(2, 3)
	fmt.Println(d1.DifferenceWith(&d2))
}
