package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	s := `{"errcode":1,"errmsg":12,"data323":"cd"}`
	ret, err := new(SCrmOpenApiErr).CheckErrCode([]byte(s))
	fmt.Println(ret, err.Error())
}

type SCrmOpenApiErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (this *SCrmOpenApiErr) NewError() SCrmOpenApiErr {
	return SCrmOpenApiErr{
		Code: 2,
		Msg:  "23",
	}
}

func (this *SCrmOpenApiErr) Error() string {
	return fmt.Sprintf("%s-%s", cast.ToString(this.Code), this.Msg)
}
func (this *SCrmOpenApiErr) CheckErrCode(data []byte) (ret []byte, err error) {
	type BasicResp struct {
		Code int         `json:"errcode"`
		Msg  string      `json:"errmsg"`
		Data interface{} `json:"data1"`
	}
	var resp BasicResp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil,errors.New("5555")
	}
	if resp.Code > 0 {
		this.Code = resp.Code
		this.Msg = resp.Msg
		return ret, this
	}
	ret, err = json.Marshal(resp.Data)
	return
}
