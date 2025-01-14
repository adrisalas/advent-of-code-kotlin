package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := "2333133121414131402"
	want := 1928

	got := part1(input)

	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func Test_part2(t *testing.T) {
	input := "2333133121414131402"
	want := 2858

	got := part2(input)

	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}
}
