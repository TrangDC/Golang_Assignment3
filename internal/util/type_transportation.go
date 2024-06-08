package util

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func SwitchType(typeString string, typeOfValue any) string {
	value := ""
	switch typeString {
	case "uuid.UUID":
		if typeOfValue.(uuid.UUID) == uuid.Nil {
			value = ""
		} else {
			value = typeOfValue.(uuid.UUID).String()
		}
	case "string":
		value = typeOfValue.(string)
	case "time.Time":
		value = typeOfValue.(time.Time).UTC().Format("2006-01-02")
	case "int":
		value = strconv.Itoa(typeOfValue.(int))
	case "int64":
		value = strconv.FormatInt(typeOfValue.(int64), 10)
	case "float64":
		value = strconv.FormatFloat(typeOfValue.(float64), 'f', 6, 64)
	case "bool":
		value = strconv.FormatBool(typeOfValue.(bool))
	default:
		value = fmt.Sprintf("%v", typeOfValue)
	}
	return value
}

func ValidDuplicateString(values []string) error {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		if _, ok := uniqueValues[value]; ok {
			return fmt.Errorf("duplicate value %s", value)
		}
		uniqueValues[value] = true
	}
	return nil
}

func ExtractKeys(input string) map[string]float64 {
	re := regexp.MustCompile(`{{\s*([\w:]+)\s*}}`)
	matches := re.FindAllStringSubmatch(input, -1)

	result := make(map[string]float64)
	for _, match := range matches {
		key := match[1]
		result[key] = 0
	}
	return result
}
