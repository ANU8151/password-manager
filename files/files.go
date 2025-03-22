package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) {
	content, err := os.Open(name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer content.Close()

	fmt.Println(content)
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
