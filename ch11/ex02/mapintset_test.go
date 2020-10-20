package mapintset

import (
	"gopl.io/ch6/intset"
	"testing"
)

func TestHasEmpty(t *testing.T) {
	is := intset.IntSet{}
	ms := MapIntSet{}
	i := 100
	want := is.Has(i)
	got := ms.Has(i)

	if got != want {
		t.Errorf("Has(%d) = %v", i, got)
	}
}

func TestAdd(t *testing.T) {
	is := intset.IntSet{}
	ms := MapIntSet{}
	i := 100
	is.Add(i)
	want := is.Has(i)
	ms.Add(i)
	got := ms.Has(i)

	if got != want {
		t.Errorf("After Add(%d), Has(%d) = %v", i, i, got)
	}
}

func TestString(t *testing.T) {
	is := intset.IntSet{}
	ms := MapIntSet{}
	is.Add(1)
	is.Add(2)
	is.Add(3)
	want := is.String()
	ms.Add(1)
	ms.Add(2)
	ms.Add(3)
	got := ms.String()

	if got != want {
		t.Errorf("MapIntSet.String() = %q, want %q", got, want)
	}
}

func TestStringEmpty(t *testing.T) {
	is := intset.IntSet{}
	ms := MapIntSet{}
	want := is.String()
	got := ms.String()

	if got != want {
		t.Errorf("MapIntSet.String() = %q, want %q", got, want)
	}
}

func TestUnionWith(t *testing.T) {
	is1 := intset.IntSet{}
	is1.Add(1)
	is1.Add(2)
	is1.Add(3)

	is2 := intset.IntSet{}
	is2.Add(2)
	is2.Add(4)
	is2.Add(6)

	is1.UnionWith(&is2)

	ms1 := MapIntSet{}
	ms1.Add(1)
	ms1.Add(2)
	ms1.Add(3)

	ms2 := MapIntSet{}
	ms2.Add(2)
	ms2.Add(4)
	ms2.Add(6)

	ms1.UnionWith(&ms2)

	want := is1.String()
	got := ms1.String()

	if got != want {
		t.Errorf("MapIntSet.String() = %q, want %q", got, want)
	}
}
