package sets

import (
	"github.com/feyeleanor/slices"
)

type u64map	map[uint64] bool

func (m u64map) Len() int {
	return len(m)
}

func (m u64map) Member(i uint64) bool {
	return m[i]
}

func (m u64map) include(v uint64) {
	m[v] = true
}

func (m u64map) delete(v uint64) {
	delete(m, v)
}

func (m u64map) Each(f interface{}) {
	switch f := f.(type) {
	case func(uint64):
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


type u64set struct {
	u64map
}

func U64Set(v... uint64) (r u64set) {
	r.u64map = make(u64map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s u64set) String() (t string) {
	elements := slices.U64Slice{}
	s.Each(func(v uint64) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s u64set) Intersection(o u64set) (r u64set) {
	r.u64map = make(u64map)
	s.Each(func(v uint64) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u64set) Union(o u64set) (r u64set) {
	r.u64map = make(u64map)
	s.Each(func(v uint64) {
		r.include(v)
	})
	o.Each(func(v uint64) {
		r.include(v)
	})
	return
}

func (s u64set) Difference(o u64set) (r u64set) {
	r.u64map = make(u64map)
	s.Each(func(v uint64) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s u64set) SubsetOf(o u64set) (r bool) {
	r = true
	s.Each(func(v uint64) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s u64set) Equal(o u64set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v uint64) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s u64set) Sum() interface{} {
	var r uint64
	s.Each(func(v uint64) {
		r += v
	})
	return r
}

func (s u64set) Product() interface{} {
	r := uint64(1)
	s.Each(func(v uint64) {
		r *= v
	})
	return r
}