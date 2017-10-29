package core

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Save metrics in a file
func Save(path, name, metrics string) error {
	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		return err
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/%s", path, name), []byte(metrics), os.ModePerm)
}
