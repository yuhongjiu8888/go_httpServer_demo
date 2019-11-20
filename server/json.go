package main

import (
	"encoding/json"
)

/*返回response*/
func (rsp *MyHandler) encode(status,mes string) string{
	rsp.Rspcode = status
	rsp.Rspinfo = mes

	j,_:=json.Marshal(rsp)
	return string(j)
}
