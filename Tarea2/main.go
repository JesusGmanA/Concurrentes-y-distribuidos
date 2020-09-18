package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

const OPC_AREA_CUADRADO = 1
const OPC_AREA_CIRCULO = 2
const OPC_AREA_TRIANGULO = 3
const OPC_CONV_TEMPS = 4
const EXIT = 5

func main() {
	var opc = 0

	for opc != EXIT {
		limpiarPantalla()
		opc = obtenerOpcMenu()
	}

	fmt.Print("Hasta luego!")
}

func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func convertirTempsACelsius(temps float64) float64 {
	return ((temps - 32) * 5 / 9)
}

func calcularAreaCuadro(lado float64) float64 {
	return lado * lado
}

func calcularAreaTriangulo(base, altura float64) float64 {
	return (base * altura / 2)
}

func calcAreaCirculo(radio float64) float64 {
	return math.Pi * math.Pow(radio, 2)
}

func obtenerOpcMenu() int {
	var option int
	fmt.Println("Escribe la accion a realizar:")
	fmt.Println("1.Calcular area del cuadrado")
	fmt.Println("2.Calcular area del circulo")
	fmt.Println("3.Calcular area del triangulo")
	fmt.Println("4.Calcular temperatura en grados Celsius")
	fmt.Println("5. Salir")
	fmt.Print("Opcion: ")
	fmt.Scanf("%d", &option)

	switch option {
	case OPC_AREA_CUADRADO:
		var lado float64
		fmt.Print("Dame el lado del cuadrado: ")
		fmt.Scanf("%f", &lado)
		fmt.Printf("El area del cuadrado es: %.2f\n", calcularAreaCuadro(lado))

	case OPC_AREA_CIRCULO:
		var radio float64
		fmt.Print("Dame el radio del circulo: ")
		fmt.Scanf("%f", &radio)
		fmt.Printf("El araea del circulo es: %.2f\n", calcAreaCirculo(radio))

	case OPC_AREA_TRIANGULO:
		var base, altura float64
		fmt.Print("Ingresa la base del triangulo: ")
		fmt.Scanf("%f", &base)
		fmt.Print("Ingresa la altura del triangulo: ")
		fmt.Scanf("%f", &altura)
		fmt.Printf("El area del cuadrado es: %.2f\n", calcularAreaTriangulo(base, altura))

	case OPC_CONV_TEMPS:
		var temps float64
		fmt.Print("Dame la temperatura en Fahrenheit : ")
		fmt.Scanf("%f", &temps)
		fmt.Printf("La temperatura en Celsius es: %.2f°C\n", convertirTempsACelsius(temps))
	case EXIT:
	default:
		fmt.Println("Opcion inexistente intente de nuevo!")
	}

	if option != EXIT {
		fmt.Print("Presiona 'Enter' para continuar...")
		fmt.Scanln() //El primero se come el "Enter" atorado en el buffer cuando se lee algo desde consola
		fmt.Scanln()
	}
	return option
}
