package saying

import "testing"

func TestGreet(t *testing.T) {
	Greet("yash")
}

func BenchmarkGreet(b *testing.B) {
	for i :=0; i<b.N; i++{
		Greet("yash")
	}
}
