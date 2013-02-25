package sets

import (
	"github.com/feyeleanor/slices"
)

type f32map	map[float32] bool

func (m f32map) Len() int {
	return len(m)
}

func (m f32map) Member(i float32) bool {
	return m[i]
}

func (m f32map) include(v float32) {
	m[v] = true
}

func (m f32map) delete(v float32) {
	delete(m, v)
}

func (m f32map) Each(f interface{}) {
	switch f := f.(type) {
	case func(float32):
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


type f32set struct {
	f32map
}

func F32Set(v... float32) (r f32set) {
	r.f32map = make(f32map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s f32set) String() (t string) {
	elements := slices.F32Slice{}
	s.Each(func(v float32) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s f32set) Intersection(o f32set) (r f32set) {
	r.f32map = make(f32map)
	s.Each(func(v float32) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s f32set) Union(o f32set) (r f32set) {
	r.f32map = make(f32map)
	s.Each(func(v float32) {
		r.include(v)
	})
	o.Each(func(v float32) {
		r.include(v)
	})
	return
}

func (s f32set) Difference(o f32set) (r f32set) {
	r.f32map = make(f32map)
	s.Each(func(v float32) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s f32set) SubsetOf(o f32set) (r bool) {
	r = true
	s.Each(func(v float32) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s f32set) Equal(o f32set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v float32) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s f32set) Sum() interface{} {
	var r float32
	s.Each(func(v float32) {
		r += v
	})
	return r
}

func (s f32set) Product() interface{} {
	r := float32(1)
	s.Each(func(v float32) {
		r *= v
	})
	return r
}