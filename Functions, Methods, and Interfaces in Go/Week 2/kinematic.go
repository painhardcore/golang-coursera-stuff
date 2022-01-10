package main

import (
	"fmt"
)

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
	return fn
}

func main() {
	var err error
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Printf("enter acceleration: ")
	_, err = fmt.Scan(&acceleration)
	if err != nil {
		fmt.Printf("failed to scan input %s*\n", err)
	}
	fmt.Printf("enter initial velocity: ")
	_, err = fmt.Scan(&initialVelocity)
	if err != nil {
		fmt.Printf("failed to scan input %s*\n", err)
	}
	fmt.Printf("enter initial displacement: ")
	_, err = fmt.Scan(&initialDisplacement)
	if err != nil {
		fmt.Printf("failed to scan input %s*\n", err)
	}

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Enter time: ")
	_, err = fmt.Scan(&time)
	if err != nil {
		fmt.Printf("failed to scan input %s*\n", err)
	}
	fmt.Printf("displacement: %f \n", fn(time))

}
