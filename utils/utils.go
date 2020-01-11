package utils

import (
	"os"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return string(""), err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return string(""), err
	}
	return string(content), nil
}

func WriteFile(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}
