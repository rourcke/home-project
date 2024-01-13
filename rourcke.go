package main

import "fmt"

// Define a struct for Furniture
type Furniture struct {
	Name     string
	Material string
}

// Define a struct for Room
type Room struct {
	Name      string
	Furniture []Furniture
}

// Define a struct for Home
type Home struct {
	Name  string
	Rooms []Room
}

func main() {
	// Create some furniture items
	sofa := Furniture{Name: "Sofa", Material: "Leather"}
	table := Furniture{Name: "Coffee Table", Material: "Wood"}

	// Create a living room and add furniture
	livingRoom := Room{Name: "Living Room", Furniture: []Furniture{sofa, table}}

	// Create a home and add rooms
	myHome := Home{Name: "My Home", Rooms: []Room{livingRoom}}

	// Display information about the home and its contents
	fmt.Printf("Welcome to %s\n", myHome.Name)
	for _, room := range myHome.Rooms {
		fmt.Printf("Room: %s\n", room.Name)
		fmt.Println("Furniture:")
		for _, furniture := range room.Furniture {
			fmt.Printf("- %s (Material: %s)\n", furniture.Name, furniture.Material)
		}
		fmt.Println()
	}
}
