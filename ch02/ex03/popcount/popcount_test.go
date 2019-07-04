package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	count := PopCount(1)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountUsingLoop(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountUsingLoop(t *testing.T) {
	count := PopCountUsingLoop(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountUsingLoop(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

var results []int

func BenchmarkPopCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCount(uint64(i)))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountUsingLoop(uint64(i)))
	}
}
