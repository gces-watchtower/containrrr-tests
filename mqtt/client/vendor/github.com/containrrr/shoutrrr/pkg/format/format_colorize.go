package format

import "github.com/fatih/color"

// ColorizeDesc colorizes the input string as "Description"
var ColorizeDesc = color.New(color.FgHiBlack).SprintFunc()

// ColorizeTrue colorizes the input string as "True"
var ColorizeTrue = color.New(color.FgHiGreen).SprintFunc()

// ColorizeFalse colorizes the input string as "False"
var ColorizeFalse = color.New(color.FgHiRed).SprintFunc()

// ColorizeNumber colorizes the input string as "Number"
var ColorizeNumber = color.New(color.FgHiBlue).SprintFunc()

// ColorizeString colorizes the input string as "String"
var ColorizeString = color.New(color.FgHiYellow).SprintFunc()

// ColorizeEnum colorizes the input string as "Enum"
var ColorizeEnum = color.New(color.FgHiCyan).SprintFunc()

// ColorizeValue colorizes the input string according to what type appears to be
func ColorizeValue(value string, isEnum bool) string {
	if isEnum {
		return ColorizeEnum(value)
	}

	if isTrue, isType := ParseBool(value, false); isType {
		if isTrue {
			return ColorizeTrue(value)
		}
		return ColorizeFalse(value)
	}

	if IsNumber(value) {
		return ColorizeNumber(value)
	}

	return ColorizeString(value)
}
