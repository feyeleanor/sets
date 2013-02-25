package sets

import (
	"testing"
)

func TestSString(t *testing.T) {
	ConfirmString := func(s sset, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(SSet(), "()")
	ConfirmString(SSet("A", "B"), "(A B)")
	ConfirmString(SSet("B", "A"), "(A B)")
	ConfirmString(SSet("A", "B", "C", "D", "E"), "(A B C D E)")
}

func TestSIntersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r sset) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(SSet(), SSet(), SSet())
	ConfirmIntersection(SSet("A"), SSet("B"), SSet())
	ConfirmIntersection(SSet("A", "B"), SSet("B", "C"), SSet("B"))
	ConfirmIntersection(SSet("A", "B", "C"), SSet("B", "C", "D"), SSet("B", "C"))
}

func TestSUnion(t *testing.T) {
	ConfirmUnion := func(s, x, r sset) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(SSet(), SSet(), SSet())
	ConfirmUnion(SSet("A"), SSet(), SSet("A"))
	ConfirmUnion(SSet(), SSet("B"), SSet("B"))
	ConfirmUnion(SSet("A"), SSet("B"), SSet("A", "B"))
	ConfirmUnion(SSet("A", "B"), SSet("B", "C"), SSet("A", "B", "C"))
	ConfirmUnion(SSet("A", "B", "C"), SSet("B", "C", "D"), SSet("A", "B", "C", "D"))
}

func TestSDifference(t *testing.T) {
	ConfirmDifference := func(s, x, r sset) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(SSet(), SSet(), SSet())
	ConfirmDifference(SSet("A"), SSet(), SSet("A"))
	ConfirmDifference(SSet(), SSet("B"), SSet())
	ConfirmDifference(SSet("A"), SSet("B"), SSet("A"))
	ConfirmDifference(SSet("A", "B"), SSet("B", "C"), SSet("A"))
	ConfirmDifference(SSet("A", "B", "C"), SSet("B", "C", "D"), SSet("A"))
	ConfirmDifference(SSet("A", "B", "C", "D"), SSet("B", "C", "D"), SSet("A"))
}

func TestSSubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x sset, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(SSet(), SSet(), true)
	ConfirmSubsetOf(SSet("A"), SSet(), false)
	ConfirmSubsetOf(SSet(), SSet("A"), true)
	ConfirmSubsetOf(SSet("A"), SSet("A"), true)
	ConfirmSubsetOf(SSet("A"), SSet("B"), false)
	ConfirmSubsetOf(SSet("A"), SSet("A", "B"), true)
	ConfirmSubsetOf(SSet("A", "B"), SSet("A", "B"), true)
	ConfirmSubsetOf(SSet("A", "B", "C"), SSet("A", "B"), false)
}

func TestSMember(t *testing.T) {
	ConfirmMember := func(s sset, x string, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(SSet(), "A", false)
	ConfirmMember(SSet("A"), "A", true)
	ConfirmMember(SSet("A", "B"), "A", true)
	ConfirmMember(SSet("A", "B"), "B", true)
	ConfirmMember(SSet("A", "B"), "C", false)
}

func TestSEqual(t *testing.T) {
	ConfirmEqual := func(s, x sset, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(SSet(), SSet(), true)
	ConfirmEqual(SSet("A"), SSet(), false)
	ConfirmEqual(SSet(), SSet("A"), false)
	ConfirmEqual(SSet("A"), SSet("A"), true)
	ConfirmEqual(SSet("A", "A"), SSet("A"), true)
	ConfirmEqual(SSet("A"), SSet("A", "A"), true)
	ConfirmEqual(SSet("A", "B"), SSet("A", "A"), false)
	ConfirmEqual(SSet("A", "B"), SSet("A", "B"), true)
	ConfirmEqual(SSet("A", "B"), SSet("B", "A"), true)
	ConfirmEqual(SSet("A", "B"), SSet("B", "B"), false)
}