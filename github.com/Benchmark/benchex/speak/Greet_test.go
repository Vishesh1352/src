package speak

import (
	"testing"

	"example.com/src/github.com/Benchmark/benchex/speak"

)
sssssssss
func TestGreet(t *testing.T) {
	speak.Greet("yash")
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		speak.Greet("yash")
	}
}
