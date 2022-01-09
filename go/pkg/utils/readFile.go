package utils

import (
	//"fmt"
	"bufio"
	"path/filepath"
	"os"
)

func NewStringSliceFromFileName(fileName string) ([]string, error) {
	filePath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "DePinto", "aoc-2021", "input", fileName)
	
	file, err := os.Open(filePath)
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
