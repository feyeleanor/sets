package sets

import (
	"github.com/feyeleanor/slices"
)

type c64map	map[complex64] bool

func (m c64map) Len() int {
	return len(m)
}

func (m c64map) Member(i complex64) bool {
	return m[i]
}

func (m c64map) include(v complex64) {
	m[v] = true
}

func (m c64map) delete(v complex64) {
	delete(m, v)
}

func (m c64map) Each(f interface{}) {
	switch f := f.(type) {
	case func(complex64):
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

func (m c64map) String() (t string) {
	elements := slices.C64Slice{}
	m.Each(func(v complex64) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}


type c64set struct {
	c64map
}

func C64Set(v... complex64) (r c64set) {
	r.c64map = make(c64map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s c64set) Intersection(o c64set) (r c64set) {
	r.c64map = make(c64map)
	s.Each(func(v complex64) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s c64set) Union(o c64set) (r c64set) {
	r.c64map = make(c64map)
	s.Each(func(v complex64) {
		r.include(v)
	})
	o.Each(func(v complex64) {
		r.include(v)
	})
	return
}

func (s c64set) Difference(o c64set) (r c64set) {
	r.c64map = make(c64map)
	s.Each(func(v complex64) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s c64set) SubsetOf(o c64set) (r bool) {
	r = true
	s.Each(func(v complex64) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s c64set) Equal(o c64set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v complex64) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s c64set) Sum() interface{} {
	var r complex64
	s.Each(func(v complex64) {
		r += v
	})
	return r
}

func (s c64set) Product() interface{} {
	r := complex64(1)
	s.Each(func(v complex64) {
		r *= v
	})
	return r
}