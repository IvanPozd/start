package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Glossary struct {
	ID        string `json:"id"`
	Acronym   string `json:"acronym"`
	GlossTerm string `json:"glossTerm"`
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

func writeFile() {
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
	err = os.WriteFile("output.json", byteValue, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
