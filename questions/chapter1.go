package questions

import (
	"fmt"
	"math"
	"strings"
)

func isUnique(input string) bool {
	charMap := make(map[string]int)
	for _, char := range input {
		if _, ok := charMap[string(char)]; ok {
			return false
		}
		charMap[string(char)] = 1
	}
	return true
}

func checkPurmutations(input1, input2 string) bool {
	inputMap1 := make(map[string]int)
	inputMap2 := make(map[string]int)

	for _, char := range input1 {
		if _, ok := inputMap1[string(char)]; ok {
			inputMap1[string(char)]++
		}
		inputMap1[string(char)] = 1
	}

	for _, char := range input2 {
		if _, ok := inputMap2[string(char)]; ok {
			inputMap2[string(char)]++
		}
		inputMap2[string(char)] = 1
	}

	for key, value := range inputMap1 {
		if val, ok := inputMap2[key]; !ok || val != value {
			return false
		}
	}

	return true
}

func urlify(input string) string {
	var stringBuilder strings.Builder

	for i := range input {
		if string(input[i]) != " " {
			stringBuilder.WriteByte(input[i])
		} else {
			stringBuilder.WriteString("%20")
		}
	}

	return stringBuilder.String()
}

func palindromePermutation(input string) bool {
	// Ok so the idea here is that if len(input) > 1 then there can only
	// be one odd count of characters.
	// I'll also ignore spaces.
	if len(input) < 1 {
		return true
	}

	input = strings.ToLower(input)

	charMap := make(map[string]int)
	for _, char := range input {
		myChar := string(char)
		if myChar != " " {
			if _, ok := charMap[myChar]; ok {
				charMap[myChar]++
			} else {
				charMap[myChar] = 1
			}
		}
	}

	oddCount := 0
	for _, value := range charMap {
		if value%2 != 0 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}

	return true
}

func oneAway(input, output string) bool {
	inputMap := make(map[rune]int)
	outputMap := make(map[rune]int)

	if math.Abs(float64(len(input))-float64(len(output))) > 1 {
		return false
	}

	errorToleration := true

	for _, char := range input {
		if _, ok := inputMap[char]; ok {
			inputMap[char]++
		} else {
			inputMap[char] = 1
		}
	}
	for _, char := range output {
		if _, ok := outputMap[char]; ok {
			outputMap[char]++
		} else {
			outputMap[char] = 1
		}
	}

	for key, value := range outputMap {
		if val, ok := inputMap[key]; !ok || val != value {
			if !errorToleration {
				return false
			}
			errorToleration = false
		}
	}

	return true
}

func stringCompression(input string) string {
	charCountMatrix := [][]byte{}

	currentChar := input[0]
	charCount := 1

	for _, char := range input[1:] {
		if char != rune(currentChar) {
			charCountMatrix = append(charCountMatrix, []byte{currentChar, byte(charCount)})
			currentChar = byte(char)
			charCount = 1
		} else {
			charCount++
		}
	}

	charCountMatrix = append(charCountMatrix, []byte{currentChar, byte(charCount)})

	var stringBuilder strings.Builder

	for _, charCount := range charCountMatrix {
		stringBuilder.WriteByte(charCount[0])
		stringBuilder.WriteString(fmt.Sprintf("%d", charCount[1]))
	}

	return stringBuilder.String()
}
