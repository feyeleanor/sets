package sets

import (
	"testing"
)

func TestU16String(t *testing.T) {
	ConfirmString := func(s u16set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(U16Set(), "()")
	ConfirmString(U16Set(0, 1), "(0 1)")
	ConfirmString(U16Set(1, 0), "(0 1)")
	ConfirmString(U16Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestU16Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r u16set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(U16Set(), U16Set(), U16Set())
	ConfirmIntersection(U16Set(0), U16Set(1), U16Set())
	ConfirmIntersection(U16Set(0, 1), U16Set(1, 2), U16Set(1))
	ConfirmIntersection(U16Set(0, 1, 2), U16Set(1, 2, 3), U16Set(1, 2))
}

func TestU16Union(t *testing.T) {
	ConfirmUnion := func(s, x, r u16set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(U16Set(), U16Set(), U16Set())
	ConfirmUnion(U16Set(0), U16Set(), U16Set(0))
	ConfirmUnion(U16Set(), U16Set(1), U16Set(1))
	ConfirmUnion(U16Set(0), U16Set(1), U16Set(0, 1))
	ConfirmUnion(U16Set(0, 1), U16Set(1, 2), U16Set(0, 1, 2))
	ConfirmUnion(U16Set(0, 1, 2), U16Set(1, 2, 3), U16Set(0, 1, 2, 3))
}

func TestU16Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r u16set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(U16Set(), U16Set(), U16Set())
	ConfirmDifference(U16Set(0), U16Set(), U16Set(0))
	ConfirmDifference(U16Set(), U16Set(1), U16Set())
	ConfirmDifference(U16Set(0), U16Set(1), U16Set(0))
	ConfirmDifference(U16Set(0, 1), U16Set(1, 2), U16Set(0))
	ConfirmDifference(U16Set(0, 1, 2), U16Set(1, 2, 3), U16Set(0))
	ConfirmDifference(U16Set(0, 1, 2, 3), U16Set(1, 2, 3), U16Set(0))
}

func TestU16SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x u16set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(U16Set(), U16Set(), true)
	ConfirmSubsetOf(U16Set(0), U16Set(), false)
	ConfirmSubsetOf(U16Set(), U16Set(0), true)
	ConfirmSubsetOf(U16Set(0), U16Set(0), true)
	ConfirmSubsetOf(U16Set(0), U16Set(1), false)
	ConfirmSubsetOf(U16Set(0), U16Set(0, 1), true)
	ConfirmSubsetOf(U16Set(0, 1), U16Set(0, 1), true)
	ConfirmSubsetOf(U16Set(0, 1, 2), U16Set(0, 1), false)
}

func TestU16Member(t *testing.T) {
	ConfirmMember := func(s u16set, x uint16, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(U16Set(), 0, false)
	ConfirmMember(U16Set(0), 0, true)
	ConfirmMember(U16Set(0, 1), 0, true)
	ConfirmMember(U16Set(0, 1), 1, true)
	ConfirmMember(U16Set(0, 1), 2, false)
}

func TestU16Equal(t *testing.T) {
	ConfirmEqual := func(s, x u16set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(U16Set(), U16Set(), true)
	ConfirmEqual(U16Set(0), U16Set(), false)
	ConfirmEqual(U16Set(), U16Set(0), false)
	ConfirmEqual(U16Set(0), U16Set(0), true)
	ConfirmEqual(U16Set(0, 0), U16Set(0), true)
	ConfirmEqual(U16Set(0), U16Set(0, 0), true)
	ConfirmEqual(U16Set(0, 1), U16Set(0, 0), false)
	ConfirmEqual(U16Set(0, 1), U16Set(0, 1), true)
	ConfirmEqual(U16Set(0, 1), U16Set(1, 0), true)
	ConfirmEqual(U16Set(0, 1), U16Set(1, 1), false)
}

func TestU16Sum(t *testing.T) {
	ConfirmSum := func(s u16set, r uint16) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(U16Set(), 0)
	ConfirmSum(U16Set(0), 0)
	ConfirmSum(U16Set(0, 1), 1)
	ConfirmSum(U16Set(0, 1, 1), 1)
	ConfirmSum(U16Set(0, 1, 2), 3)
}

func TestU16Product(t *testing.T) {
	ConfirmProduct := func(s u16set, r uint16) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(U16Set(), 1)
	ConfirmProduct(U16Set(0), 0)
	ConfirmProduct(U16Set(0, 1), 0)
	ConfirmProduct(U16Set(1), 1)
	ConfirmProduct(U16Set(1, 1), 1)
	ConfirmProduct(U16Set(0, 1, 1), 0)
	ConfirmProduct(U16Set(0, 1, 2), 0)
	ConfirmProduct(U16Set(1, 2), 2)
	ConfirmProduct(U16Set(1, 2, 3), 6)
	ConfirmProduct(U16Set(1, 2, 3, 3), 6)
}