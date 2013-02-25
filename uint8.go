package sets

import (
	"github.com/feyeleanor/slices"
)

type u8map	map[uint8] bool

func (m u8map) Len() int {
	return len(m)
}

func (m u8map) Member(i uint8) bool {
	return m[i]
}

func (m u8map) include(v uint8) {
	m[v] = true
}

func (m u8map) delete(v uint8) {
	delete(m, v)
}

func (m u8map) Each(f interface{}) {
	switch f := f.(type) {
	case func(uint8):
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


type u8set struct {
	u8map
}

func U8Set(v... uint8) (r u8set) {
	r.u8map = make(u8map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s u8set) String() (t string) {
	elements := slices.U8Slice{}
	s.Each(func(v uint8) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s u8set) Intersection(o u8set) (r u8set) {
	r.u8map = make(u8map)
	s.Each(func(v uint8) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u8set) Union(o u8set) (r u8set) {
	r.u8map = make(u8map)
	s.Each(func(v uint8) {
		r.include(v)
	})
	o.Each(func(v uint8) {
		r.include(v)
	})
	return
}

func (s u8set) Difference(o u8set) (r u8set) {
	r.u8map = make(u8map)
	s.Each(func(v uint8) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u8set) SubsetOf(o u8set) (r bool) {
	r = true
	s.Each(func(v uint8) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s u8set) Equal(o u8set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v uint8) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s u8set) Sum() interface{} {
	var r uint8
	s.Each(func(v uint8) {
		r += v
	})
	return r
}

func (s u8set) Product() interface{} {
	r := uint8(1)
	s.Each(func(v uint8) {
		r *= v
	})
	return r
}