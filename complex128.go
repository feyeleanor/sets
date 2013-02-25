package sets

import (
	"github.com/feyeleanor/slices"
)

type c128map	map[complex128] bool

func (m c128map) Len() int {
	return len(m)
}

func (m c128map) Member(i complex128) bool {
	return m[i]
}

func (m c128map) include(v complex128) {
	m[v] = true
}

func (m c128map) delete(v complex128) {
	delete(m, v)
}

func (m c128map) Each(f interface{}) {
	switch f := f.(type) {
	case func(complex128):
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

func (m c128map) String() (t string) {
	elements := slices.C128Slice{}
	m.Each(func(v complex128) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}


type c128set struct {
	c128map
}

func C128Set(v... complex128) (r c128set) {
	r.c128map = make(c128map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s c128set) Intersection(o c128set) (r c128set) {
	r.c128map = make(c128map)
	s.Each(func(v complex128) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s c128set) Union(o c128set) (r c128set) {
	r.c128map = make(c128map)
	s.Each(func(v complex128) {
		r.include(v)
	})
	o.Each(func(v complex128) {
		r.include(v)
	})
	return
}

func (s c128set) Difference(o c128set) (r c128set) {
	r.c128map = make(c128map)
	s.Each(func(v complex128) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s c128set) SubsetOf(o c128set) (r bool) {
	r = true
	s.Each(func(v complex128) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s c128set) Equal(o c128set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v complex128) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s c128set) Sum() interface{} {
	var r complex128
	s.Each(func(v complex128) {
		r += v
	})
	return r
}

func (s c128set) Product() interface{} {
	r := complex128(1)
	s.Each(func(v complex128) {
		r *= v
	})
	return r
}