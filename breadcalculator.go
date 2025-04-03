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
// salt is 2% of flour
// total_weight is (water + flour + inoculation) * saltFactor + salt
// total_weight = (water + 1.02*flour + inoculation) * saltFactor
// saltFactor = total_weight / (water + 1.02*flour + inoculation)
//
// multiply flour, fluid and inoculation by saltFactor
// saltFactor is total weight minus new flour, fluid, inoculation
//
// The steps to solving this system are as follows.
// 1. Calculate total_flour and total_water based on the weight and the hydration percentage
// 2.
package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

func main() {
	var (
		weight      int
		hydration   int
		inoculation int
		salinity    int
	)
	flag.IntVar(
		&weight,
		"weight",
		700,
		"final dough weight",
	)
	flag.IntVar(
		&hydration,
		"hydration",
		75,
		"hydration percentage (0-100)",
	)
	flag.IntVar(
		&inoculation,
		"inoculation",
		20,
		"inoculation percentage (0-100)",
	)
	flag.IntVar(
		&salinity,
		"salinity",
		3,
		"salt percentage (0-100)",
	)
	flag.Parse()

	// Create recipe
	r, err := createRecipe(
		float64(weight),
		float64(inoculation),
		float64(hydration),
		float64(salinity),
	)
	if err != nil {
		panic(err)
	}

	// Display recipe
	fmt.Println(r)
}

func createRecipe(weight, inoculation, hydration, salinity float64) (*Recipe, error) {
	if weight <= 0 || inoculation <= 0 || hydration <= 0 || salinity <= 0 {
		return nil, errors.New("non-negative weight, inoculation, hydration values in recipe only")
	}

	saltPercentage := salinity / 100

	hydration_percentage := hydration / 100
	total_flour := weight / (1 + hydration_percentage)
	total_fluid := weight - total_flour

	inoculation_percentage := inoculation / 100
	flour := total_flour / (1 + (0.5 * inoculation_percentage))
	fluid := total_fluid - (0.5 * inoculation_percentage * flour)
	starter := weight - flour - fluid

	saltFactor := weight / (((1 + saltPercentage) * flour) + fluid + starter)
	flour = saltFactor * flour
	fluid = saltFactor * fluid
	starter = saltFactor * starter
	salt := weight - flour - fluid - starter

	return &Recipe{
		Flour:       flour,
		Water:       fluid,
		Starter:     starter,
		Salt:        salt,
		TotalWeight: weight,
	}, nil
}

// Recipe is a dough recipe that assumes a starter hydration of 100%.
// Flour is the mass of non-starter flour.
// Water is the mass of non-starter fluid.
// Starter is the mass of starter.
// Salt is the mass in salt.
type Recipe struct {
	Flour       float64
	Water       float64
	Starter     float64
	Salt        float64
	TotalWeight float64
}

func (r Recipe) String() string {
	return fmt.Sprintf(
		"Weight: %v\nFlour: %v\nWater: %v\nStarter: %v\nSalt: %v\n",
		math.Round(r.TotalWeight),
		math.Round(r.Flour),
		math.Round(r.Water),
		math.Round(r.Starter),
		math.Round(r.Salt),
	)
}

func (r Recipe) Valid() bool {
	if int(r.TotalWeight) != int(r.Flour+r.Water+r.Water+r.Salt) {
		return false
	}
	return true
}
