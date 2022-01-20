package main

import (
	"fmt"
	"sort"
)

func main() {
	var ranks map[string]int
	//映射变量的零值为nil
	fmt.Println(ranks == nil)
	ranks = make(map[string]int)
	fmt.Println(ranks == nil)
	ranks["gold"] = 1
	ranks["silver"] = 2
	fmt.Println(ranks)
	m1 := make(map[string]int)
	fmt.Println(m1)
	//空映射
	m2 := map[int]int{}
	//映射字面量
	m3 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println(m2, m3)
	//映射值的零值为值类型的零值
	fmt.Println(m3["not found"]) //0
	//如何区分零值的来源,使用第二个参数判断是否存储过
	v, ok := m3["not found"]
	fmt.Println(v, ok)
	_, exist := m3["test exist"]
	fmt.Println(exist)
	//删除映射
	delete(m3, "a")
	fmt.Println(m3["a"]) //0
	//iterate k,v
	for k, val := range m3 {
		fmt.Println(k, val)
	}
	//only keys,unordered
	for k := range m3 {
		fmt.Println(k)
	}
	OrderedIterate(map[string]int{"dd": 133, "bs": 12, "ad": 15})
}

// OrderedIterate 可以借助切片先做key排序，再实现有序遍历
func OrderedIterate(m map[string]int) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, ":", m[v])
	}
}
