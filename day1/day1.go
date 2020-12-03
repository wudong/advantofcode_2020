package day1

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

func Day1()  {
	content, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatalln("Cannot read input file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	var result []int

	for scanner.Scan() {
		intV, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln("cannot convert to integer", intV)
		}
		result = append(result, intV)
	}

	sort.Ints(result)
	i:=0
	j:=len(result) -1

	for i < j {
		sum :=result[i] + result[j]
		if sum == 2020 {
			log.Println("two result found:", result[i] , result[j], result[i] * result[j])
			break
		}else if sum >2020 {
			j= j - 1
		} else {
			i = i+1
		}
	}

	i=0
	j=len(result) -1
	out: for i < j {
		k:=i+1
		l:=j
		newSum := 2020 - result[i]

		for  k<l {
			sum := result[l] + result[k]
			if sum == newSum {
				log.Println("tree result found:", result[i] , result[k], result[l], result[i] * result[k] * result[l])
				break out
			}else if sum > newSum {
				l=l-1
			}else {
				k=k+1
			}
		}
		i=i+1
	}


}
