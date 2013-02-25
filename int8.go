package sets

import (
	"github.com/feyeleanor/slices"
)

type i8map	map[int8] bool

func (m i8map) Len() int {
	return len(m)
}

func (m i8map) Member(i int8) bool {
	return m[i]
}

func (m i8map) include(v int8) {
	m[v] = true
}

func (m i8map) delete(v int8) {
	delete(m, v)
}

func (m i8map) Each(f interface{}) {
	switch f := f.(type) {
	case func(int8):
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


type i8set struct {
	i8map
}

func I8Set(v... int8) (r i8set) {
	r.i8map = make(i8map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s i8set) String() (t string) {
	elements := slices.I8Slice{}
	s.Each(func(v int8) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s i8set) Intersection(o i8set) (r i8set) {
	r.i8map = make(i8map)
	s.Each(func(v int8) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i8set) Union(o i8set) (r i8set) {
	r.i8map = make(i8map)
	s.Each(func(v int8) {
		r.include(v)
	})
	o.Each(func(v int8) {
		r.include(v)
	})
	return
}

func (s i8set) Difference(o i8set) (r i8set) {
	r.i8map = make(i8map)
	s.Each(func(v int8) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i8set) SubsetOf(o i8set) (r bool) {
	r = true
	s.Each(func(v int8) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s i8set) Equal(o i8set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v int8) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s i8set) Sum() interface{} {
	var r int8
	s.Each(func(v int8) {
		r += v
	})
	return r
}

func (s i8set) Product() interface{} {
	r := int8(1)
	s.Each(func(v int8) {
		r *= v
	})
	return r
}