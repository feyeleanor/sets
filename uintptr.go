package sets

import (
	"github.com/feyeleanor/slices"
)

type amap	map[uintptr] bool

func (m amap) Len() int {
	return len(m)
}

func (m amap) Member(i uintptr) bool {
	return m[i]
}

func (m amap) include(v uintptr) {
	m[v] = true
}

func (m amap) delete(v uintptr) {
	delete(m, v)
}

func (m amap) Each(f interface{}) {
	switch f := f.(type) {
	case func(uintptr):
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


type aset struct {
	amap
}

func ASet(v... uintptr) (r aset) {
	r.amap = make(amap)
	for i := len(v) - 1; i > -1; i-- {
		r.include(v[i])
	}
	return
}

func (s aset) String() (t string) {
	elements := slices.ASlice{}
	s.Each(func(v uintptr) {
		elements = append(elements, v)
	})
	slices.Sort(elements)
	return elements.String()
}

func (s aset) Intersection(o aset) (r aset) {
	r.amap = make(amap)
	s.Each(func(v uintptr) {
		if o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s aset) Union(o aset) (r aset) {
	r.amap = make(amap)
	s.Each(func(v uintptr) {
		r.include(v)
	})
	o.Each(func(v uintptr) {
		r.include(v)
	})
	return
}

func (s aset) Difference(o aset) (r aset) {
	r.amap = make(amap)
	s.Each(func(v uintptr) {
		if !o.Member(v) {
			r.include(v)
		}
	})
	return
}

func (s aset) SubsetOf(o aset) (r bool) {
	r = true
	s.Each(func(v uintptr) {
		if !o.Member(v) {
			r = false
			return
		}
	})
	return
}

func (s aset) Equal(o aset) (r bool) {
	if r = s.Len() == o.Len(); r {
		s.Each(func(v uintptr) {
			if !o.Member(v) {
				r = false
			}
		})
	}
	return
}

func (s aset) Sum() interface{} {
	var r uintptr
	s.Each(func(v uintptr) {
		r += v
	})
	return r
}

func (s aset) Product() interface{} {
	r := uintptr(1)
	s.Each(func(v uintptr) {
		r *= v
	})
	return r
}