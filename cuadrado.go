package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func areaCuadrado(number int) {
	fmt.Printf("Cuadrado es: %d ", number*number)
}

func areaCubo(number int) int {
	return number * number * number
}

func returnDosValores(number int) (int, int) {
	return number * 2, number * 3
}

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	//const baseCuadrado = 10
	//areaCuadrado := baseCuadrado * baseCuadrado

	//fmt.Println("El área del cuadrado es:", areaCuadrado)

	areaCuadrado(10)

	cubo := areaCubo(3)
	fmt.Printf("Cubo es: %d ", cubo)
	fmt.Println("")

	var1, var2 := returnDosValores(5)
	fmt.Printf("Valor 1: %d, Valor 2: %d ", var1, var2)
	fmt.Println("")

	var3, _ := returnDosValores(3)
	fmt.Printf("Valor 3: %d ", var3)
	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i
	fmt.Println("")
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		time.Sleep(time.Second / 3)
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
	fmt.Println("")
	for i := 10; i > 0; i-- {
		time.Sleep(time.Second / 2)
		fmt.Println(i)
	}

	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}

	// Convertir testo a numero
	value, err := strconv.Atoi("34")
	if err != nil {
		fmt.Println("Xopaaa eso no es un número")
		log.Fatal(err)
	}
	fmt.Printf("Tu numero es: %d", value)

	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("Switch - Múltiples condiciones ¿anidadas?")
	fmt.Println("=========================================")
	fmt.Println()

	fmt.Println("Sirve para simplificar múltiples condiciones en secuencia que implementadas con IF hace un código extenso.")
	fmt.Println()

	fmt.Println("A continuación en el código se declara de la forma mas parecida a C/Java/Kotlin donde de nuevo la condición" +
		" se colocas sin paréntesis conteniéndola.")
	fmt.Println("Interesante también ver que cada opción CASE no termina en un 'break' como en los otros lenguajes.")
	fmt.Println()

	fmt.Println("el resultado de la evaluación de si es par o no < modulo := 5 % 2 > es:")
	fmt.Println()

	const moduloInicial int8 = 5 % 2

	switch moduloInicial {
	case 0:
		fmt.Println("> Es par")
	default:
		fmt.Println("> Es impar")
	}

	fmt.Println()
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println("Se intenta una variación, donde se pone diréctamente la condición en el switch => switch 7 & 2 {....")
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println()

	switch 7 & 2 {
	case 0:
		fmt.Println("> Es par")
	default:
		fmt.Println("> Es impar")
	}

	fmt.Println()
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println("Forma directa de la instrucción donde se asigna el resultado en la variable.\nPor ahora le veo poca utilidad" +
		" a esta forma.")
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println()

	switch moduloVariableContextoInterno := 4 % 2; moduloVariableContextoInterno {
	case 0:
		fmt.Println("Aquí entró la evaluación de la condición, que dió", moduloVariableContextoInterno)
	default:
		fmt.Println("2> Es impar")
	}

	fmt.Println()
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println("SWITCH sin condición, abierto - Anidas múltiples condiciones")
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Println()

	moduloVariableExterno := 200

	switch {
	case moduloVariableExterno < 100 && moduloInicial == 1:
		fmt.Println("Es mayor a 100")
	case moduloVariableExterno > 150 || moduloInicial == 1:
		fmt.Println("Es mayor a 150")
	default:
		fmt.Println("Este es si no se cumple alguno otro")
	}

	fmt.Println()

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 4 {
			fmt.Println("Break")
			break
		}
	}

}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}
