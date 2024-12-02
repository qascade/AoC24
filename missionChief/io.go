package missionChief

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This takes input 2 columns.
func ReadIODay1(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input1, input2 := make([]int, 0), make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		a1, err := strconv.Atoi(fields[0])
		if err != nil {
			err = fmt.Errorf("not able to convert string to int: %v: %v", a1, err)
			return nil, nil, err
		}
		a2, err := strconv.Atoi(fields[1])
		if err != nil {
			err = fmt.Errorf("not able to convert string to int: %v: %v", a2, err)
			return nil, nil, err
		}
		input1 = append(input1, a1)
		input2 = append(input2, a2)

	}
	return input1, input2, nil
}

// This reads 2D array with no. of columns not fixed. 
func ReadIODay2(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	reports := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		levels := make([]int, len(fields))
		for i, field := range fields {
			f, err := strconv.Atoi(field)
			if err != nil {
				err = fmt.Errorf("not able to convert string to int: %v: %v", f, err)
				return nil, err
			}
			levels[i] = f
		}
		reports = append(reports, levels)
	}

	return reports, nil
}