package utils

import (
	//"fmt"
	"bufio"
	"os"
)

func NewStringSliceFromFileName(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	xs := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xs = append(xs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return xs, nil
}
