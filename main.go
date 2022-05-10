package main

import (
	"fmt"
)



func main() {
	ns := GetNutritionalScore(
		NutritionalData {
			Energy: EnergyFromKcal(100),
			Sugars: SugarGram(10),
			SaturatedFattyAcids: SaturatedFattyAcids(2),
			Sodium: SodiumMilliGram(500),
			Fruits: FruitsPercent(60),
			Fiber: FiberGram(4),
			Protein: ProteinGram(2),
		}, Food)

	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}

