package util

import (
	"bytes"
	"io"
	"os"
)

const (
	defaultPerm = 0666
)

func LoadFile(filePath string) (string, error) {
	writer := new(bytes.Buffer)

	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	io.Copy(writer, f)

	return writer.String(), nil
}

func CreateFile(filePath string, data string) error {
	reader := new(bytes.Buffer)
	_, err := reader.WriteString(data)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, defaultPerm)
	if err != nil {
		return err
	}

	io.Copy(f, reader)

	return nil
}
