package set_test

import (
	"reflect"
	"testing"

	"github.com/maxbet1507/set"
)

func TestNew(t *testing.T) {
	s1 := set.New(1, 2, 3)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2", "3"}) {
		t.Fatal(s1)
	}
}

func TestIn(t *testing.T) {
	s1 := set.New(1, 2, 3)

	if !s1.In(1) || !s1.In(2) || !s1.In(3) {
		t.Fatal(s1)
	}
	if !s1.In(1, 2) || !s1.In(2, 3) || !s1.In(1, 3) {
		t.Fatal(s1)
	}
	if !s1.In(1, 2, 3) || s1.In(1, 2, 4) {
		t.Fatal(s1)
	}
}

func TestSubset(t *testing.T) {
	s1 := set.New(1, 2, 3)

	if !s1.Subset(set.New(1, 2, 3, 4)) || !s1.Subset(set.New(1, 2, 3)) {
		t.Fatal(s1)
	}
	if s1.Subset(set.New(2, 3, 4)) {
		t.Fatal(s1)
	}
}

func TestSuperset(t *testing.T) {
	s1 := set.New(1, 2, 3)

	if !s1.Superset(set.New(1, 2)) || !s1.Superset(set.New(2, 3)) || !s1.Superset(set.New(1, 3)) {
		t.Fatal(s1)
	}
	if s1.Superset(set.New(1, 2, 3, 4)) {
		t.Fatal(s1)
	}
}

func TestAdd(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s1.Add(3, 4)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2", "3", "4"}) {
		t.Fatal(s1)
	}
}

func TestRemove(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s1.Remove(3, 4)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2"}) {
		t.Fatal(s1)
	}
}

func TestClear(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s1.Clear()

	if !reflect.DeepEqual(s1.Keys(), []string{}) {
		t.Fatal(s1)
	}
}

func TestOr(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(2, 3, 4)
	s3 := s1.Or(s2)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2", "3"}) {
		t.Fatal(s1)
	}
	if !reflect.DeepEqual(s2.Keys(), []string{"2", "3", "4"}) {
		t.Fatal(s2)
	}
	if !reflect.DeepEqual(s3.Keys(), []string{"1", "2", "3", "4"}) {
		t.Fatal(s3)
	}
}

func TestAnd(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(2, 3, 4)
	s3 := s1.And(s2)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2", "3"}) {
		t.Fatal(s1)
	}
	if !reflect.DeepEqual(s2.Keys(), []string{"2", "3", "4"}) {
		t.Fatal(s2)
	}
	if !reflect.DeepEqual(s3.Keys(), []string{"2", "3"}) {
		t.Fatal(s3)
	}
}

func TestXor(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(2, 3, 4)
	s3 := s1.Xor(s2)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2", "3"}) {
		t.Fatal(s1)
	}
	if !reflect.DeepEqual(s2.Keys(), []string{"2", "3", "4"}) {
		t.Fatal(s2)
	}
	if !reflect.DeepEqual(s3.Keys(), []string{"1", "4"}) {
		t.Fatal(s3)
	}
}

func TestCopy(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := s1.Copy()
	s2.Add(4)
	s2.Remove(1)

	if !reflect.DeepEqual(s1.Keys(), []string{"1", "2", "3"}) {
		t.Fatal(s1)
	}
	if !reflect.DeepEqual(s2.Keys(), []string{"2", "3", "4"}) {
		t.Fatal(s2)
	}
}
