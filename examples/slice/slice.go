package slice

import (
	"bufio"
	"io/fs"
	"strconv"
)

// GetFloats 使用切片处理不确定数量的情况
//程序出错时，不应返回无效数据，切片通常返回nil,通常返回对应类型的零值
func GetFloats(file fs.File) ([]float64, error) {
	var nums []float64
	scanner := bufio.NewScanner(file)
	i := 0
	var err error
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, f)
		i++
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return nums, nil
}
