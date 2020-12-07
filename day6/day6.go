package day6

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
)

func Day6() {
	input := loadInput()
	count := 0
	for _, i := range input {
		count += len(i)
	}

	log.Println("day6 | Total answers anyone answered: ", count)

	count = 0
	for _, i := range input {
		num := i[255]
		for key, value := range i {
			if key != 255 && value == num {
				count++
			}
		}
	}

	log.Println("day6 | Total answers everyone answered: ", count)

}

func loadInput() []map[byte]int {
	file, err := ioutil.ReadFile("day6/input.txt")
	if err != nil {
		log.Fatalln("Cannot load file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(file))

	var allAnswers []map[byte]int
	var curAnswer map[byte]int

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			if curAnswer == nil {
				curAnswer = make(map[byte]int)
				curAnswer[255] = 0
			}

			for _, t := range text {
				i, OK := curAnswer[byte(t)]
				if OK {
					curAnswer[byte(t)] = i + 1
				} else {
					curAnswer[byte(t)] = 1
				}
			}
			curAnswer[255] = curAnswer[255] + 1 //saving the people count in the special slot.
		} else {
			if curAnswer != nil {
				allAnswers = append(allAnswers, curAnswer)
				curAnswer = nil
			}
		}
	}

	if curAnswer != nil {
		allAnswers = append(allAnswers, curAnswer)
	}

	return allAnswers

}
