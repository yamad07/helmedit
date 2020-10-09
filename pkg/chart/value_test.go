package chart

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()
	fname := "../../fixtures/values.yaml"
	valf, err := NewValueFile(fname)
	if err != nil {
		t.Error("load values error")
	}
	expected := 3
	actual := valf.Values["replicaCount"]

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("\n(-expected, +actual)\n%s", diff)
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	fname := "../../fixtures/values.yaml"
	valf, err := NewValueFile(fname)
	if err != nil {
		t.Error("load values error")
	}

	expected := "civilization.solution"
	actual := valf.Get("service.annotations.meta")
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("\n(-expected, +actual)\n%s", diff)
	}
}

func TestSet(t *testing.T) {
	t.Parallel()
	fname := "../../fixtures/values.yaml"
	valf, err := NewValueFile(fname)
	if err != nil {
		t.Error("load values error")
	}

	key := "service.annotations.meta"
	expected := "set.civilization.soulutions"

	err = valf.Set(key, expected)
	if err != nil {
		t.Error("load values error")
	}
	actual := valf.Get(key)

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("\n(-expected, +actual)\n%s", diff)
	}
}
