package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile() {

}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	color.Green("File written successfully!")
}
