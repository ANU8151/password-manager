package cloud

import "github.com/ANU8151/password-manager/output"

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
	output.PrintError("FILE_WRITTEN_SUCCESSFULLY")
}
