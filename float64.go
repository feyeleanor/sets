package sets

import (
	"github.com/feyeleanor/slices"
)

type f64map	map[float64] bool

func (m f64map) Len() int {
	return len(m)
}

func (m f64map) Member(i float64) bool {
	return m[i]
}

func (m f64map) include(v float64) {
	m[v] = true
}

func (m f64map) delete(v float64) {
	delete(m, v)
}

func (m f64map) Each(f interface{}) {
	switch f := f.(type) {
	case func(float64):
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


type f64set struct {
	f64map
}

func F64Set(v... float64) (r f64set) {
	r.f64map = make(f64map)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s f64set) String() (t string) {
	elements := slices.F64Slice{}
	s.Each(func(v float64) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s f64set) Intersection(o f64set) (r f64set) {
	r.f64map = make(f64map)
	s.Each(func(v float64) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s f64set) Union(o f64set) (r f64set) {
	r.f64map = make(f64map)
	s.Each(func(v float64) {
		r.include(v)
	})
	o.Each(func(v float64) {
		r.include(v)
	})
	return
}

func (s f64set) Difference(o f64set) (r f64set) {
	r.f64map = make(f64map)
	s.Each(func(v float64) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s f64set) SubsetOf(o f64set) (r bool) {
	r = true
	s.Each(func(v float64) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s f64set) Equal(o f64set) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v float64) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s f64set) Sum() interface{} {
	var r float64
	s.Each(func(v float64) {
		r += v
	})
	return r
}

func (s f64set) Product() interface{} {
	r := float64(1)
	s.Each(func(v float64) {
		r *= v
	})
	return r
}