package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{
		filename: filename,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	content, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
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
