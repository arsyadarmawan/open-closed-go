package main

import "fmt"

func main() {
	buildingType := "house"
	switch buildingType {
	case "house":
		house := House{}
		fmt.Println(house.GetComposition())
		fmt.Println(house.MaxHeight())
	case "skyscrapper":
		building := SkyScrapper{}
		fmt.Println(building.MaxHeight())
		fmt.Println(building.GetComposition())
	}
}
