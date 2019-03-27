package main

import "fmt"

func hoge() string, error {
	ans, err := somefunc(aa)
	if err!= nil{
		return nil, err
	}
	return ans, nil
}

func main() string {
	modorichi, err := hoge()
	if err != nil{
		fmt.Println("load error")
		return nil
	}
	return modorichi
}
