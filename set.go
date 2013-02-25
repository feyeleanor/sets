package sets

import (
	"github.com/feyeleanor/slices"
)

type vmap map[interface{}] bool

func (m vmap) Len() int {
	return len(m)
}

func (m vmap) Member(i interface{}) bool {
	return m[i]
}

func (m vmap) include(v interface{}) {
	m[v] = true
}

func (m vmap) delete(v interface{}) {
	delete(m, v)
}

func (m vmap) Each(f interface{}) {
	switch f := f.(type) {
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

func (m vmap) String() (t string) {
	elements := slices.Slice{}
	m.Each(func(v interface{}) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}


type vset struct {
	vmap
}

func Set(v... interface{}) (r vset) {
	r.vmap = make(vmap)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s vset) Intersection(o vset) (r vset) {
	r.vmap = make(vmap)
	s.Each(func(v interface{}) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s vset) Union(o vset) (r vset) {
	r.vmap = make(vmap)
	s.Each(func(v interface{}) {
		r.include(v)
	})
	o.Each(func(v interface{}) {
		r.include(v)
	})
	return
}

func (s vset) Difference(o vset) (r vset) {
	r.vmap = make(vmap)
	s.Each(func(v interface{}) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s vset) SubsetOf(o vset) (r bool) {
	r = true
	s.Each(func(v interface{}) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s vset) Equal(o vset) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v interface{}) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}