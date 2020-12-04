package day3

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"math"
)

func Day3() {
	mm, width := loadMap()

	var slopes = [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	counter := 1
	for _, slope := range slopes {
		c := getNumberOfTree(width, mm, slope[0], slope[1])
		counter = c * counter
	}

	log.Println("the multiple of all number of trees encountered: ", counter)
}

func getNumberOfTree(width int, mm [][]bool, right int, down int) int {
	length := len(mm)
	counter := 0
	x := 0
	y := 0
	for y < length {
		//take a step
		x = int(math.Mod(float64(x+right), float64(width)))
		y = y + down
		if y < length && mm[y][x] {
			counter++
		}
	}
	return counter
}

func loadMap() ([][]bool, int) {
	file, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatalln("cannot load file")
	}

	reader := bytes.NewReader(file)
	scanner := bufio.NewScanner(reader)

	var treeMap [][]bool
	var width int

	for scanner.Scan() {
		text := scanner.Text()
		width = len(text)
		bools := make([]bool, width)

		for i, v := range text {
			bools[i] = v == '#'
		}
		treeMap = append(treeMap, bools)
	}

	return treeMap, width

}
