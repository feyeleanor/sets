package sets

import (
	"testing"
)

func TestC128String(t *testing.T) {
	ConfirmString := func(s c128set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(C128Set(), "()")
	ConfirmString(C128Set(0, 1), "((0+0i) (1+0i))")
	ConfirmString(C128Set(1, 0), "((0+0i) (1+0i))")
	ConfirmString(C128Set(0, 1, 2, 3, 4), "((0+0i) (1+0i) (2+0i) (3+0i) (4+0i))")
}

func TestC128Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r c128set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(C128Set(), C128Set(), C128Set())
	ConfirmIntersection(C128Set(0), C128Set(1), C128Set())
	ConfirmIntersection(C128Set(0, 1), C128Set(1, 2), C128Set(1))
	ConfirmIntersection(C128Set(0, 1, 2), C128Set(1, 2, 3), C128Set(1, 2))
}

func TestC128Union(t *testing.T) {
	ConfirmUnion := func(s, x, r c128set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(C128Set(), C128Set(), C128Set())
	ConfirmUnion(C128Set(0), C128Set(), C128Set(0))
	ConfirmUnion(C128Set(), C128Set(1), C128Set(1))
	ConfirmUnion(C128Set(0), C128Set(1), C128Set(0, 1))
	ConfirmUnion(C128Set(0, 1), C128Set(1, 2), C128Set(0, 1, 2))
	ConfirmUnion(C128Set(0, 1, 2), C128Set(1, 2, 3), C128Set(0, 1, 2, 3))
}

func TestC128Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r c128set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(C128Set(), C128Set(), C128Set())
	ConfirmDifference(C128Set(0), C128Set(), C128Set(0))
	ConfirmDifference(C128Set(), C128Set(1), C128Set())
	ConfirmDifference(C128Set(0), C128Set(1), C128Set(0))
	ConfirmDifference(C128Set(0, 1), C128Set(1, 2), C128Set(0))
	ConfirmDifference(C128Set(0, 1, 2), C128Set(1, 2, 3), C128Set(0))
	ConfirmDifference(C128Set(0, 1, 2, 3), C128Set(1, 2, 3), C128Set(0))
}

func TestC128SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x c128set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(C128Set(), C128Set(), true)
	ConfirmSubsetOf(C128Set(0), C128Set(), false)
	ConfirmSubsetOf(C128Set(), C128Set(0), true)
	ConfirmSubsetOf(C128Set(0), C128Set(0), true)
	ConfirmSubsetOf(C128Set(0), C128Set(1), false)
	ConfirmSubsetOf(C128Set(0), C128Set(0, 1), true)
	ConfirmSubsetOf(C128Set(0, 1), C128Set(0, 1), true)
	ConfirmSubsetOf(C128Set(0, 1, 2), C128Set(0, 1), false)
}

func TestC128Member(t *testing.T) {
	ConfirmMember := func(s c128set, x complex128, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(C128Set(), 0, false)
	ConfirmMember(C128Set(0), 0, true)
	ConfirmMember(C128Set(0, 1), 0, true)
	ConfirmMember(C128Set(0, 1), 1, true)
	ConfirmMember(C128Set(0, 1), 2, false)
}

func TestC128Equal(t *testing.T) {
	ConfirmEqual := func(s, x c128set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(C128Set(), C128Set(), true)
	ConfirmEqual(C128Set(0), C128Set(), false)
	ConfirmEqual(C128Set(), C128Set(0), false)
	ConfirmEqual(C128Set(0), C128Set(0), true)
	ConfirmEqual(C128Set(0, 0), C128Set(0), true)
	ConfirmEqual(C128Set(0), C128Set(0, 0), true)
	ConfirmEqual(C128Set(0, 1), C128Set(0, 0), false)
	ConfirmEqual(C128Set(0, 1), C128Set(0, 1), true)
	ConfirmEqual(C128Set(0, 1), C128Set(1, 0), true)
	ConfirmEqual(C128Set(0, 1), C128Set(1, 1), false)
}

func TestC128Sum(t *testing.T) {
	ConfirmSum := func(s c128set, r complex128) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(C128Set(), 0)
	ConfirmSum(C128Set(0), 0)
	ConfirmSum(C128Set(0, 1), 1)
	ConfirmSum(C128Set(0, 1, 1), 1)
	ConfirmSum(C128Set(0, 1, 2), 3)
}

func TestC128Product(t *testing.T) {
	ConfirmProduct := func(s c128set, r complex128) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(C128Set(), 1)
	ConfirmProduct(C128Set(0), 0)
	ConfirmProduct(C128Set(0, 1), 0)
	ConfirmProduct(C128Set(1), 1)
	ConfirmProduct(C128Set(1, 1), 1)
	ConfirmProduct(C128Set(0, 1, 1), 0)
	ConfirmProduct(C128Set(0, 1, 2), 0)
	ConfirmProduct(C128Set(1, 2), 2)
	ConfirmProduct(C128Set(1, 2, 3), 6)
	ConfirmProduct(C128Set(1, 2, 3, 3), 6)
}