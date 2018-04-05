package set

import (
	"fmt"
	"sort"
)

// Set -
type Set interface {
	Have(...interface{}) bool
	Subset(Set) bool
	Superset(Set) bool
	Add(...interface{})
	Remove(...interface{})
	Clear()
	Or(Set) Set
	And(Set) Set
	Xor(Set) Set
	Copy() Set
	Keys() []string
}

type set struct {
	Values map[interface{}]struct{}
}

func (s *set) Have(keys ...interface{}) bool {
	for _, key := range keys {
		if _, ok := s.Values[key]; !ok {
			return false
		}
	}
	return true
}

func (s *set) Subset(t Set) bool {
	keys := []interface{}{}
	for key := range s.Values {
		keys = append(keys, key)
	}
	return t.Have(keys...)
}

func (s *set) Superset(t Set) bool {
	return t.Subset(s)
}

func (s *set) Add(keys ...interface{}) {
	for _, key := range keys {
		s.Values[key] = struct{}{}
	}
}

func (s *set) Remove(keys ...interface{}) {
	for _, key := range keys {
		delete(s.Values, key)
	}
}

func (s *set) Clear() {
	s.Values = map[interface{}]struct{}{}
}

func (s *set) Or(t Set) Set {
	r := t.Copy()
	for key := range s.Values {
		r.Add(key)
	}
	return r
}

func (s *set) And(t Set) Set {
	r := &set{}
	r.Clear()
	for key := range s.Values {
		if t.Have(key) {
			r.Values[key] = struct{}{}
		}
	}
	return r
}

func (s *set) Xor(t Set) Set {
	r := t.Copy()
	for key := range s.Values {
		if t.Have(key) {
			r.Remove(key)
		} else {
			r.Add(key)
		}
	}
	return r
}

func (s *set) Copy() Set {
	r := &set{}
	r.Clear()
	for key := range s.Values {
		r.Values[key] = struct{}{}
	}
	return r
}

func (s *set) Keys() []string {
	r := []string{}
	for key := range s.Values {
		r = append(r, fmt.Sprint(key))
	}
	sort.Strings(r)
	return r
}

// New -
func New(keys ...interface{}) Set {
	r := &set{}
	r.Clear()
	r.Add(keys...)
	return r
}
