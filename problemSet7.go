package main

import (
	"fmt"
	"math"
	"math/rand"
)

func calcCircle(radius float64) (float64, float64) {

	var circunference = 2 * math.Pi * radius
	var area = math.Pi * radius * radius

	return circunference, area

}

func main1() {

	var rebanada = make([]int, 3)
	var suma = 0

	for i := 0; i < 3; i++ {
		rebanada[i] = rand.Intn(100)
		suma = suma + rebanada[i]
	}

	fmt.Println("Hola mundo")
	fmt.Print(rebanada[0], " + ", rebanada[1], " + ", rebanada[2], " = ", suma, "\n")

	radius := 3.5
	circunference, area := calcCircle(radius)
	fmt.Print("Circle of radius: ", radius, " , circle's circunferce: ", circunference, ", and circle's area: ", area)

}
