package output

import "github.com/fatih/color"

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Blue("Error Code: %d", intValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		color.Green(stringValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Unknown Error")
}
