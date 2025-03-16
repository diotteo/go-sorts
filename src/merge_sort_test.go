package main

import (
		"testing"
		"reflect"
		)

func TestAlreadySorted(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10}

	got := merge_sort(arr)
	want := []int{1,2,3,4,5,6,7,8,9,10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestReversed(t *testing.T) {
	arr := []int{10,9,8,7,6,5,4,3,2,1}

	got := merge_sort(arr)
	want := []int{1,2,3,4,5,6,7,8,9,10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRandom(t *testing.T) {
	arr := []int{6,4,7,6,7,7,0,6,5,3}

	got := merge_sort(arr)
	want := []int{0,3,4,5,6,6,6,7,7,7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
