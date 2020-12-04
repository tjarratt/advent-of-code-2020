package passport

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

type validator struct {
	passports []Passport
}

type Validator interface {
	Valid() []Passport
	Invalid() []Passport
}

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportId     string
	CountryId      string
}

func PassportValidator(reader io.Reader) Validator {
	return &validator{passports: parse(reader)}
}

func (v *validator) Valid() []Passport {
	valid := []Passport{}
	for _, passport := range v.passports {
		if isValid(passport) {
			valid = append(valid, passport)
		}
	}

	return valid
}

func (v *validator) Invalid() []Passport {
	invalid := []Passport{}

	for _, passport := range v.passports {
		if isValid(passport) {
			continue
		}

		invalid = append(invalid, passport)
	}

	return invalid
}

func isValid(passport Passport) bool {
	if passport.BirthYear == "" ||
		passport.IssueYear == "" ||
		passport.ExpirationYear == "" ||
		passport.Height == "" ||
		passport.HairColor == "" ||
		passport.EyeColor == "" ||
		passport.PassportId == "" {
		return false
	}

	return true
}

func parse(reader io.Reader) []Passport {
	passports := []Passport{}

	input, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var passport = Passport{}
	re := regexp.MustCompile("([a-z]+:[#a-zA-Z0-9]+)")

	for _, line := range lines {
		if line == "\n" || line == "" {
			passports = append(passports, passport)
			passport = Passport{}
			continue
		}

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			pieces := strings.Split(match[0], ":")
			switch pieces[0] {
			case "ecl":
				passport.EyeColor = pieces[1]
			case "pid":
				passport.PassportId = pieces[1]
			case "eyr":
				passport.ExpirationYear = pieces[1]
			case "hcl":
				passport.HairColor = pieces[1]
			case "byr":
				passport.BirthYear = pieces[1]
			case "iyr":
				passport.IssueYear = pieces[1]
			case "cid":
				passport.CountryId = pieces[1]
			case "hgt":
				passport.Height = pieces[1]
			default:
				panic(fmt.Sprintf("Strange unknown passport field '%s'", pieces[0]))
			}
		}
	}

	return passports
}
