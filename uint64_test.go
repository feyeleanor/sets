package sets

import (
	"testing"
)

func TestU64String(t *testing.T) {
	ConfirmString := func(s u64set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(U64Set(), "()")
	ConfirmString(U64Set(0, 1), "(0 1)")
	ConfirmString(U64Set(1, 0), "(0 1)")
	ConfirmString(U64Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestU64Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r u64set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(U64Set(), U64Set(), U64Set())
	ConfirmIntersection(U64Set(0), U64Set(1), U64Set())
	ConfirmIntersection(U64Set(0, 1), U64Set(1, 2), U64Set(1))
	ConfirmIntersection(U64Set(0, 1, 2), U64Set(1, 2, 3), U64Set(1, 2))
}

func TestU64Union(t *testing.T) {
	ConfirmUnion := func(s, x, r u64set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(U64Set(), U64Set(), U64Set())
	ConfirmUnion(U64Set(0), U64Set(), U64Set(0))
	ConfirmUnion(U64Set(), U64Set(1), U64Set(1))
	ConfirmUnion(U64Set(0), U64Set(1), U64Set(0, 1))
	ConfirmUnion(U64Set(0, 1), U64Set(1, 2), U64Set(0, 1, 2))
	ConfirmUnion(U64Set(0, 1, 2), U64Set(1, 2, 3), U64Set(0, 1, 2, 3))
}

func TestU64Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r u64set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(U64Set(), U64Set(), U64Set())
	ConfirmDifference(U64Set(0), U64Set(), U64Set(0))
	ConfirmDifference(U64Set(), U64Set(1), U64Set())
	ConfirmDifference(U64Set(0), U64Set(1), U64Set(0))
	ConfirmDifference(U64Set(0, 1), U64Set(1, 2), U64Set(0))
	ConfirmDifference(U64Set(0, 1, 2), U64Set(1, 2, 3), U64Set(0))
	ConfirmDifference(U64Set(0, 1, 2, 3), U64Set(1, 2, 3), U64Set(0))
}

func TestU64SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x u64set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(U64Set(), U64Set(), true)
	ConfirmSubsetOf(U64Set(0), U64Set(), false)
	ConfirmSubsetOf(U64Set(), U64Set(0), true)
	ConfirmSubsetOf(U64Set(0), U64Set(0), true)
	ConfirmSubsetOf(U64Set(0), U64Set(1), false)
	ConfirmSubsetOf(U64Set(0), U64Set(0, 1), true)
	ConfirmSubsetOf(U64Set(0, 1), U64Set(0, 1), true)
	ConfirmSubsetOf(U64Set(0, 1, 2), U64Set(0, 1), false)
}

func TestU64Member(t *testing.T) {
	ConfirmMember := func(s u64set, x uint64, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(U64Set(), 0, false)
	ConfirmMember(U64Set(0), 0, true)
	ConfirmMember(U64Set(0, 1), 0, true)
	ConfirmMember(U64Set(0, 1), 1, true)
	ConfirmMember(U64Set(0, 1), 2, false)
}

func TestU64Equal(t *testing.T) {
	ConfirmEqual := func(s, x u64set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(U64Set(), U64Set(), true)
	ConfirmEqual(U64Set(0), U64Set(), false)
	ConfirmEqual(U64Set(), U64Set(0), false)
	ConfirmEqual(U64Set(0), U64Set(0), true)
	ConfirmEqual(U64Set(0, 0), U64Set(0), true)
	ConfirmEqual(U64Set(0), U64Set(0, 0), true)
	ConfirmEqual(U64Set(0, 1), U64Set(0, 0), false)
	ConfirmEqual(U64Set(0, 1), U64Set(0, 1), true)
	ConfirmEqual(U64Set(0, 1), U64Set(1, 0), true)
	ConfirmEqual(U64Set(0, 1), U64Set(1, 1), false)
}

func TestU64Sum(t *testing.T) {
	ConfirmSum := func(s u64set, r uint64) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(U64Set(), 0)
	ConfirmSum(U64Set(0), 0)
	ConfirmSum(U64Set(0, 1), 1)
	ConfirmSum(U64Set(0, 1, 1), 1)
	ConfirmSum(U64Set(0, 1, 2), 3)
}

func TestU64Product(t *testing.T) {
	ConfirmProduct := func(s u64set, r uint64) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(U64Set(), 1)
	ConfirmProduct(U64Set(0), 0)
	ConfirmProduct(U64Set(0, 1), 0)
	ConfirmProduct(U64Set(1), 1)
	ConfirmProduct(U64Set(1, 1), 1)
	ConfirmProduct(U64Set(0, 1, 1), 0)
	ConfirmProduct(U64Set(0, 1, 2), 0)
	ConfirmProduct(U64Set(1, 2), 2)
	ConfirmProduct(U64Set(1, 2, 3), 6)
	ConfirmProduct(U64Set(1, 2, 3, 3), 6)
}