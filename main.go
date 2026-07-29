package main

import (
	"fmt"
	"log"
	"math"
)

var globalVar string = "This is global variable"

const pi = math.Pi

func getRadius(radius int8) float32 {
	return pi * float32(radius) * float32(radius)
}

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

	var circleRadius int8 = 6
	result := getRadius(circleRadius)
	log.Printf("Радиус круга: %d см.\n", circleRadius)
	log.Println("Формула расчёта площади круга: A=Pi*r2")
	log.Printf("Площадь круга: %f32 см. кв.\n", result)
	fmt.Println(delimiter)
}
