// breadcalculator is a tool for creating dough recipes based on a recipe's
// total weight, inoculation percentage, and hydration percentage.
//
// The resulting recipe is the mass for each component: starter, flour, water
//
// Below I've outlined my system of equations.
//
// parameters

// total dough weight
// percent hydration
// percent inoculation

// total dough weight = flour + inoculation_flour + inoculation_hydration + hydration
// total dough weight = flour + water

// 75% hydration = 100g flour, 75g water
// 50% inoculation = x + .25x = 100g
//                 = y + .25x = 75g
// flour:
// 100 = 1.25x
// x = 80
//
// fluid:
// .25(80) + y = 75
// 20 + y = 75
// y = 55

// x is flour (non-inoculation)
// y is fluid (non-inoculation)
// at 50% inoculation, the total flour is 1/2 of 0.5 * flour

// weight = total_flour + total_hydration
// weight = total_flour + hydration(total_flour)
// weight = total_flour * (1 + hydration)
// total_flour = weight / (1 + hydration)
// total_fluid = weight - total_flour
//
// The steps to solving this system are as follows.
// 1. Calculate total_flour and total_water based on the weight and the hydration percentage
// 2.
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Read params from terminal

	// Create recipe
	r, err := createRecipe(500, 40, 70)
	if err != nil {
		panic(err)
	}

	// Display recipe
	fmt.Println(r)
}

func createRecipe(weight, inoculation, hydration int) (*Recipe, error) {
	if weight <= 0 || inoculation <= 0 || hydration <= 0 {
		return nil, errors.New("non-negative weight, inoculation, hydration values in recipe only")
	}

	hydration_percentage := float32(hydration) / 100
	total_flour := float32(weight) / (1 + hydration_percentage)
	total_fluid := float32(weight) - total_flour

	// total_flour = (1 + .5(inoculation_percentage)) * flour
	// flour = total_flour / (1 + .5(inoculation_percentage)
	inoculation_percentage := float32(inoculation) / 100
	flour := total_flour / (1 + (0.5 * inoculation_percentage))
	fluid := total_fluid - (0.5 * inoculation_percentage * flour)

	return &Recipe{
		Flour:       flour,
		Fluid:       fluid,
		Inoculation: float32(weight) - flour - fluid,
		TotalWeight: float32(weight),
	}, nil
}

// Recipe is a dough recipe that assumes a starter hydration of 100%.
// Flour is the mass of non-inoculation flour.
// Fluid is the mass of non-inoculation fluid.
// Inoculation is the mass of starter.
type Recipe struct {
	Flour       float32
	Fluid       float32
	Inoculation float32
	TotalWeight float32
}

func (r Recipe) String() string {
	return fmt.Sprintf("Weight: %v\nFlour: %v\nFluid: %v\nInoculation: %v\n", r.TotalWeight, r.Flour, r.Fluid, r.Inoculation)
}

func (r Recipe) Valid() bool {
	if int(r.TotalWeight) != int(r.Flour+r.Fluid+r.Fluid) {
		return false
	}
	return true
}
