package sets

import (
	"github.com/feyeleanor/slices"
)

type smap	map[string] bool

func (m smap) Len() int {
	return len(m)
}

func (m smap) Member(s string) bool {
	return m[s]
}

func (m smap) include(v string) {
	m[v] = true
}

func (m smap) delete(v string) {
	delete(m, v)
}

func (m smap) Each(f interface{}) {
	switch f := f.(type) {
	case func(string):
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


type sset struct {
	smap
}

func SSet(v... string) (r sset) {
	r.smap = make(smap)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s sset) String() (t string) {
	elements := slices.SSlice{}
	s.Each(func(v string) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s sset) Intersection(o sset) (r sset) {
	r.smap = make(smap)
	s.Each(func(v string) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s sset) Union(o sset) (r sset) {
	r.smap = make(smap)
	s.Each(func(v string) {
		r.include(v)
	})
	o.Each(func(v string) {
		r.include(v)
	})
	return
}

func (s sset) Difference(o sset) (r sset) {
	r.smap = make(smap)
	s.Each(func(v string) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s sset) SubsetOf(o sset) (r bool) {
	r = true
	s.Each(func(v string) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s sset) Equal(o sset) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v string) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}