// Berkay Deniz
// Exercise 2.1
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64 // Added for Kelvin temperature type

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// This function converts Kelvin temperature to Celcius
func KToC(k Kelvin) Celsius { return Celsius(k - Kelvin(AbsoluteZeroC)) }

// This function converts Celcius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }
