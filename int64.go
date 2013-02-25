package sets

import (
	"github.com/feyeleanor/slices"
)

type i64map	map[int64] bool

func (m i64map) Len() int {
	return len(m)
}

func (m i64map) Member(i int64) bool {
	return m[i]
}

func (m i64map) include(v int64) {
	m[v] = true
}

func (m i64map) delete(v int64) {
	delete(m, v)
}

func (m i64map) Each(f interface{}) {
	switch f := f.(type) {
	case func(int64):
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


type i64set struct {
	i64map
}

func I64Set(v... int64) (r i64set) {
	r.i64map = make(i64map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s i64set) String() (t string) {
	elements := slices.I64Slice{}
	s.Each(func(v int64) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s i64set) Intersection(o i64set) (r i64set) {
	r.i64map = make(i64map)
	s.Each(func(v int64) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i64set) Union(o i64set) (r i64set) {
	r.i64map = make(i64map)
	s.Each(func(v int64) {
		r.include(v)
	})
	o.Each(func(v int64) {
		r.include(v)
	})
	return
}

func (s i64set) Difference(o i64set) (r i64set) {
	r.i64map = make(i64map)
	s.Each(func(v int64) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s i64set) SubsetOf(o i64set) (r bool) {
	r = true
	s.Each(func(v int64) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s i64set) Equal(o i64set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v int64) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s i64set) Sum() interface{} {
	var r int64
	s.Each(func(v int64) {
		r += v
	})
	return r
}

func (s i64set) Product() interface{} {
	r := int64(1)
	s.Each(func(v int64) {
		r *= v
	})
	return r
}