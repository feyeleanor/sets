package sets

import (
	"github.com/feyeleanor/slices"
)

type imap	map[int] bool

func (m imap) Len() int {
	return len(m)
}

func (m imap) Member(i int) bool {
	return m[i]
}

func (m imap) include(v int) {
	m[v] = true
}

func (m imap) delete(v int) {
	delete(m, v)
}

func (m imap) Each(f interface{}) {
	switch f := f.(type) {
	case func(int):
		for k, v := range m {
			if v {
				f(k)
			}
		}
	case func(interface{}):
		for k, v := range m {
			if v {
				f(k)
			}
		}
	default:
		panic(f)
	}
}


type iset struct {
	imap
}

func ISet(v... int) (r iset) {
	r.imap = make(imap)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s iset) String() (t string) {
	elements := slices.ISlice{}
	s.Each(func(v int) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s iset) Intersection(o iset) (r iset) {
	r.imap = make(imap)
	s.Each(func(v int) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s iset) Union(o iset) (r iset) {
	r.imap = make(imap)
	s.Each(func(v int) {
		r.include(v)
	})
	o.Each(func(v int) {
		r.include(v)
	})
	return
}

func (s iset) Difference(o iset) (r iset) {
	r.imap = make(imap)
	s.Each(func(v int) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s iset) SubsetOf(o iset) (r bool) {
	r = true
	s.Each(func(v int) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s iset) Equal(o iset) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v int) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s iset) Sum() interface{} {
	var r int
	s.Each(func(v int) {
		r += v
	})
	return r
}

func (s iset) Product() interface{} {
	r := int(1)
	s.Each(func(v int) {
		r *= v
	})
	return r
}