package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	color.Green("File written successfully!")
}
