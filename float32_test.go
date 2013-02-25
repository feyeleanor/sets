package sets

import (
	"testing"
)

func TestF32String(t *testing.T) {
	ConfirmString := func(s f32set, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(F32Set(), "()")
	ConfirmString(F32Set(0, 1), "(0 1)")
	ConfirmString(F32Set(1, 0), "(0 1)")
	ConfirmString(F32Set(0, 1, 2, 3, 4), "(0 1 2 3 4)")
}

func TestF32Intersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r f32set) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(F32Set(), F32Set(), F32Set())
	ConfirmIntersection(F32Set(0), F32Set(1), F32Set())
	ConfirmIntersection(F32Set(0, 1), F32Set(1, 2), F32Set(1))
	ConfirmIntersection(F32Set(0, 1, 2), F32Set(1, 2, 3), F32Set(1, 2))
}

func TestF32Union(t *testing.T) {
	ConfirmUnion := func(s, x, r f32set) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(F32Set(), F32Set(), F32Set())
	ConfirmUnion(F32Set(0), F32Set(), F32Set(0))
	ConfirmUnion(F32Set(), F32Set(1), F32Set(1))
	ConfirmUnion(F32Set(0), F32Set(1), F32Set(0, 1))
	ConfirmUnion(F32Set(0, 1), F32Set(1, 2), F32Set(0, 1, 2))
	ConfirmUnion(F32Set(0, 1, 2), F32Set(1, 2, 3), F32Set(0, 1, 2, 3))
}

func TestF32Difference(t *testing.T) {
	ConfirmDifference := func(s, x, r f32set) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(F32Set(), F32Set(), F32Set())
	ConfirmDifference(F32Set(0), F32Set(), F32Set(0))
	ConfirmDifference(F32Set(), F32Set(1), F32Set())
	ConfirmDifference(F32Set(0), F32Set(1), F32Set(0))
	ConfirmDifference(F32Set(0, 1), F32Set(1, 2), F32Set(0))
	ConfirmDifference(F32Set(0, 1, 2), F32Set(1, 2, 3), F32Set(0))
	ConfirmDifference(F32Set(0, 1, 2, 3), F32Set(1, 2, 3), F32Set(0))
}

func TestF32SubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x f32set, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(F32Set(), F32Set(), true)
	ConfirmSubsetOf(F32Set(0), F32Set(), false)
	ConfirmSubsetOf(F32Set(), F32Set(0), true)
	ConfirmSubsetOf(F32Set(0), F32Set(0), true)
	ConfirmSubsetOf(F32Set(0), F32Set(1), false)
	ConfirmSubsetOf(F32Set(0), F32Set(0, 1), true)
	ConfirmSubsetOf(F32Set(0, 1), F32Set(0, 1), true)
	ConfirmSubsetOf(F32Set(0, 1, 2), F32Set(0, 1), false)
}

func TestF32Member(t *testing.T) {
	ConfirmMember := func(s f32set, x float32, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(F32Set(), 0, false)
	ConfirmMember(F32Set(0), 0, true)
	ConfirmMember(F32Set(0, 1), 0, true)
	ConfirmMember(F32Set(0, 1), 1, true)
	ConfirmMember(F32Set(0, 1), 2, false)
}

func TestF32Equal(t *testing.T) {
	ConfirmEqual := func(s, x f32set, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(F32Set(), F32Set(), true)
	ConfirmEqual(F32Set(0), F32Set(), false)
	ConfirmEqual(F32Set(), F32Set(0), false)
	ConfirmEqual(F32Set(0), F32Set(0), true)
	ConfirmEqual(F32Set(0, 0), F32Set(0), true)
	ConfirmEqual(F32Set(0), F32Set(0, 0), true)
	ConfirmEqual(F32Set(0, 1), F32Set(0, 0), false)
	ConfirmEqual(F32Set(0, 1), F32Set(0, 1), true)
	ConfirmEqual(F32Set(0, 1), F32Set(1, 0), true)
	ConfirmEqual(F32Set(0, 1), F32Set(1, 1), false)
}

func TestF32Sum(t *testing.T) {
	ConfirmSum := func(s f32set, r float32) {
		if v := s.Sum(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmSum(F32Set(), 0)
	ConfirmSum(F32Set(0), 0)
	ConfirmSum(F32Set(0, 1), 1)
	ConfirmSum(F32Set(0, 1, 1), 1)
	ConfirmSum(F32Set(0, 1, 2), 3)
}

func TestF32Product(t *testing.T) {
	ConfirmProduct := func(s f32set, r float32) {
		if v := s.Product(); r != v {
			t.Errorf("%v.Sum() expected %v but produced %v", s, r, v)
		}
	}

	ConfirmProduct(F32Set(), 1)
	ConfirmProduct(F32Set(0), 0)
	ConfirmProduct(F32Set(0, 1), 0)
	ConfirmProduct(F32Set(1), 1)
	ConfirmProduct(F32Set(1, 1), 1)
	ConfirmProduct(F32Set(0, 1, 1), 0)
	ConfirmProduct(F32Set(0, 1, 2), 0)
	ConfirmProduct(F32Set(1, 2), 2)
	ConfirmProduct(F32Set(1, 2, 3), 6)
	ConfirmProduct(F32Set(1, 2, 3, 3), 6)
}