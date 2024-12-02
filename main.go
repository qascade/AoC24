package main

import (
	"fmt"
	"github.com/qascade/AoC24/scripts"
)

func RunDay1Part1() {
	input1, input2, err := scripts.ReadIODay1("./resource/day1/input1.txt")

	distance := scripts.Day1Part1(input1, input2)
	fmt.Println(distance)

	if err != nil {
		panic(err)
	}
}

func RunDay1Part2() {
	input1, input2, err := scripts.ReadIODay1("./resource/day1/input2.txt")

	distance := scripts.Day1Part2(input1, input2)
	fmt.Println(distance)

	if err != nil {
		panic(err)
	}
}

func main() {
	//RunDay1Part1()
	RunDay1Part2()
}
