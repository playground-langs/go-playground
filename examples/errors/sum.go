package errors

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("opening ", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("closing file ", file.Name())
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func GetFloats(fileName string) ([]float64, error) {
	var nums []float64
	file, err := OpenFile(fileName)
	defer CloseFile(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, f)
		i++
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return nums, nil
}
