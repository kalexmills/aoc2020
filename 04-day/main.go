package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(4)
	defer input.Close()
	fmt.Println(solve1(parseInput(input)))
}

func solve1(passports []Passport) int {
	count := 0
	for _, p := range passports {
		if p.IsValid() {
			count++
		}
	}
	return count
}

type Passport struct { // stringly-typed! ^_^
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p Passport) IsValid() bool {
	return validBirthYear(p.BirthYear) &&
		validIssueYear(p.IssueYear) &&
		validExpirationYear(p.ExpirationYear) &&
		validHeight(p.Height) &&
		validHairColor(p.HairColor) &&
		validEyeColor(p.EyeColor) &&
		validPassportID(p.PassportID)
}

func validBirthYear(byr string) bool {
	return byr != "" && validYear(byr, 1920, 2002)
}

func validIssueYear(iyr string) bool {
	return iyr != "" && validYear(iyr, 2010, 2020)
}

func validExpirationYear(eyr string) bool {
	return eyr != "" && validYear(eyr, 2020, 2030)
}

func validYear(str string, min, max int) bool {
	year, err := strconv.ParseInt(str, 10, 64)
	return err == nil && int64(min) <= year && year <= int64(max)
}

func validHeight(str string) bool {
	if str == "" {
		return false
	}
	if strings.HasSuffix(str, "cm") {
		cm, err := strconv.ParseInt(strings.ReplaceAll(str, "cm", ""), 10, 64)
		return err == nil && 150 <= cm && cm <= 193
	} else if strings.HasSuffix(str, "in") {
		in, err := strconv.ParseInt(strings.ReplaceAll(str, "in", ""), 10, 64)
		return err == nil && 59 <= in && in <= 76
	}
	return false
}

func validHairColor(str string) bool {
	if str == "" {
		return false
	}
	if !strings.HasPrefix(str, "#") {
		return false
	}
	_, err := strconv.ParseInt(strings.ReplaceAll(str, "#", ""), 16, 64)
	return err == nil && len(str) == 7
}

func validEyeColor(str string) bool {
	if str == "" {
		return false
	}
	for _, test := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if str == test {
			return true
		}
	}
	return false
}

func validPassportID(str string) bool {
	if str == "" {
		return false
	}
	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil && len(str) == 9
}

func parseInput(reader io.Reader) []Passport {
	var result []Passport
	var currTokens []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, parsePassport(currTokens))
			currTokens = nil
		} else {
			currTokens = append(currTokens, strings.Split(scanner.Text(), " ")...)
		}
	}
	return result
}

func parsePassport(tokens []string) Passport {
	result := Passport{}
	for _, token := range tokens {
		fields := strings.Split(token, ":")
		if len(fields) != 2 {
			log.Printf("ignored malformed token: %s", token)
			continue
		}
		switch fields[0] {
		case "byr":
			result.BirthYear = fields[1]
		case "iyr":
			result.IssueYear = fields[1]
		case "eyr":
			result.ExpirationYear = fields[1]
		case "hgt":
			result.Height = fields[1]
		case "ecl":
			result.EyeColor = fields[1]
		case "hcl":
			result.HairColor = fields[1]
		case "pid":
			result.PassportID = fields[1]
		case "cid":
			result.CountryID = fields[1]
		default:
			log.Printf("unknown key %s in token %s", fields[0], token)
		}
	}
	return result
}
