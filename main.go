package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
)

var globalVar string = "This is global variable"

const pi = math.Pi

func calculateCirlceAre(radius int8) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус на может быть равен нулю или иметь отрицательное значение")
	}
	return pi * float32(radius) * float32(radius), nil
}

func printCircleRadius(radius int8, result float32) {
	log.Printf("Радиус круга: %d см.\n", radius)
	log.Println("Формула расчёта площади круга: A=Pi*r2")
	log.Printf("Площадь круга: %f32 см. кв.\n", result)
}

func practice() {
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

	var circleRadius int8 = 9
	result, err := calculateCirlceAre(circleRadius)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		printCircleRadius(circleRadius, result)
	}
	fmt.Println(delimiter)

	x := 10
	log.Printf("x = %d", x)
	log.Println("&x =", &x)

	p := &x
	log.Println("p =", p)
	*p = 15
	log.Println("*p =", *p)
	log.Println("x =", x)
	fmt.Println(delimiter)
}

type Glossary struct {
	Id        string `json: "ID"`
	Acronym   string `json: "Acronym"`
	GlossTerm string `json: "GlossTerm"`
}

func read_file() {
	jsonFile, err := os.Open("example.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var glos Glossary
	err = json.Unmarshal(byteValue, &glos)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Gloss: %+v\n", glos)
	fmt.Println(glos)
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}

	// Marshal to JSON with indentation
	byteValue, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write to file
	err = ioutil.WriteFile("output.json", byteValue, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
