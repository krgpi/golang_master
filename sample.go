//
package main

import (
	"fmt"
	"math"
)

func cal(a int, b int) (sum int, dif int) {
	sum = a + b
	dif = a - b
	return sum, dif
}

func main() {
	var r float64 = 2         //可変
	const pi float64 = 3.1415 //固定値 （実際のところ円周率はmath.Piで呼び出す）

	var s = func(radius float64) float64 { //クロージャー
		return math.Pow(radius, 2) * pi
	}
	defer println(s(r)) //defer スタックに回され、関数の最後に実行

	resSum, resDif := cal(int(r), int(math.Sqrt(r)))
	fmt.Printf("%[1]d + %[2]d = %[3]d, %[1]d - %[2]d = %[4]d\n", int(r), int(math.Sqrt(r)), resSum, resDif)

	var array [5]int //要素数を最初に決める（ここでは5つ）固定長
	array[3] = 3
	fmt.Printf("array:%v\n", array) //初期値を決めなくても、0が自動的に初期値になっている

	var slice []int // 要素数は定義しない 可変長の配列

	count := 10 //var count int = 10 の省略形 型推論される
	for i := 0; i < count; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("slice:%v\n", slice)

}
