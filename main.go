package main

import (
	"fmt"
	"log"
	"math"
)

var globalVar string = "This is global variable"

const pi = math.Pi

func main() {
	var delimiter string = "=========================================================================================="

	var phrase string = "Go rules"
	log.Print(phrase)
	fmt.Println(delimiter)

	var number int8 = 1
	var number1 uint8 = 2
	var isGood bool = true

	log.Print(number)
	log.Print(number1)
	log.Println("Число ПИ:", pi)
	log.Print(isGood)
	fmt.Println(delimiter)

	localVar1 := "Local variable"
	log.Println(localVar1)
	localVar1 = "some different"

	log.Println(globalVar)
	log.Println(localVar1)
	fmt.Println(delimiter)
}
