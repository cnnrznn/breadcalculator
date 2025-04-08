# breadcalculator

Welcome! Bread calculator is a tool for calculating a bread recipe based on the final dough weight and baker's percentages.

## Why?

I did this because existing tools [1, 2, 3, 4, 5] do not calculate the recipe based on the final dough weight.
This is annoying because I want to keep my final loaf weight the same while changing the percentages (salt, hydration, inoculation).

## Installation

1. Install [Go](https://go.dev/doc/install)
1. ```go install github.com/cnnrznn/breadcalculator```

## Usage

```sh
breadcalculator -h
```

### Default recipe

The default recipe by `breadcalculator` is 75% hydration, 20% inoculation, and 2% salt.

## References
1. http://brdclc.com/
2. https://foodgeek.io/en/bread-calculator/
3. https://asuratoom.com/bread-calculator
4. https://breadcalc.com/
5. https://www.breadratiocalculator.com/
