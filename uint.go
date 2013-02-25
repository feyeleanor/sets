package sets

import (
	"github.com/feyeleanor/slices"
)

type umap	map[uint] bool

func (m umap) Len() int {
	return len(m)
}

func (m umap) Member(i uint) bool {
	return m[i]
}

func (m umap) include(v uint) {
	m[v] = true
}

func (m umap) delete(v uint) {
	delete(m, v)
}

func (m umap) Each(f interface{}) {
	switch f := f.(type) {
	case func(uint):
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


type uset struct {
	umap
}

func USet(v... uint) (r uset) {
	r.umap = make(umap)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s uset) String() (t string) {
	elements := slices.USlice{}
	s.Each(func(v uint) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s uset) Intersection(o uset) (r uset) {
	r.umap = make(umap)
	s.Each(func(v uint) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s uset) Union(o uset) (r uset) {
	r.umap = make(umap)
	s.Each(func(v uint) {
		r.include(v)
	})
	o.Each(func(v uint) {
		r.include(v)
	})
	return
}

func (s uset) Difference(o uset) (r uset) {
	r.umap = make(umap)
	s.Each(func(v uint) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s uset) SubsetOf(o uset) (r bool) {
	r = true
	s.Each(func(v uint) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s uset) Equal(o uset) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v uint) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s uset) Sum() interface{} {
	var r uint
	s.Each(func(v uint) {
		r += v
	})
	return r
}

func (s uset) Product() interface{} {
	r := uint(1)
	s.Each(func(v uint) {
		r *= v
	})
	return r
}