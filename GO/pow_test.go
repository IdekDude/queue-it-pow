package queueit

import "testing"

const (
	input     = "af085a5f-ae56-4450-8bf8-11cabf2b140a"
	zeroCount = 3
	runs      = 25
)

func TestGetHash(t *testing.T) {
	solution, err := GetHash(input, zeroCount, runs)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(solution)
	}
}

func BenchmarkGetHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := GetHash(input, zeroCount, runs); err != nil {
			b.Error(err)
		}
	}
}
