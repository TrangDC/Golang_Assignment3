package util

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

var outputString []string

func ProcessIfStatement(input string) string {
	ifStatementArr := splitIfStatement(input)
	for _, ifStatement := range ifStatementArr {
		stringAfterReplace := replaceIfStatement(ifStatement)
		input = strings.ReplaceAll(input, ifStatement, stringAfterReplace)
		outputString = []string{}
	}
	return input
}

func replaceIfStatement(inputString string) string {
	processString(inputString)
	result := ""
	outputString = lo.Reverse(outputString)
	outputStringCopy := make([]string, len(outputString))
	copy(outputStringCopy, outputString)
	for i := 0; i < len(outputString); i++ {
		arr := strings.Split(string(outputStringCopy[i]), ";")
		newString := fmt.Sprintf("%s?%s:%s", arr[0], arr[1], arr[2])
		if i == len(outputString)-1 {
			result = newString
			break
		}
		outputStringCopy[i+1] = strings.ReplaceAll(string(outputStringCopy[i+1]), string(outputString[i]), newString)
	}
	result = strings.ReplaceAll(result, "if", "")
	result = result[:len(result)-1]
	return result
}

func splitIfStatement(inputString string) []string {
	ifStatementArr := []string{}
	result := ""
	openBrackets := 0
	for i := 0; i <= len(inputString); i++ {
		replace := ""
		if openBrackets == 0 {
			if i+4 <= len(inputString) && string(inputString[i]) == "(" && string(inputString[i:i+4]) == "(if(" {
				openBrackets += 2
				replace = "(if("
				i += 3
			}
		} else {
			replace = string(inputString[i])
			if string(inputString[i]) == "(" {
				openBrackets += 1
			}
			if string(inputString[i]) == ")" && openBrackets == 1 {
				result += replace
				openBrackets -= 1
				ifStatementArr = append(ifStatementArr, result)
				result = ""
				continue
			}
			if string(inputString[i]) == ")" {
				openBrackets -= 1
			}
		}
		result += replace
	}
	return ifStatementArr
}

func processString(inputString string) {
	result := ""
	openBrackets := 0
	for i := 0; i <= len(inputString); i++ {
		replace := ""
		if openBrackets == 0 {
			if i+4 <= len(inputString) && string(inputString[i]) == "(" && string(inputString[i:i+4]) == "(if(" {
				openBrackets += 2
				replace = "if("
				i += 3
			}
		} else {
			replace = string(inputString[i])
			if string(inputString[i]) == "(" {
				openBrackets += 1
			}
			if string(inputString[i]) == ")" {
				openBrackets -= 1
			}
		}
		result += replace
	}
	outputString = append(outputString, result)
	if strings.Contains(result, "(if(") {
		processString(result)
	}
}
