package main

import (
	"fmt"
	"math"
)

type InitialDisplacement struct {
	acceleration float64
	velocity float64
	displacement float64
}

func RaiseError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetQuestionInput(subject string) float64 {
	var f float64

	fmt.Printf("Enter %s value: ", subject)
	_, err := fmt.Scanf("%f", &f)
	RaiseError(err)

	return f
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	/*
		s = ½ a t^2 + vot + so

		a: acceleration
		vo: initial velocity
		so: initial displacement
	*/

	fn := func(t float64) float64 {
		/*
			t: time
		*/

		return .5 * a * math.Pow(t, 2) + vo * t + so
	}

	return fn
}


func main() {
	id := InitialDisplacement{
		acceleration: GetQuestionInput("acceleration"),
		velocity: GetQuestionInput("initial velocity"),
		displacement: GetQuestionInput("initial displacement"),
	}
	// fmt.Printf("acceleration: %f, velocity: %f, displacement: %f\n", id.acceleration, id.velocity, id.displacement)
	fn := GenDisplaceFn(id.acceleration, id.velocity, id.displacement)
	t := GetQuestionInput("time")
	fmt.Println("Displacement:", fn(t))
}