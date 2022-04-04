package main

import (
	"flag"
	"fmt"
)

type celsiusFlag struct {
	Celsius
}
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kaerw float64

func (c Celsius) String() string {
	return fmt.Sprintf("%fC", c)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kaerw) Celsius {
	return Celsius(k - 273.15)
}

func (c *celsiusFlag) String() string {
	return fmt.Sprintf("%sC", *c)
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		c.Celsius = Celsius(value)
		return nil
	case "F":
		c.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		c.Celsius = KToC(Kaerw(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
