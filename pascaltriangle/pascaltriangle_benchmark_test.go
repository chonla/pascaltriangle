package pascaltriangle

import "testing"

func BenchmarkPascalLineMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PascalLineMemo(1000)
	}
}

func BenchmarkPascalLineRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PascalLineRecursive(1000)
	}
}
