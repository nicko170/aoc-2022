package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/* The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

The input is stored in a file called day1.txt
*/

func main() {
	// Read the input file
	input, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}

	// Split the input into lines
	lines := strings.Split(string(input), "\n")

	// The number of empty lines in the number of elves
	elves := 1
	for _, line := range lines {
		if line == "" {
			elves++
		}
	}

	// create a slice of slices to store the calories for each elf
	calories := make([][]int, elves)

	// loop through the lines and add the calories to the correct slice for each elf
	elf := 0
	for _, line := range lines {
		if line == "" {
			elf++
		} else {
			res, _ := strconv.Atoi(line)
			calories[elf] = append(calories[elf], res)
		}
	}

	// Find the elf with the most calories, and print the total
	max := 0
	for _, elf := range calories {
		total := 0
		for _, calorie := range elf {
			total += calorie
		}
		if total > max {
			max = total
		}
	}

	// Find the top 3 elves with the most calories, and print the total
	max1 := 0
	max2 := 0
	max3 := 0
	for _, elf := range calories {
		total := 0
		for _, calorie := range elf {
			total += calorie
		}
		if total > max1 {
			max3 = max2
			max2 = max1
			max1 = total
		} else if total > max2 {
			max3 = max2
			max2 = total
		} else if total > max3 {
			max3 = total
		}
	}

	fmt.Println(max)
	fmt.Println(max1 + max2 + max3)
}
