package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http server start!")
	m:=new(MyHandler)
	http.HandleFunc("/test",m.RspHTTP)
	http.ListenAndServe(":1111",nil)

}
