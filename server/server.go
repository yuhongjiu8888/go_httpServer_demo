package main

import (
	"fmt"
	"net/http"
	"strconv"
)


/*解决跨域*/
func CrossAdmin(w http.ResponseWriter,r *http.Request){
	orgin := r.Header.Get("Access-Control-Allow-Origin")
	if orgin == "" {
		orgin = r.Header.Get("Origin")
	}
	w.Header().Set("Access-Control-Allow-Origin", orgin)
	w.Header().Add("Access-Control-Allow-Headers",  "X-Requested-With,Content-Type,x-csrftoken")    //header的类型
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.Header().Set("content-type", "application/json;charset=UTF-8") //返回数据格式是json
}

type MyHandler struct {
	Rspcode string `json:"rspcode"` 
	Rspinfo string	`json:"rspinfo"`
}
/*自定义serverHTTP*/
func  (m *MyHandler)RspHTTP(w http.ResponseWriter, r *http.Request) {
	CrossAdmin(w,r)
	mm:=GetBody(r)

	var respone string
	k, _ := strconv.Atoi("18")
	age,_:= strconv.Atoi(mm["age"])
	if  age>= k {
		respone=m.encode("200","OK")
	} else{
		respone = m.encode("-100", "年龄非法，需成年18岁以上！")
	}
	fmt.Println("return:",respone)
	fmt.Fprintf(w,respone)
	GetBody2(r)
}

/*获取请求报文体*/
func GetBody(r *http.Request) map[string]string{
	r.ParseForm()
	//return r.Form
	m_map:=make(map[string]string)
	if len(r.Form) > 0 {
		for k,v := range r.Form {
			fmt.Printf("%s=%s\n", k, v[0])
			m_map[k] = v[0]
		}

	}
	return m_map
}

func GetBody2(r *http.Request)  {
	if (r.Form.Get("name") == "yuwei") && (r.Form.Get("age") == "18") {
		fmt.Println("hello,验证成功！")
	} else {
		fmt.Println("hello,验证失败了！")
	}
}