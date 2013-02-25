package sets

import (
	"testing"
)

func TestSetString(t *testing.T) {
	ConfirmString := func(s vset, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(Set(), "()")
//	ConfirmString(Set("A", "B"), "(A B)")
//	ConfirmString(Set("B", "A"), "(A B)")
//	ConfirmString(Set("A", "B", "C", "D", "E"), "(A B C D E)")
}

func TestSetIntersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r vset) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(Set(), Set(), Set())
	ConfirmIntersection(Set("A"), Set("B"), Set())
	ConfirmIntersection(Set("A", 1), Set(1, "C"), Set(1))
	ConfirmIntersection(Set("A", 1, 2), Set(1, 2, "D"), Set(1, 2))
}

func TestSetUnion(t *testing.T) {
	ConfirmUnion := func(s, x, r vset) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(Set(), Set(), Set())
	ConfirmUnion(Set("A"), Set(), Set("A"))
	ConfirmUnion(Set(), Set("B"), Set("B"))
	ConfirmUnion(Set("A"), Set("B"), Set("A", "B"))
	ConfirmUnion(Set("A", "B"), Set("B", "C"), Set("A", "B", "C"))
	ConfirmUnion(Set(0, "B", "C"), Set("B", "C", 3), Set(0, "B", "C", 3))
}

func TestSetDifference(t *testing.T) {
	ConfirmDifference := func(s, x, r vset) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(Set(), Set(), Set())
	ConfirmDifference(Set("A"), Set(), Set("A"))
	ConfirmDifference(Set(), Set("B"), Set())
	ConfirmDifference(Set("A"), Set("B"), Set("A"))
	ConfirmDifference(Set("A", "B"), Set("B", "C"), Set("A"))
	ConfirmDifference(Set("A", "B", "C"), Set("B", "C", "D"), Set("A"))
	ConfirmDifference(Set("A", "B", "C", "D"), Set("B", "C", "D"), Set("A"))
}

func TestSetSubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x vset, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(Set(), Set(), true)
	ConfirmSubsetOf(Set("A"), Set(), false)
	ConfirmSubsetOf(Set(), Set("A"), true)
	ConfirmSubsetOf(Set("A"), Set("A"), true)
	ConfirmSubsetOf(Set("A"), Set("B"), false)
	ConfirmSubsetOf(Set("A"), Set("A", "B"), true)
	ConfirmSubsetOf(Set("A", "B"), Set("A", "B"), true)
	ConfirmSubsetOf(Set("A", "B", "C"), Set("A", "B"), false)
}

func TestSetMember(t *testing.T) {
	ConfirmMember := func(s vset, x string, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(Set(), "A", false)
	ConfirmMember(Set("A"), "A", true)
	ConfirmMember(Set("A", "B"), "A", true)
	ConfirmMember(Set("A", "B"), "B", true)
	ConfirmMember(Set("A", "B"), "C", false)
}

func TestSetEqual(t *testing.T) {
	ConfirmEqual := func(s, x vset, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(Set(), Set(), true)
	ConfirmEqual(Set("A"), Set(), false)
	ConfirmEqual(Set(), Set("A"), false)
	ConfirmEqual(Set("A"), Set("A"), true)
	ConfirmEqual(Set("A", "A"), Set("A"), true)
	ConfirmEqual(Set("A"), Set("A", "A"), true)
	ConfirmEqual(Set("A", "B"), Set("A", "A"), false)
	ConfirmEqual(Set("A", "B"), Set("A", "B"), true)
	ConfirmEqual(Set("A", "B"), Set("B", "A"), true)
	ConfirmEqual(Set("A", "B"), Set("B", "B"), false)
}