package fractals

import (
	"image/color"
	"math/big"
	"math/cmplx"
)

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	realZ := big.NewFloat(real(z))
	imagZ := big.NewFloat(imag(z))
	realV := new(big.Float)
	imagV := new(big.Float)
	for n := uint8(0); n < iterations; n++ {
		tempR := new(big.Float)
		tempI := new(big.Float)
		tempR.Mul(realV, realV).Sub(tempR, (&big.Float{}).Mul(imagV, imagV)).Add(tempR, realZ)
		tempI.Mul(realV, imagV).Mul(tempI, big.NewFloat(2)).Add(tempI, imagZ)
		realV, imagV = tempR, tempI
		sum := new(big.Float)
		sum.Mul(realV, realV).Add(sum, (&big.Float{}).Mul(imagV, imagV))
		if sum.Cmp(big.NewFloat(4)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	realZ := new(big.Rat).SetFloat64(real(z))
	imagZ := new(big.Rat).SetFloat64(imag(z))
	realV := new(big.Rat)
	imagV := new(big.Rat)
	for n := uint8(0); n < iterations; n++ {
		tempR := new(big.Rat)
		tempI := new(big.Rat)
		tempR.Mul(realV, realV).Sub(tempR, (&big.Rat{}).Mul(imagV, imagV)).Add(tempR, realZ)
		tempI.Mul(realV, imagV).Mul(tempI, big.NewRat(2, 1)).Add(tempI, imagZ)
		realV, imagV = tempR, tempI
		sum := new(big.Rat)
		sum.Mul(realV, realV).Add(sum, (&big.Rat{}).Mul(imagV, imagV))
		if sum.Cmp(big.NewRat(4, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
