package day2

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// 7-14 m: mmcmqmmxmmmnmmrmcxc
type Policy struct {
	Low      int8
	High     int8
	Letter   byte
	Password string
}

func (policy *Policy) isValid() bool {
	var count int8 = 0

	for i := 0; i < len(policy.Password); i++ {
		if policy.Password[i] == policy.Letter {
			count++
			if count > policy.High {
				return false
			}
		}
	}

	return count >= policy.Low
}

func (policy *Policy) isValid2() bool {
	a := policy.Password[policy.Low-1] == policy.Letter
	b := policy.Password[policy.High-1] == policy.Letter
	return a != b
}

func Day2() {
	inputs := getInput()
	count := 0
	count2 := 0

	for _, policy := range *inputs {
		if policy.isValid() {
			count++
		}
		if policy.isValid2() {
			count2++
		}
	}

	log.Println("Total valid password for first policy: ", count)
	log.Println("Total valid password for second policy: ", count2)
}

func getInput() *[]Policy {
	content, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatalln("cannot read file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	var inputs []Policy

	for scanner.Scan() {
		row := scanner.Text()
		split := strings.Split(row, ":")
		part1 := strings.Split(split[0], " ")
		part0 := strings.Split(part1[0], "-")
		low, _ := strconv.Atoi(part0[0])
		high, _ := strconv.Atoi(part0[1])

		i := Policy{
			Password: strings.TrimSpace(split[1]),
			Letter:   part1[1][0],
			Low:      int8(low),
			High:     int8(high),
		}

		inputs = append(inputs, i)
	}

	return &inputs
}
