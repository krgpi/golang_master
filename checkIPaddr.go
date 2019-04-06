package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const postURL = "https://hooks.slack.com/services/---------"

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {
  filename := "file/to/path"
	//古いIPアドレスの読み込み
	fp, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	str, err := Readln(reader)
	if err != nil {
		log.Fatal(err)
	}
	old := string(str)

	//現在のIPアドレスの取得
	out, err := exec.Command("curl", "ifconfig.io").Output()
	out = out[:len(out)-1] //末尾の改行コードを削除

	//現在のIPアドレスを上書き
	fp, err = os.Create("./ipaddr.txt")
	defer fp.Close()
	writer := bufio.NewWriter(fp)
	_, err = writer.WriteString(string(out))
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()

	//IPアドレスが変更されてたらSlackに投げる
	if string(out) != old {
		param := string("{\"text\":\"" + string(out) + "\"}")
		req, err := http.NewRequest("POST", postURL, bytes.NewBufferString(param))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded") //リクエストに必要なヘッダーを追加
		//作成したリクエストを確認
		// fmt.Println("Request:")
		// fmt.Println(req)

		//リクエストをPOST
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response)
	}
}
