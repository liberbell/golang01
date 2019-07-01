package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseSignatureFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sigs := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		fileds := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", path, lnum)
		}
	}
}
