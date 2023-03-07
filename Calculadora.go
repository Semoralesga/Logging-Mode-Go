package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func division(a, b float32) float32 {
	if int(b) == 0 {
		fmt.Println("Error: El divisor no puede ser igual a 0")
	}

	return a / b
}

func stringToInt(s string) int {
	integer, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(" Input invalido: ", err.Error())
	}

	return integer
}

func checkInput(input string) []string {
	thereComma := strings.Contains(input, ",")
	if !thereComma {
	}
	numbers := strings.Split(input, ",")

	if len(numbers) > 2 {
		log.Fatal(" El input no es valido contiene mas de dos digitos, ej: 4,6")
	}

	return numbers
}

func doOperation(option int, numbers []string) {
	log.Print(" [Logging] La división es: ", division(float32(stringToInt(numbers[0])), float32(stringToInt(numbers[1]))))
	switch option {
	case 1:
		fmt.Printf(" La suma es: %d\n", sumar(stringToInt(numbers[0]), stringToInt(numbers[1])))
	case 2:
		fmt.Printf(" La resta es: %d\n", restar(stringToInt(numbers[0]), stringToInt(numbers[1])))
	default:
		fmt.Println("********** Ejecución Finalizada **********")
		os.Exit(0)
	}
}

func main() {

	fmt.Println("************** Calculadora de 2 Digitos **************")
	fmt.Println(" 1. Sumar \n 2. Restar \n * Cualquier otro valor finalizara la ejecución")
	fmt.Print(" Digite el numero de la operación que desea hacer: ")

	var option int
	_, err := fmt.Scan(&option)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(" Digite los numeros a operar separados por coma(1,2): ")

	var input string
	_, err = fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	numbers := checkInput(input)

	doOperation(option, numbers)

}
