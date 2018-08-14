package convert

import "fmt"

type Celsius float64
type Fahrenheit float64

type Meter float64
type Feet float64

type Kilogram float64
type Pound float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string       { return fmt.Sprintf("%g f", f) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string      { return fmt.Sprintf("%g lb", p) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MToF converts a Meter length to Feet.
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// FToM converts a Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// KToP converts a Kilogram weight to Pound.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }

// PToK converts a Pound weight to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
