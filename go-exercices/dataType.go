package main

import (
	"fmt"
)

func main() {
	// int - números inteiros com sinal (8, 16, 32, 64 bits)
	var num1 int8 = 127
	var num2 int16 = 32767
	var num3 int32 = 2147483647
	var num4 int64 = 9223372036854775807

	// uint - números inteiros sem sinal (8, 16, 32, 64 bits)
	//"uint" é uma abreviação de "unsigned integer" em Go, que se refere a um tipo de dado numérico inteiro sem sinal (ou seja, apenas números positivos ou zero). Em Go, uint pode ser usado com diferentes bits de tamanho, como uint8, uint16, uint32 e uint64, dependendo da quantidade de bits que se deseja usar.
	var num5 uint8 = 255
	var num6 uint16 = 65535
	var num7 uint32 = 4294967295
	var num8 uint64 = 18446744073709551615

	// float - números de ponto flutuante (32, 64 bits)
	var num9 float32 = 3.14159265358979323846
	var num10 float64 = 3.14159265358979323846264338327950288419716939937510582097494459

	// boolean - valores verdadeiros ou falsos
	var bool1 bool = true
	var bool2 bool = false

	// string - sequências de caracteres
	var str1 string = "Hello, world!"

	// rune - representação Unicode (32 bits)
	var char rune = '♥'

	fmt.Printf("num1: %d\n", num1)
	fmt.Printf("num2: %d\n", num2)
	fmt.Printf("num3: %d\n", num3)
	fmt.Printf("num4: %d\n", num4)
	fmt.Printf("num5: %d\n", num5)
	fmt.Printf("num6: %d\n", num6)
	fmt.Printf("num7: %d\n", num7)
	fmt.Printf("num8: %d\n", num8)
	fmt.Printf("num9: %f\n", num9)
	fmt.Printf("num10: %f\n", num10)
	fmt.Printf("bool1: %t\n", bool1)
	fmt.Printf("bool2: %t\n", bool2)
	fmt.Printf("str1: %s\n", str1)
	fmt.Printf("char: %c\n", char)
}
