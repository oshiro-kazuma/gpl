package main

import (
	"fmt"
	"github.com/oshiro-kazuma/gpl/ch02/ex01/tempconv"
	"github.com/yutachaos/programming-go-study2019/ch02/ex02/distconv"
	"os"
	"strconv"
)

func main() {
	v := os.Args[1]
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convert error: %v\n", err)
		os.Exit(1)
	}

	// Temperature
	f1 := tempconv.Fahrenheit(value)
	c1 := tempconv.Celsius(value)
	fmt.Printf("Fahrenheit %s = Celsius %s \n", f1, tempconv.FToC(f1))
	fmt.Printf("Celsius %s = Fahrenheit %s \n", c1, tempconv.CToF(c1))

	// Distance
	f2 := distconv.Feet(value)
	m2 := distconv.Meter(value)
	fmt.Printf("Feet %s = Meter %s \n", f2, distconv.FToM(f2))
	fmt.Printf("Meter %s = Feet %s \n", m2, distconv.MToF(m2))
}
