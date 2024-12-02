package main

import (
	"fmt"
	"github.com/qascade/AoC24/missionChief"
)

func RunDay1Part1() {
	input1, input2, err := missionChief.ReadIODay1("./resource/day1/input1.txt")

	distance := missionChief.Day1Part1(input1, input2)
	fmt.Println(distance)

	if err != nil {
		panic(err)
	}
}

func RunDay1Part2() {
	input1, input2, err := missionChief.ReadIODay1("./resource/day1/input2.txt")
	if err != nil {
		panic(err)
	}
	distance := missionChief.Day1Part2(input1, input2)
	fmt.Println(distance)

}

func RunDay2Part1() {
	input1, err := missionChief.ReadIODay2("./resource/day2/input1.txt")
	fmt.Println(input1)
	if err != nil {
		panic(err)
	}
	safeReports := missionChief.Day2Part1(input1)
	fmt.Println(safeReports)
}


func main() {
	//RunDay1Part1()
	// RunDay1Part2()
	RunDay2Part1()
}
