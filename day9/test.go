package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	res,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",res.StatusCode)
		return
	}

	uf8Reader := transform.NewReader(res.Body, simplifiedchinese.GBK.NewDecoder()) //gbk转为UTF8
	all,err := ioutil.ReadAll(uf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",all)
}
