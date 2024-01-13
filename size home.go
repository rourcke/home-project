package main

import "fmt"

// Define a struct for Furniture
type Furniture struct {
	Name     string
	Material string
}

// Define a struct for Room with size information
type Room struct {
	Name      string
	Size      string // Add size information, e.g., "20x15"
	Furniture []Furniture
}

// Define a struct for Home with size information
type Home struct {
	Name  string
	Size  string // Add size information, e.g., "2000 sq ft"
	Rooms []Room
}

func main() {
	// Create some furniture items
	sofa := Furniture{Name: "Sofa", Material: "Leather"}
	table := Furniture{Name: "Coffee Table", Material: "Wood"}

	// Create a living room and add furniture
	livingRoom := Room{Name: "Living Room", Size: "25x20", Furniture: []Furniture{sofa, table}}

	// Create a home and add rooms
	myHome := Home{Name: "My Home", Size: "2000 sq ft", Rooms: []Room{livingRoom}}

	// Display information about the home and its contents
	fmt.Printf("Welcome to %s\n", myHome.Name)
	fmt.Printf("Size: %s\n\n", myHome.Size)
	for _, room := range myHome.Rooms {
		fmt.Printf("Room: %s\n", room.Name)
		fmt.Printf("Size: %s\n", room.Size)
		fmt.Println("Furniture:")
		for _, furniture := range room.Furniture {
			fmt.Printf("- %s (Material: %s)\n", furniture.Name, furniture.Material)
		}
		fmt.Println()
	}
}
