package fractals

import (
	"testing"
)

func BenchmarkMandelbrot64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot64(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrot128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot128(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotBigFloat(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotRat(complex(float64(i), float64(i)))
	}
}
