package sets

import (
	"testing"
)

func TestF64String(t *testing.T) {
	ConfirmString := func(s f64set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(F64Set(), "()")
	ConfirmString(F64Set(0, 1), "(0 1)")
	ConfirmString(F64Set(1, 0), "(0 1)")
	ConfirmString(F64Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestF64Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r f64set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(F64Set(), F64Set(), F64Set())
	ConfirmIntersection(F64Set(0), F64Set(1), F64Set())
	ConfirmIntersection(F64Set(0, 1), F64Set(1, 2), F64Set(1))
	ConfirmIntersection(F64Set(0, 1, 2), F64Set(1, 2, 3), F64Set(1, 2))
}

func TestF64Union(t *testing.T) {
	ConfirmUnion := func(s, x, r f64set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(F64Set(), F64Set(), F64Set())
	ConfirmUnion(F64Set(0), F64Set(), F64Set(0))
	ConfirmUnion(F64Set(), F64Set(1), F64Set(1))
	ConfirmUnion(F64Set(0), F64Set(1), F64Set(0, 1))
	ConfirmUnion(F64Set(0, 1), F64Set(1, 2), F64Set(0, 1, 2))
	ConfirmUnion(F64Set(0, 1, 2), F64Set(1, 2, 3), F64Set(0, 1, 2, 3))
}

func TestF64Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r f64set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(F64Set(), F64Set(), F64Set())
	ConfirmDifference(F64Set(0), F64Set(), F64Set(0))
	ConfirmDifference(F64Set(), F64Set(1), F64Set())
	ConfirmDifference(F64Set(0), F64Set(1), F64Set(0))
	ConfirmDifference(F64Set(0, 1), F64Set(1, 2), F64Set(0))
	ConfirmDifference(F64Set(0, 1, 2), F64Set(1, 2, 3), F64Set(0))
	ConfirmDifference(F64Set(0, 1, 2, 3), F64Set(1, 2, 3), F64Set(0))
}

func TestF64SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x f64set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(F64Set(), F64Set(), true)
	ConfirmSubsetOf(F64Set(0), F64Set(), false)
	ConfirmSubsetOf(F64Set(), F64Set(0), true)
	ConfirmSubsetOf(F64Set(0), F64Set(0), true)
	ConfirmSubsetOf(F64Set(0), F64Set(1), false)
	ConfirmSubsetOf(F64Set(0), F64Set(0, 1), true)
	ConfirmSubsetOf(F64Set(0, 1), F64Set(0, 1), true)
	ConfirmSubsetOf(F64Set(0, 1, 2), F64Set(0, 1), false)
}

func TestF64Member(t *testing.T) {
	ConfirmMember := func(s f64set, x float64, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(F64Set(), 0, false)
	ConfirmMember(F64Set(0), 0, true)
	ConfirmMember(F64Set(0, 1), 0, true)
	ConfirmMember(F64Set(0, 1), 1, true)
	ConfirmMember(F64Set(0, 1), 2, false)
}

func TestF64Equal(t *testing.T) {
	ConfirmEqual := func(s, x f64set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(F64Set(), F64Set(), true)
	ConfirmEqual(F64Set(0), F64Set(), false)
	ConfirmEqual(F64Set(), F64Set(0), false)
	ConfirmEqual(F64Set(0), F64Set(0), true)
	ConfirmEqual(F64Set(0, 0), F64Set(0), true)
	ConfirmEqual(F64Set(0), F64Set(0, 0), true)
	ConfirmEqual(F64Set(0, 1), F64Set(0, 0), false)
	ConfirmEqual(F64Set(0, 1), F64Set(0, 1), true)
	ConfirmEqual(F64Set(0, 1), F64Set(1, 0), true)
	ConfirmEqual(F64Set(0, 1), F64Set(1, 1), false)
}

func TestF64Sum(t *testing.T) {
	ConfirmSum := func(s f64set, r float64) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(F64Set(), 0)
	ConfirmSum(F64Set(0), 0)
	ConfirmSum(F64Set(0, 1), 1)
	ConfirmSum(F64Set(0, 1, 1), 1)
	ConfirmSum(F64Set(0, 1, 2), 3)
}

func TestF64Product(t *testing.T) {
	ConfirmProduct := func(s f64set, r float64) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(F64Set(), 1)
	ConfirmProduct(F64Set(0), 0)
	ConfirmProduct(F64Set(0, 1), 0)
	ConfirmProduct(F64Set(1), 1)
	ConfirmProduct(F64Set(1, 1), 1)
	ConfirmProduct(F64Set(0, 1, 1), 0)
	ConfirmProduct(F64Set(0, 1, 2), 0)
	ConfirmProduct(F64Set(1, 2), 2)
	ConfirmProduct(F64Set(1, 2, 3), 6)
	ConfirmProduct(F64Set(1, 2, 3, 3), 6)
}