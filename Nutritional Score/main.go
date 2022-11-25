package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(500),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumFromSalt(500),
		Fruits:              FruitsPercent(60),
		Fibre:               FibreGram(4),
		Protien:             ProtienGram(2),
	}, Food)

	fmt.Printf("Nutritional Score:%d\n", ns.Value)
}
