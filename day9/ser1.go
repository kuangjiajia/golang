package main

//import (
//	"net/http"
//	"net/http/httputil"
//	"fmt"
//)
//
//func main() {
//	res,err := http.Get("http://www.imooc.com")
//	if err != nil {
//		panic(err)
//	}
//	defer res.Body.Close()
//
//	s,err := httputil.DumpResponse(res,true)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%s\n",s)
//}