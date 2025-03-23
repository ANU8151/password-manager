package files

import (
	"os"

	"github.com/ANU8151/password-manager/output"
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
		output.PrintError("ERROR_CREATING_FILE")
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError("ERROR_WRITING_TO_FILE")
		return
	}
	output.PrintError("FILE_WRITTEN_SUCCESSFULLY")
}
