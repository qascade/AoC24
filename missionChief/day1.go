package missionChief

import (
	"math"
	"sort"
)

func Day1Part1(input1, input2 []int) int {
	// As we want the smallest in front of the smallest lets just sort them
	sort.Ints(input1)
	sort.Ints(input2)

	distance := 0

	for i, _ := range input1 {
		distance = distance + int(math.Abs(float64(input1[i])-float64(input2[i])))
	}
	return distance
}

func Day1Part2(input1, input2 []int) int {
	// Calculate frequency of numbers in input2
	counter := make(map[int]int)

	for _, value := range input2 {
		if _, ok := counter[value]; !ok {
			counter[value] = 0
		}
		counter[value] += 1
	}

	similarityScore := 0
	for _, input := range input1 {
		if count, ok := counter[input]; ok {
			similarityScore += input * count
		}
	}
	return similarityScore
}
