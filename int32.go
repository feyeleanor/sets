package sets

import (
	"github.com/feyeleanor/slices"
)

type i32map	map[int32] bool

func (m i32map) Len() int {
	return len(m)
}

func (m i32map) Member(i int32) bool {
	return m[i]
}

func (m i32map) include(v int32) {
	m[v] = true
}

func (m i32map) delete(v int32) {
	delete(m, v)
}

func (m i32map) Each(f interface{}) {
	switch f := f.(type) {
	case func(int32):
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


type i32set struct {
	i32map
}

func I32Set(v... int32) (r i32set) {
	r.i32map = make(i32map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s i32set) String() (t string) {
	elements := slices.I32Slice{}
	s.Each(func(v int32) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s i32set) Intersection(o i32set) (r i32set) {
	r.i32map = make(i32map)
	s.Each(func(v int32) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i32set) Union(o i32set) (r i32set) {
	r.i32map = make(i32map)
	s.Each(func(v int32) {
		r.include(v)
	})
	o.Each(func(v int32) {
		r.include(v)
	})
	return
}

func (s i32set) Difference(o i32set) (r i32set) {
	r.i32map = make(i32map)
	s.Each(func(v int32) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i32set) SubsetOf(o i32set) (r bool) {
	r = true
	s.Each(func(v int32) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s i32set) Equal(o i32set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v int32) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s i32set) Sum() interface{} {
	var r int32
	s.Each(func(v int32) {
		r += v
	})
	return r
}

func (s i32set) Product() interface{} {
	r := int32(1)
	s.Each(func(v int32) {
		r *= v
	})
	return r
}
