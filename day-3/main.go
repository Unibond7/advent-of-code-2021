package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	fileData, _ := ioutil.ReadFile("input.txt")
	data := string(fileData)

	slice := strings.Split(data, "\n")

	ones := make([]int, len(slice[1]))
	zeros := make([]int, len(slice[1]))

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		bit := strings.Split(i, "")

		for k := range bit {
			if bit[k] == "1" {
				ones[k] += 1
			} else if bit[k] == "0" {
				zeros[k] += 1
			}
		}
	}

	var gamma_bit []string
	var epsilon_bit []string

	for j := range ones {
		if ones[j] > zeros[j] {
			gamma_bit = append(gamma_bit, "0")
			epsilon_bit = append(epsilon_bit, "1")
		} else if ones[j] < zeros[j] {
			gamma_bit = append(gamma_bit, "1")
			epsilon_bit = append(epsilon_bit, "0")
		}
	}

	gamma, _ := strconv.ParseInt(strings.Join(gamma_bit, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilon_bit, ""), 2, 64)

	fmt.Println(gamma * epsilon)
}
