package main

import "fmt"

func main() {
	var array [5]int //要素数を最初に決める（ここでは5つ）固定長
	array[3] = 3
	fmt.Printf("array:%v", array)
	var slice []int // 要素数は定義しない 可変長
	count := 10
	for i := 0; i < count; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("Slice:%v/n", slice)

}
