package sets

import (
	"github.com/feyeleanor/slices"
)

type emap	map[error] bool

func (m emap) Len() int {
	return len(m)
}

func (m emap) Member(i error) bool {
	return m[i]
}

func (m emap) include(v error) {
	m[v] = true
}

func (m emap) delete(v error) {
	delete(m, v)
}

func (m emap) Each(f interface{}) {
	switch f := f.(type) {
	case func(error):
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

func (m emap) String() (t string) {
	elements := slices.ESlice{}
	m.Each(func(v error) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}


type eset struct {
	emap
}

func ESet(v... error) (r eset) {
	r.emap = make(emap)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s eset) Intersection(o eset) (r eset) {
	r.emap = make(emap)
	s.Each(func(v error) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s eset) Union(o eset) (r eset) {
	r.emap = make(emap)
	s.Each(func(v error) {
		r.include(v)
	})
	o.Each(func(v error) {
		r.include(v)
	})
	return
}

func (s eset) Difference(o eset) (r eset) {
	r.emap = make(emap)
	s.Each(func(v error) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s eset) SubsetOf(o eset) (r bool) {
	r = true
	s.Each(func(v error) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s eset) Equal(o eset) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v error) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}