package sets

import (
	"testing"
)

func TestU32String(t *testing.T) {
	ConfirmString := func(s u32set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(U32Set(), "()")
	ConfirmString(U32Set(0, 1), "(0 1)")
	ConfirmString(U32Set(1, 0), "(0 1)")
	ConfirmString(U32Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestU32Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r u32set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(U32Set(), U32Set(), U32Set())
	ConfirmIntersection(U32Set(0), U32Set(1), U32Set())
	ConfirmIntersection(U32Set(0, 1), U32Set(1, 2), U32Set(1))
	ConfirmIntersection(U32Set(0, 1, 2), U32Set(1, 2, 3), U32Set(1, 2))
}

func TestU32Union(t *testing.T) {
	ConfirmUnion := func(s, x, r u32set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(U32Set(), U32Set(), U32Set())
	ConfirmUnion(U32Set(0), U32Set(), U32Set(0))
	ConfirmUnion(U32Set(), U32Set(1), U32Set(1))
	ConfirmUnion(U32Set(0), U32Set(1), U32Set(0, 1))
	ConfirmUnion(U32Set(0, 1), U32Set(1, 2), U32Set(0, 1, 2))
	ConfirmUnion(U32Set(0, 1, 2), U32Set(1, 2, 3), U32Set(0, 1, 2, 3))
}

func TestU32Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r u32set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(U32Set(), U32Set(), U32Set())
	ConfirmDifference(U32Set(0), U32Set(), U32Set(0))
	ConfirmDifference(U32Set(), U32Set(1), U32Set())
	ConfirmDifference(U32Set(0), U32Set(1), U32Set(0))
	ConfirmDifference(U32Set(0, 1), U32Set(1, 2), U32Set(0))
	ConfirmDifference(U32Set(0, 1, 2), U32Set(1, 2, 3), U32Set(0))
	ConfirmDifference(U32Set(0, 1, 2, 3), U32Set(1, 2, 3), U32Set(0))
}

func TestU32SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x u32set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(U32Set(), U32Set(), true)
	ConfirmSubsetOf(U32Set(0), U32Set(), false)
	ConfirmSubsetOf(U32Set(), U32Set(0), true)
	ConfirmSubsetOf(U32Set(0), U32Set(0), true)
	ConfirmSubsetOf(U32Set(0), U32Set(1), false)
	ConfirmSubsetOf(U32Set(0), U32Set(0, 1), true)
	ConfirmSubsetOf(U32Set(0, 1), U32Set(0, 1), true)
	ConfirmSubsetOf(U32Set(0, 1, 2), U32Set(0, 1), false)
}

func TestU32Member(t *testing.T) {
	ConfirmMember := func(s u32set, x uint32, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(U32Set(), 0, false)
	ConfirmMember(U32Set(0), 0, true)
	ConfirmMember(U32Set(0, 1), 0, true)
	ConfirmMember(U32Set(0, 1), 1, true)
	ConfirmMember(U32Set(0, 1), 2, false)
}

func TestU32Equal(t *testing.T) {
	ConfirmEqual := func(s, x u32set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(U32Set(), U32Set(), true)
	ConfirmEqual(U32Set(0), U32Set(), false)
	ConfirmEqual(U32Set(), U32Set(0), false)
	ConfirmEqual(U32Set(0), U32Set(0), true)
	ConfirmEqual(U32Set(0, 0), U32Set(0), true)
	ConfirmEqual(U32Set(0), U32Set(0, 0), true)
	ConfirmEqual(U32Set(0, 1), U32Set(0, 0), false)
	ConfirmEqual(U32Set(0, 1), U32Set(0, 1), true)
	ConfirmEqual(U32Set(0, 1), U32Set(1, 0), true)
	ConfirmEqual(U32Set(0, 1), U32Set(1, 1), false)
}

func TestU32Sum(t *testing.T) {
	ConfirmSum := func(s u32set, r uint32) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(U32Set(), 0)
	ConfirmSum(U32Set(0), 0)
	ConfirmSum(U32Set(0, 1), 1)
	ConfirmSum(U32Set(0, 1, 1), 1)
	ConfirmSum(U32Set(0, 1, 2), 3)
}

func TestU32Product(t *testing.T) {
	ConfirmProduct := func(s u32set, r uint32) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(U32Set(), 1)
	ConfirmProduct(U32Set(0), 0)
	ConfirmProduct(U32Set(0, 1), 0)
	ConfirmProduct(U32Set(1), 1)
	ConfirmProduct(U32Set(1, 1), 1)
	ConfirmProduct(U32Set(0, 1, 1), 0)
	ConfirmProduct(U32Set(0, 1, 2), 0)
	ConfirmProduct(U32Set(1, 2), 2)
	ConfirmProduct(U32Set(1, 2, 3), 6)
	ConfirmProduct(U32Set(1, 2, 3, 3), 6)
}