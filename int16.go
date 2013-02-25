package sets

import (
	"github.com/feyeleanor/slices"
)

type i16map	map[int16] bool

func (m i16map) Len() int {
	return len(m)
}

func (m i16map) Member(i int16) bool {
	return m[i]
}

func (m i16map) include(v int16) {
	m[v] = true
}

func (m i16map) delete(v int16) {
	delete(m, v)
}

func (m i16map) Each(f interface{}) {
	switch f := f.(type) {
	case func(int16):
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


type i16set struct {
	i16map
}

func I16Set(v... int16) (r i16set) {
	r.i16map = make(i16map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s i16set) String() (t string) {
	elements := slices.I16Slice{}
	s.Each(func(v int16) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s i16set) Intersection(o i16set) (r i16set) {
	r.i16map = make(i16map)
	s.Each(func(v int16) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i16set) Union(o i16set) (r i16set) {
	r.i16map = make(i16map)
	s.Each(func(v int16) {
		r.include(v)
	})
	o.Each(func(v int16) {
		r.include(v)
	})
	return
}

func (s i16set) Difference(o i16set) (r i16set) {
	r.i16map = make(i16map)
	s.Each(func(v int16) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i16set) SubsetOf(o i16set) (r bool) {
	r = true
	s.Each(func(v int16) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s i16set) Equal(o i16set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v int16) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s i16set) Sum() interface{} {
	var r int16
	s.Each(func(v int16) {
		r += v
	})
	return r
}

func (s i16set) Product() interface{} {
	r := int16(1)
	s.Each(func(v int16) {
		r *= v
	})
	return r
}