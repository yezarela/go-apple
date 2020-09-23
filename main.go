package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cakes := prompt("Please input the number of cakes")
	apples := prompt("Please input the number of apples")

	boxes := CountBoxes(cakes, apples)
	fmt.Println("Number of boxes: ", boxes)

	cakesQty, applesQty := CountItemAmount(cakes, apples, boxes)
	fmt.Printf("Each box contains: %d cakes and %d apples\n", cakesQty, applesQty)
}

// CountBoxes calculates how many boxes we can make
func CountBoxes(x, y int) int {

	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}

// CountItemAmount calculates how many apple and cakes in the box
func CountItemAmount(x, y, b int) (int, int) {
	return x / b, y / b
}

func prompt(s string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		v, err := strconv.Atoi(strings.TrimSuffix(response, "\n"))
		if err != nil {
			log.Fatal(err)
		}

		if v > 0 {
			return v
		}
	}
}
