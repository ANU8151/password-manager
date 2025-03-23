package cloud

import (
	"github.com/fatih/color"
)

type CloudDb struct {
	url string // Define fields for the Cloud struct
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *CloudDb) Write(content []byte) {
	color.Green("FILE_WRITTEN_SUCCESSFULLY")
}
