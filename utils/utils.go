package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	Check(err)
	return intValue
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(fileName string) []string {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}

func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	Check(err)
	return string(file)
}

func SplitStringInLines(value string) []string {
	return strings.Split(value, "\n")
}
