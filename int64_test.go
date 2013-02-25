package sets

import (
	"testing"
)

func TestI64String(t *testing.T) {
	ConfirmString := func(s i64set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(I64Set(), "()")
	ConfirmString(I64Set(0, 1), "(0 1)")
	ConfirmString(I64Set(1, 0), "(0 1)")
	ConfirmString(I64Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestI64Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r i64set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(I64Set(), I64Set(), I64Set())
	ConfirmIntersection(I64Set(0), I64Set(1), I64Set())
	ConfirmIntersection(I64Set(0, 1), I64Set(1, 2), I64Set(1))
	ConfirmIntersection(I64Set(0, 1, 2), I64Set(1, 2, 3), I64Set(1, 2))
}

func TestI64Union(t *testing.T) {
	ConfirmUnion := func(s, x, r i64set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(I64Set(), I64Set(), I64Set())
	ConfirmUnion(I64Set(0), I64Set(), I64Set(0))
	ConfirmUnion(I64Set(), I64Set(1), I64Set(1))
	ConfirmUnion(I64Set(0), I64Set(1), I64Set(0, 1))
	ConfirmUnion(I64Set(0, 1), I64Set(1, 2), I64Set(0, 1, 2))
	ConfirmUnion(I64Set(0, 1, 2), I64Set(1, 2, 3), I64Set(0, 1, 2, 3))
}

func TestI64Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r i64set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(I64Set(), I64Set(), I64Set())
	ConfirmDifference(I64Set(0), I64Set(), I64Set(0))
	ConfirmDifference(I64Set(), I64Set(1), I64Set())
	ConfirmDifference(I64Set(0), I64Set(1), I64Set(0))
	ConfirmDifference(I64Set(0, 1), I64Set(1, 2), I64Set(0))
	ConfirmDifference(I64Set(0, 1, 2), I64Set(1, 2, 3), I64Set(0))
	ConfirmDifference(I64Set(0, 1, 2, 3), I64Set(1, 2, 3), I64Set(0))
}

func TestI64SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x i64set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(I64Set(), I64Set(), true)
	ConfirmSubsetOf(I64Set(0), I64Set(), false)
	ConfirmSubsetOf(I64Set(), I64Set(0), true)
	ConfirmSubsetOf(I64Set(0), I64Set(0), true)
	ConfirmSubsetOf(I64Set(0), I64Set(1), false)
	ConfirmSubsetOf(I64Set(0), I64Set(0, 1), true)
	ConfirmSubsetOf(I64Set(0, 1), I64Set(0, 1), true)
	ConfirmSubsetOf(I64Set(0, 1, 2), I64Set(0, 1), false)
}

func TestI64Member(t *testing.T) {
	ConfirmMember := func(s i64set, x int64, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(I64Set(), 0, false)
	ConfirmMember(I64Set(0), 0, true)
	ConfirmMember(I64Set(0, 1), 0, true)
	ConfirmMember(I64Set(0, 1), 1, true)
	ConfirmMember(I64Set(0, 1), 2, false)
}

func TestI64Equal(t *testing.T) {
	ConfirmEqual := func(s, x i64set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(I64Set(), I64Set(), true)
	ConfirmEqual(I64Set(0), I64Set(), false)
	ConfirmEqual(I64Set(), I64Set(0), false)
	ConfirmEqual(I64Set(0), I64Set(0), true)
	ConfirmEqual(I64Set(0, 0), I64Set(0), true)
	ConfirmEqual(I64Set(0), I64Set(0, 0), true)
	ConfirmEqual(I64Set(0, 1), I64Set(0, 0), false)
	ConfirmEqual(I64Set(0, 1), I64Set(0, 1), true)
	ConfirmEqual(I64Set(0, 1), I64Set(1, 0), true)
	ConfirmEqual(I64Set(0, 1), I64Set(1, 1), false)
}

func TestI64Sum(t *testing.T) {
	ConfirmSum := func(s i64set, r int64) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(I64Set(), 0)
	ConfirmSum(I64Set(0), 0)
	ConfirmSum(I64Set(0, 1), 1)
	ConfirmSum(I64Set(0, 1, 1), 1)
	ConfirmSum(I64Set(0, 1, 2), 3)
}

func TestI64Product(t *testing.T) {
	ConfirmProduct := func(s i64set, r int64) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(I64Set(), 1)
	ConfirmProduct(I64Set(0), 0)
	ConfirmProduct(I64Set(0, 1), 0)
	ConfirmProduct(I64Set(1), 1)
	ConfirmProduct(I64Set(1, 1), 1)
	ConfirmProduct(I64Set(0, 1, 1), 0)
	ConfirmProduct(I64Set(0, 1, 2), 0)
	ConfirmProduct(I64Set(1, 2), 2)
	ConfirmProduct(I64Set(1, 2, 3), 6)
	ConfirmProduct(I64Set(1, 2, 3, 3), 6)
}