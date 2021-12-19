package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Printf("You have %d gallons of fuel left!", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}
func greetPlanet(planet string) {
	fmt.Printf("Welcome to the planet: %v", planet)
}
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaing, fuelCost int
	fuelRemaing = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaing >= fuelCost {
		greetPlanet(planet)
		fuelRemaing -= fuelCost
	} else if fuelCost >= fuelRemaing {
		cantFly()
	}
	return fuelRemaing
}
func main() {
	fuel := 1000000
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
