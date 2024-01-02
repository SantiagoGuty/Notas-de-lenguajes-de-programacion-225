package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ISBN() {
	var resultado string
	var x int = 10
	var suma int = 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < x-1; i++ {

		if (i == 1) || (i == 3) || (i == 9) {
			num := r.Intn(10) - 1
			resultado += fmt.Sprint("-", num)
			suma += (i + 1) * num
		} else {
			num := r.Intn(10) - 1
			resultado += fmt.Sprint(num)
			suma += (i + 1) * num
		}
	}

	divideSum := suma % 11
	if divideSum < 10 {
		resultado += fmt.Sprint("-", divideSum)
	} else {
		resultado += fmt.Sprint("-X")
	}

	fmt.Print("\nISBN: ", resultado, "\n")
	fmt.Println("\ndivido: ", divideSum)
	fmt.Println("\nsuma: ", suma)

}

func main() {
	ISBN()
}
