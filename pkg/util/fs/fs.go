package fs

import (
	"io/ioutil"
	"os"
)

// Exists checks if file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Mkdir creates directory if it does not exist
func Mkdir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

// WriteFile saves data to file. If the file does not
// exist, it will create a file and write data to it.
func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
