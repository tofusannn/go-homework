package main

import (
	"fmt"
	"go-homework/intern"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("Hello, World")

	// Select prompt
	levels := []string{"Intern", "Junior", "Middle", "Senior"}

	// Create a Select prompt
	prompt := promptui.Select{
		Label: "Choose a level",
		Items: levels,
	}

	// Run the prompt
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You selected: %s\n", result)
	fmt.Printf("------------------------------------\n")
	switch result {
	case "Intern":
		intern.Main()
	case "Junior":
		fmt.Println("You are a junior")
	case "Middle":
		fmt.Println("You are a middle")
	case "Senior":
		fmt.Println("You are a senior")
	}
}
