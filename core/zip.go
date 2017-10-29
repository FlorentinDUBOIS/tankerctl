package core

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// ReadAndUnzip read a http response and unzip the response body
func ReadAndUnzip(r *http.Response) ([][]byte, error) {
	defer r.Body.Close()

	log.Info("Read zip archive")
	archive, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return nil, err
	}

	log.Info("Unzip archive")
	return Unzip(archive, size)
}

// Unzip an bytes archive into an array of bytes files
func Unzip(archive []byte, size int64) ([][]byte, error) {

	log.Info("Read archive")
	reader, err := zip.NewReader(bytes.NewReader(archive), size)
	if err != nil {
		return nil, err
	}

	files := make([][]byte, 0)
	for _, file := range reader.File {

		log.Infof("Open file %s", file.Name)
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		log.Info("Extract file")
		c, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}

		files = append(files, c)
	}

	log.Info("Ending unzipping")
	return files, nil
}
