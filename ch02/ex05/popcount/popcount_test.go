package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	count := PopCount(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCount(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountLoop(t *testing.T) {
	count := PopCountUsingLoop(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountUsingLoop(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountShift64(t *testing.T) {
	count := PopCountShift64(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountShift64(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

var input uint64 = 0x1234567890ABCDEF
var results []int

func BenchmarkPopCount(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCount(input))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountUsingLoop(input))
	}
}

func BenchmarkPopCountShift64(b *testing.B) {
	results = []int{}
	for i := 0; i < b.N; i++ {
		results = append(results, PopCountShift64(input))
	}
}
