package bar

import (
	"io"
	"os"
	"testing"
)

func ExampleGraph() {
	Graph(os.Stdout, 3, map[string]float64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
	})
	// Output:
	// 8 █
	// 7 ▉
	// 6 ▊
	// 5 ▋
	// 4 ▌
	// 3 ▍
	// 2 ▎
	// 1 ▏
	// 0
}

func BenchmarkGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Graph(io.Discard, 80, map[string]float64{
			"":    0.9999999999999,
			"c":   0.99,
			"b":   0.98,
			"d":   0.01,
			"de":  0.005,
			"dfe": 0.001,
		})
	}
}
