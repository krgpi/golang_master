package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func sort(str string) (ary []string) {
	ary = strings.Split(str, "")
	for index := range ary {
		if index > 0 {
			for i := index; ; i = i - 1 {
				if i > 0 {
					cur, _ := strconv.Atoi(ary[i])
					bef, _ := strconv.Atoi(ary[i-1])
					if bef > cur {
						save := ary[i-1]
						ary[i-1] = ary[i]
						ary[i] = save
					} else {
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
	flag.Parse()
	var args []string = flag.Args()
	fmt.Println(sort(args[0]))
}
