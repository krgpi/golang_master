package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//sort バブルソート
func sort(str string) (ary []string) {
	//1文字ずつ分割して配列にする
	ary = strings.Split(str, "")
	//ソート開始
	for index := range ary {
		if index > 0 {
			for i := index; ; i = i - 1 {
				if i > 0 {
					cur, _ := strconv.Atoi(ary[i])   //current
					bef, _ := strconv.Atoi(ary[i-1]) //before
					if bef > cur {                   //小さい順でなかったら入れ替える
						save := ary[i-1]
						ary[i-1] = ary[i]
						ary[i] = save
					} else { //小さい順になってたら何もしない
						break
					}
				} else {
					break
				}
			}
		}
	}
	return
}

func main() {
	//コマンドライン引数を取得
	flag.Parse()
	var args []string = flag.Args()
	fmt.Println(sort(args[0]))
}
