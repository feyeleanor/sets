package sets

import (
	"github.com/feyeleanor/slices"
)

type u16map	map[uint16] bool

func (m u16map) Len() int {
	return len(m)
}

func (m u16map) Member(i uint16) bool {
	return m[i]
}

func (m u16map) include(v uint16) {
	m[v] = true
}

func (m u16map) delete(v uint16) {
	delete(m, v)
}

func (m u16map) Each(f interface{}) {
	switch f := f.(type) {
	case func(uint16):
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


type u16set struct {
	u16map
}

func U16Set(v... uint16) (r u16set) {
	r.u16map = make(u16map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s u16set) String() (t string) {
	elements := slices.U16Slice{}
	s.Each(func(v uint16) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s u16set) Intersection(o u16set) (r u16set) {
	r.u16map = make(u16map)
	s.Each(func(v uint16) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u16set) Union(o u16set) (r u16set) {
	r.u16map = make(u16map)
	s.Each(func(v uint16) {
		r.include(v)
	})
	o.Each(func(v uint16) {
		r.include(v)
	})
	return
}

func (s u16set) Difference(o u16set) (r u16set) {
	r.u16map = make(u16map)
	s.Each(func(v uint16) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u16set) SubsetOf(o u16set) (r bool) {
	r = true
	s.Each(func(v uint16) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s u16set) Equal(o u16set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v uint16) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s u16set) Sum() interface{} {
	var r uint16
	s.Each(func(v uint16) {
		r += v
	})
	return r
}

func (s u16set) Product() interface{} {
	r := uint16(1)
	s.Each(func(v uint16) {
		r *= v
	})
	return r
}