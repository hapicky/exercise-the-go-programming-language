package main

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint
}

const pBit = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/pBit, uint(x%pBit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/pBit, uint(x%pBit)
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

	var ret IntSet
	for i := 0; i < min_len; i++ {
		ret.words = append(ret.words, s.words[i]&t.words[i])
	}
	return &ret
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var ret IntSet
	for i := 0; i < len(s.words); i++ {
		var t_word uint
		if len(t.words) > i {
			t_word = t.words[i]
		}
		ret.words = append(ret.words, (s.words[i] & ^(t_word)))
	}
	return &ret
}

func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var max_len int
	if len(s.words) > len(t.words) {
		max_len = len(s.words)
	} else {
		max_len = len(t.words)
	}

	var ret IntSet
	for i := 0; i < max_len; i++ {
		var s_word, t_word uint
		if len(s.words) > i {
			s_word = s.words[i]
		}
		if len(t.words) > i {
			t_word = t.words[i]
		}
		ret.words = append(ret.words, s_word^t_word)
	}
	return &ret
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < pBit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", pBit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		l += bits.OnesCount(word)
	}
	return l
}

func (s *IntSet) Remove(x int) {
	word, bit := x/pBit, uint(x%pBit)
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

func (s *IntSet) Elems() []int {
	var ret []int
	for i, word := range s.words {
		for n := 0; n < pBit; n++ {
			if (word & (1 << n)) > 0 {
				ret = append(ret, i*pBit+n)
			}
		}
	}
	return ret
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

	var s1, s2 IntSet
	s1.AddAll(1, 2, 3, 64)
	s2.AddAll(2, 3, 150)
	fmt.Println(s1.SymmetricDifference(&s2))

	var e IntSet
	fmt.Println(e.Elems())
	e.AddAll(0, 2, 4, 64)
	fmt.Println(e.Elems())
}
