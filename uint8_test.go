package sets

import (
	"testing"
)

func TestU8String(t *testing.T) {
	ConfirmString := func(s u8set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(U8Set(), "()")
	ConfirmString(U8Set(0, 1), "(0 1)")
	ConfirmString(U8Set(1, 0), "(0 1)")
	ConfirmString(U8Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestU8Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r u8set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(U8Set(), U8Set(), U8Set())
	ConfirmIntersection(U8Set(0), U8Set(1), U8Set())
	ConfirmIntersection(U8Set(0, 1), U8Set(1, 2), U8Set(1))
	ConfirmIntersection(U8Set(0, 1, 2), U8Set(1, 2, 3), U8Set(1, 2))
}

func TestU8Union(t *testing.T) {
	ConfirmUnion := func(s, x, r u8set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(U8Set(), U8Set(), U8Set())
	ConfirmUnion(U8Set(0), U8Set(), U8Set(0))
	ConfirmUnion(U8Set(), U8Set(1), U8Set(1))
	ConfirmUnion(U8Set(0), U8Set(1), U8Set(0, 1))
	ConfirmUnion(U8Set(0, 1), U8Set(1, 2), U8Set(0, 1, 2))
	ConfirmUnion(U8Set(0, 1, 2), U8Set(1, 2, 3), U8Set(0, 1, 2, 3))
}

func TestU8Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r u8set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(U8Set(), U8Set(), U8Set())
	ConfirmDifference(U8Set(0), U8Set(), U8Set(0))
	ConfirmDifference(U8Set(), U8Set(1), U8Set())
	ConfirmDifference(U8Set(0), U8Set(1), U8Set(0))
	ConfirmDifference(U8Set(0, 1), U8Set(1, 2), U8Set(0))
	ConfirmDifference(U8Set(0, 1, 2), U8Set(1, 2, 3), U8Set(0))
	ConfirmDifference(U8Set(0, 1, 2, 3), U8Set(1, 2, 3), U8Set(0))
}

func TestU8SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x u8set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(U8Set(), U8Set(), true)
	ConfirmSubsetOf(U8Set(0), U8Set(), false)
	ConfirmSubsetOf(U8Set(), U8Set(0), true)
	ConfirmSubsetOf(U8Set(0), U8Set(0), true)
	ConfirmSubsetOf(U8Set(0), U8Set(1), false)
	ConfirmSubsetOf(U8Set(0), U8Set(0, 1), true)
	ConfirmSubsetOf(U8Set(0, 1), U8Set(0, 1), true)
	ConfirmSubsetOf(U8Set(0, 1, 2), U8Set(0, 1), false)
}

func TestU8Member(t *testing.T) {
	ConfirmMember := func(s u8set, x uint8, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(U8Set(), 0, false)
	ConfirmMember(U8Set(0), 0, true)
	ConfirmMember(U8Set(0, 1), 0, true)
	ConfirmMember(U8Set(0, 1), 1, true)
	ConfirmMember(U8Set(0, 1), 2, false)
}

func TestU8Equal(t *testing.T) {
	ConfirmEqual := func(s, x u8set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(U8Set(), U8Set(), true)
	ConfirmEqual(U8Set(0), U8Set(), false)
	ConfirmEqual(U8Set(), U8Set(0), false)
	ConfirmEqual(U8Set(0), U8Set(0), true)
	ConfirmEqual(U8Set(0, 0), U8Set(0), true)
	ConfirmEqual(U8Set(0), U8Set(0, 0), true)
	ConfirmEqual(U8Set(0, 1), U8Set(0, 0), false)
	ConfirmEqual(U8Set(0, 1), U8Set(0, 1), true)
	ConfirmEqual(U8Set(0, 1), U8Set(1, 0), true)
	ConfirmEqual(U8Set(0, 1), U8Set(1, 1), false)
}

func TestU8Sum(t *testing.T) {
	ConfirmSum := func(s u8set, r uint8) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(U8Set(), 0)
	ConfirmSum(U8Set(0), 0)
	ConfirmSum(U8Set(0, 1), 1)
	ConfirmSum(U8Set(0, 1, 1), 1)
	ConfirmSum(U8Set(0, 1, 2), 3)
}

func TestU8Product(t *testing.T) {
	ConfirmProduct := func(s u8set, r uint8) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(U8Set(), 1)
	ConfirmProduct(U8Set(0), 0)
	ConfirmProduct(U8Set(0, 1), 0)
	ConfirmProduct(U8Set(1), 1)
	ConfirmProduct(U8Set(1, 1), 1)
	ConfirmProduct(U8Set(0, 1, 1), 0)
	ConfirmProduct(U8Set(0, 1, 2), 0)
	ConfirmProduct(U8Set(1, 2), 2)
	ConfirmProduct(U8Set(1, 2, 3), 6)
	ConfirmProduct(U8Set(1, 2, 3, 3), 6)
}