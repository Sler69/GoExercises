package main

import "strings"

var romanStringToInt = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900}

var romanStrings = [13]string{
	"IV",
	"IX",
	"XL",
	"XC",
	"CD",
	"CM",
	"I",
	"V",
	"X",
	"L",
	"C",
	"D",
	"M"}

func main() {
	var value string = "MCMXCIV"
	var result int = romanToInteger(value)
}

func romanToInteger(value string) int {
	var result int = 0
	var index int = 0
	var lengthOfRomans int = len(romanStrings)
	for index < lengthOfRomans || value != "" {
		var romanStringPriority string = romanStrings[index]
		if strings.Contains(value, romanStringPriority) {
			value = strings.Replace(value, romanStringPriority, "", 1)
			result = result + romanStringToInt[romanStringPriority]
		} else {
			index++
		}
	}
}
