package sets

import (
	"fmt"
	"testing"
)

type Errno int

func (e Errno) Error() (err string) {
	return fmt.Sprintf("(E%v)", int(e))
}

const (
	E0 = Errno(iota)
	E1
	E2
	E3
	E4
)

func TestEString(t *testing.T) {
	ConfirmString := func(s eset, r string) {
		if v := s.String(); r != v {
			t.Errorf("String() expected %v but produced %v", r, v)
		}
	}

	ConfirmString(ESet(), "()")
//	ConfirmString(ESet(E0, E1), "(E0 E1)")
//	ConfirmString(ESet(E1, E0), "(E0 E1)")
//	ConfirmString(ESet(E0, E1, E2, E3, E4), "(E0 E1 E2 E3 E4)")
}

func TestEIntersection(t *testing.T) {
	ConfirmIntersection := func(s, x, r eset) {
		if v := s.Intersection(x); !r.Equal(v) {
			t.Errorf("%v.Intersection(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmIntersection(ESet(), ESet(), ESet())
	ConfirmIntersection(ESet(E0), ESet(E1), ESet())
	ConfirmIntersection(ESet(E0, E1), ESet(E1, E2), ESet(E1))
	ConfirmIntersection(ESet(E0, E1, E2), ESet(E1, E2, E3), ESet(E1, E2))
}

func TestEUnion(t *testing.T) {
	ConfirmUnion := func(s, x, r eset) {
		if v := s.Union(x); !r.Equal(v) {
			t.Errorf("%v.Union(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmUnion(ESet(), ESet(), ESet())
	ConfirmUnion(ESet(E0), ESet(), ESet(E0))
	ConfirmUnion(ESet(), ESet(E1), ESet(E1))
	ConfirmUnion(ESet(E0), ESet(E1), ESet(E0, E1))
	ConfirmUnion(ESet(E0, E1), ESet(E1, E2), ESet(E0, E1, E2))
	ConfirmUnion(ESet(E0, E1, E2), ESet(E1, E2, E3), ESet(E0, E1, E2, E3))
}

func TestEDifference(t *testing.T) {
	ConfirmDifference := func(s, x, r eset) {
		if v := s.Difference(x); !r.Equal(v) {
			t.Errorf("%v.Difference(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmDifference(ESet(), ESet(), ESet())
	ConfirmDifference(ESet(E0), ESet(), ESet(E0))
	ConfirmDifference(ESet(), ESet(E1), ESet())
	ConfirmDifference(ESet(E0), ESet(E1), ESet(E0))
	ConfirmDifference(ESet(E0, E1), ESet(E1, E2), ESet(E0))
	ConfirmDifference(ESet(E0, E1, E2), ESet(E1, E2, E3), ESet(E0))
	ConfirmDifference(ESet(E0, E1, E2, E3), ESet(E1, E2, E3), ESet(E0))
}

func TestESubsetOf(t *testing.T) {
	ConfirmSubsetOf := func(s, x eset, r bool) {
		if v := s.SubsetOf(x); r != v {
			t.Errorf("%v.SubsetOf(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmSubsetOf(ESet(), ESet(), true)
	ConfirmSubsetOf(ESet(E0), ESet(), false)
	ConfirmSubsetOf(ESet(), ESet(E0), true)
	ConfirmSubsetOf(ESet(E0), ESet(E0), true)
	ConfirmSubsetOf(ESet(E0), ESet(E1), false)
	ConfirmSubsetOf(ESet(E0), ESet(E0, E1), true)
	ConfirmSubsetOf(ESet(E0, E1), ESet(E0, E1), true)
	ConfirmSubsetOf(ESet(E0, E1, E2), ESet(E0, E1), false)
}

func TestEMember(t *testing.T) {
	ConfirmMember := func(s eset, x error, r bool) {
		if v := s.Member(x); r != v {
			t.Errorf("%v.Member(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmMember(ESet(), E0, false)
	ConfirmMember(ESet(E0), E0, true)
	ConfirmMember(ESet(E0, E1), E0, true)
	ConfirmMember(ESet(E0, E1), E1, true)
	ConfirmMember(ESet(E0, E1), E2, false)
}

func TestEEqual(t *testing.T) {
	ConfirmEqual := func(s, x eset, r bool) {
		if v := s.Equal(x); r != v {
			t.Errorf("%v.Equal(%v) expected %v but produced %v", s, x, r, v)
		}
	}

	ConfirmEqual(ESet(), ESet(), true)
	ConfirmEqual(ESet(E0), ESet(), false)
	ConfirmEqual(ESet(), ESet(E0), false)
	ConfirmEqual(ESet(E0), ESet(E0), true)
	ConfirmEqual(ESet(E0, E0), ESet(E0), true)
	ConfirmEqual(ESet(E0), ESet(E0, E0), true)
	ConfirmEqual(ESet(E0, E1), ESet(E0, E0), false)
	ConfirmEqual(ESet(E0, E1), ESet(E0, E1), true)
	ConfirmEqual(ESet(E0, E1), ESet(E1, E0), true)
	ConfirmEqual(ESet(E0, E1), ESet(E1, E1), false)
}