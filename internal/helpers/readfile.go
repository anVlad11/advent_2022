package helpers

import (
	"bufio"
	"os"
)

func ReadFile(filename string) ([]string, error) {
	var err error
	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}

	if file == nil {
		return nil, nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	content := []string{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return content, nil
}
