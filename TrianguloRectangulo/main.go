package main

import (
	"fmt"
	"math"
)

func main() {
	//Declaración de variables
	var lado1, lado2, hipotenusa, area, perimetro float64
	const precision = 2

	//Inicio del Programa
	fmt.Println("****************PROGRAMA PARA CALCULAR EL ÁREA Y PERÍMETRO DE UN TRÍANGULO******************************")

	//Ingreso de información
	fmt.Println("Ingrese el lado 1:")
	fmt.Scanln(&lado1)

	fmt.Println("Ingrese el lado 2:")
	fmt.Scanln(&lado2)

	//Calculamos la hipotenusa
	hipotenusa = math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))

	//Calcular el área
	area = (lado1 * lado2) / 2

	//Calcular el perimetro
	perimetro = lado1 + lado2 + hipotenusa

	//Resultados
	fmt.Printf("\nÁrea: %.*f \n", precision, area)
	fmt.Printf("Perímetro: %.*f \n", precision, perimetro)

}
