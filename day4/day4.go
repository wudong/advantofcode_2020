package day4

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *Passport) isValid() bool {
	return p.byr != "" && validNumber(p.byr, 1920, 2002) &&
		p.iyr != "" && validNumber(p.iyr, 2010, 2020) &&
		p.eyr != "" && validNumber(p.eyr, 2020, 2030) &&
		p.hgt != "" && (validHgt(p.hgt, "cm", 150, 193) || validHgt(p.hgt, "in", 59, 76)) &&
		p.hcl != "" && validHcl(p.hcl) &&
		p.ecl != "" && validEcl(p.ecl) &&
		p.pid != "" && validPid(p.pid)
}

func validHcl(str string) bool {
	b := len(str) == 7 && strings.HasPrefix(str, "#")
	if !b {
		return false
	}

	for _, s := range strings.TrimPrefix(str, "#") {
		v := (s >= '0' && s <= '9') || (s >= 'a' && s <= 'f')
		if !v {
			return false
		}
	}

	return true
}

func validEcl(str string) bool {
	switch str {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}
	return true
}

func validPid(str string) bool {
	b := len(str) == 9
	if !b {
		return false
	}
	for _, s := range str {
		if s < '0' || s > '9' {
			return false
		}
	}
	return true
}

func validNumber(str string, low int, high int) bool {
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return atoi >= low && atoi <= high
}

func validHgt(str string, sufix string, low int, high int) bool {
	if strings.HasSuffix(str, sufix) {
		num := strings.TrimSuffix(str, sufix)
		return validNumber(num, low, high)
	} else {
		return false
	}
}

func Day4() {
	passports := loadInput()
	counter := 0
	for _, v := range *passports {
		if v.isValid() {
			counter++
		}
	}
	log.Println("day4 | Total valid passport: ", counter)
}

func loadInput() *[]Passport {
	file, err := ioutil.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatalln("cannt load file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(file))
	var passports []Passport
	var current *Passport
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			if current == nil {
				current = &Passport{}
			}
			loadLine(text, current)
		} else { //emtpy line
			if current != nil {
				passports = append(passports, *current)
				current = nil
			}
		}
	}

	if current != nil {
		passports = append(passports, *current)
	}

	return &passports
}

func loadLine(text string, current *Passport) {
	splits := strings.Split(text, " ")
	for _, s := range splits {
		parts := strings.Split(s, ":")
		switch parts[0] {
		case "byr":
			current.byr = parts[1]
		case "iyr":
			current.iyr = parts[1]
		case "eyr":
			current.eyr = parts[1]
		case "hgt":
			current.hgt = parts[1]
		case "hcl":
			current.hcl = parts[1]
		case "ecl":
			current.ecl = parts[1]
		case "pid":
			current.pid = parts[1]
		case "cid":
			current.cid = parts[1]
		}
	}
}
