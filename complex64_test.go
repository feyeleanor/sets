package sets

import (
	"testing"
)

func TestC64String(t *testing.T) {
	ConfirmString := func(s c64set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(C64Set(), "()")
	ConfirmString(C64Set(0, 1), "((0+0i) (1+0i))")
	ConfirmString(C64Set(1, 0), "((0+0i) (1+0i))")
	ConfirmString(C64Set(0, 1, 2, 3, 4), "((0+0i) (1+0i) (2+0i) (3+0i) (4+0i))")
}

func TestC64Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r c64set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(C64Set(), C64Set(), C64Set())
	ConfirmIntersection(C64Set(0), C64Set(1), C64Set())
	ConfirmIntersection(C64Set(0, 1), C64Set(1, 2), C64Set(1))
	ConfirmIntersection(C64Set(0, 1, 2), C64Set(1, 2, 3), C64Set(1, 2))
}

func TestC64Union(t *testing.T) {
	ConfirmUnion := func(s, x, r c64set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(C64Set(), C64Set(), C64Set())
	ConfirmUnion(C64Set(0), C64Set(), C64Set(0))
	ConfirmUnion(C64Set(), C64Set(1), C64Set(1))
	ConfirmUnion(C64Set(0), C64Set(1), C64Set(0, 1))
	ConfirmUnion(C64Set(0, 1), C64Set(1, 2), C64Set(0, 1, 2))
	ConfirmUnion(C64Set(0, 1, 2), C64Set(1, 2, 3), C64Set(0, 1, 2, 3))
}

func TestC64Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r c64set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(C64Set(), C64Set(), C64Set())
	ConfirmDifference(C64Set(0), C64Set(), C64Set(0))
	ConfirmDifference(C64Set(), C64Set(1), C64Set())
	ConfirmDifference(C64Set(0), C64Set(1), C64Set(0))
	ConfirmDifference(C64Set(0, 1), C64Set(1, 2), C64Set(0))
	ConfirmDifference(C64Set(0, 1, 2), C64Set(1, 2, 3), C64Set(0))
	ConfirmDifference(C64Set(0, 1, 2, 3), C64Set(1, 2, 3), C64Set(0))
}

func TestC64SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x c64set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(C64Set(), C64Set(), true)
	ConfirmSubsetOf(C64Set(0), C64Set(), false)
	ConfirmSubsetOf(C64Set(), C64Set(0), true)
	ConfirmSubsetOf(C64Set(0), C64Set(0), true)
	ConfirmSubsetOf(C64Set(0), C64Set(1), false)
	ConfirmSubsetOf(C64Set(0), C64Set(0, 1), true)
	ConfirmSubsetOf(C64Set(0, 1), C64Set(0, 1), true)
	ConfirmSubsetOf(C64Set(0, 1, 2), C64Set(0, 1), false)
}

func TestC64Member(t *testing.T) {
	ConfirmMember := func(s c64set, x complex64, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(C64Set(), 0, false)
	ConfirmMember(C64Set(0), 0, true)
	ConfirmMember(C64Set(0, 1), 0, true)
	ConfirmMember(C64Set(0, 1), 1, true)
	ConfirmMember(C64Set(0, 1), 2, false)
}

func TestC64Equal(t *testing.T) {
	ConfirmEqual := func(s, x c64set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(C64Set(), C64Set(), true)
	ConfirmEqual(C64Set(0), C64Set(), false)
	ConfirmEqual(C64Set(), C64Set(0), false)
	ConfirmEqual(C64Set(0), C64Set(0), true)
	ConfirmEqual(C64Set(0, 0), C64Set(0), true)
	ConfirmEqual(C64Set(0), C64Set(0, 0), true)
	ConfirmEqual(C64Set(0, 1), C64Set(0, 0), false)
	ConfirmEqual(C64Set(0, 1), C64Set(0, 1), true)
	ConfirmEqual(C64Set(0, 1), C64Set(1, 0), true)
	ConfirmEqual(C64Set(0, 1), C64Set(1, 1), false)
}

func TestC64Sum(t *testing.T) {
	ConfirmSum := func(s c64set, r complex64) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(C64Set(), 0)
	ConfirmSum(C64Set(0), 0)
	ConfirmSum(C64Set(0, 1), 1)
	ConfirmSum(C64Set(0, 1, 1), 1)
	ConfirmSum(C64Set(0, 1, 2), 3)
}

func TestC64Product(t *testing.T) {
	ConfirmProduct := func(s c64set, r complex64) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(C64Set(), 1)
	ConfirmProduct(C64Set(0), 0)
	ConfirmProduct(C64Set(0, 1), 0)
	ConfirmProduct(C64Set(1), 1)
	ConfirmProduct(C64Set(1, 1), 1)
	ConfirmProduct(C64Set(0, 1, 1), 0)
	ConfirmProduct(C64Set(0, 1, 2), 0)
	ConfirmProduct(C64Set(1, 2), 2)
	ConfirmProduct(C64Set(1, 2, 3), 6)
	ConfirmProduct(C64Set(1, 2, 3, 3), 6)
}