package calcengine

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type FileReader struct {
}

func (r *FileReader) Read(filename string) ([]int, error) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
