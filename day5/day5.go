package day5

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"sort"
)

func Day5() {
	input := *loadInput()
	highestId := 0
	for _, v := range input {
		id := v.Id()
		if id > highestId {
			highestId = id
		}
	}
	log.Println("day5 | Highest id: ", highestId)

	//sorting.
	sort.Slice(input, func(i, j int) bool {
		return input[i].Id() < input[j].Id()
	})

	gap := 0
	for idx, v := range input {
		if idx == 0 || idx == len(input)-1 {
			continue
		} else {
			//looking for the gap.
			if v.Id()-input[idx-1].Id() == 2 {
				gap = v.Id() - 1
				log.Println("day5 | Missing id found:", gap)
			}
		}

	}
}

type BoardPass struct {
	row    int
	column int
}

func (b *BoardPass) Id() int {
	return b.row*8 + b.column
}

func loadInput() *[]BoardPass {
	file, err := ioutil.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatalln("cannot load file")
	}

	var passes []BoardPass

	scanner := bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		text := scanner.Text()
		pass := parse(text)
		passes = append(passes, pass)
	}

	return &passes
}

func parse(text string) BoardPass {
	pass := BoardPass{}
	low := 0
	high := 127
	for i := 0; i < 7; i++ { //parsing row.
		low, high = div(low, high, text[i] == 'F')
	}
	pass.row = low
	low = 0
	high = 7
	for i := 7; i < 10; i++ { //parsing row.
		low, high = div(low, high, text[i] == 'L')
	}
	pass.column = low
	return pass
}

func div(low int, high int, firstHalf bool) (newLow int, newHigh int) {
	if firstHalf {
		return low, low + (high-low)/2
	} else {
		return low + (high-low)/2 + 1, high
	}
}
