package loader

import (
	"testing"
)

var testFiles = []string{
	"../data/file1.csv",
	"../data/file2.csv",
}

func BenchmarkSearchParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchTxnParallel(testFiles, "TXN100128")
	}
}
