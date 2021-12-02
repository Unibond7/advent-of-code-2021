package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	forward := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		v := strings.Split(i, " ")

		n, err := strconv.Atoi(v[1])
		if err != nil {
			fmt.Println(err)
		}

		if v[0] == "forward" {
			forward += n
			depth += aim * n
		} else if v[0] == "down" {
			aim += n
		} else if v[0] == "up" {
			aim -= n
		}
	}

	ans := forward * depth

	fmt.Println(ans)

}
