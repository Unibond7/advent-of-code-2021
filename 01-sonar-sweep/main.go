package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var list []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, i)
	}

	top_list := list[0 : len(list)-2]
	mid_list := list[1 : len(list)-1]
	bot_list := list[2:]

	var sweep []int
	for k := range top_list {
		sum_1 := top_list[k] + mid_list[k]
		sum_2 := sum_1 + bot_list[k]
		sweep = append(sweep, sum_2)
	}

	top_sweep := sweep[0 : len(sweep)-1]
	bot_sweep := sweep[1:]

	drop := 0

	for k := range top_sweep {
		ans := bot_sweep[k] - top_sweep[k]
		if ans > 0 {
			drop += 1
		}
	}
	fmt.Println(drop)
}
