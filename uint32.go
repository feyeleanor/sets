package sets

import (
	"github.com/feyeleanor/slices"
)

type u32map	map[uint32] bool

func (m u32map) Len() int {
	return len(m)
}

func (m u32map) Member(i uint32) bool {
	return m[i]
}

func (m u32map) include(v uint32) {
	m[v] = true
}

func (m u32map) delete(v uint32) {
	delete(m, v)
}

func (m u32map) Each(f interface{}) {
	switch f := f.(type) {
	case func(uint32):
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


type u32set struct {
	u32map
}

func U32Set(v... uint32) (r u32set) {
	r.u32map = make(u32map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s u32set) String() (t string) {
	elements := slices.U32Slice{}
	s.Each(func(v uint32) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s u32set) Intersection(o u32set) (r u32set) {
	r.u32map = make(u32map)
	s.Each(func(v uint32) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u32set) Union(o u32set) (r u32set) {
	r.u32map = make(u32map)
	s.Each(func(v uint32) {
		r.include(v)
	})
	o.Each(func(v uint32) {
		r.include(v)
	})
	return
}

func (s u32set) Difference(o u32set) (r u32set) {
	r.u32map = make(u32map)
	s.Each(func(v uint32) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u32set) SubsetOf(o u32set) (r bool) {
	r = true
	s.Each(func(v uint32) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s u32set) Equal(o u32set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v uint32) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s u32set) Sum() interface{} {
	var r uint32
	s.Each(func(v uint32) {
		r += v
	})
	return r
}

func (s u32set) Product() interface{} {
	r := uint32(1)
	s.Each(func(v uint32) {
		r *= v
	})
	return r
}