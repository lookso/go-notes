package main

import (
	"context"
	"encoding/json"
	"fmt"
	//"git.100tal.com/wangxiao_go_lib/xesTools/httputil"
	"github.com/spf13/cast"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Param struct {
	WxUserid    string `json:"wx_userid"`
	MsgType     int    `json:"msg_type"`
	List        []List `json:"list"`
	StartTime   string `json:"start_time"`
	ValidType   int    `json:"valid_type"`
	CallbackURL string `json:"callback_url"`
}

type Content struct {
	Type     int    `json:"type"`
	Content  string `json:"content"`
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

type List struct {
	Content        []Content `json:"content"`
	ExternalUserid string    `json:"external_userid"`
}

func main() {
	isGenerated := cast.ToBool(cast.ToString("true"))
	if !isGenerated {
		fmt.Println("this is false")
	}
	var info interface{} = 123
	fmt.Println(cast.ToInt("123"))
	fmt.Println(cast.ToString(info))

	var sid = []interface{}{1, 2.0, "gocn"}
	fmt.Println("req.StuIds", sid)
	fmt.Println("cast.ToStringSlice(req.StuIds)", cast.ToStringSlice(sid))
	fmt.Println("--------------------------------------------")

	var pStr = `{
    "wx_userid": "tangtianlai@77198",
    "msg_type": 2,
   	"list": [
	{
		"content": [
		{
			"type": 3,
			"content": "this is wenben",
			"url": "http://yapi.xesv5.com/project/757/interface/api/17686",
			"filename": "Lorem"
		}
	],
		"external_userid": "wmmDoiDgAAlHALeoIi53PAzxp7UfXYwA"
	}
],
    "start_time": "",
    "valid_type": 1,
    "callback_url": ""
}`
	var p Param
	if err := json.Unmarshal([]byte(pStr), &p); err != nil {
		fmt.Println("err", err)
	}


	//bt, err := json.Marshal(&p)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("btstr", string(bt))
	//fmt.Println("WxUserid", p.WxUserid)
	ctx := context.Background()
	reqHeader := map[string]string{}
	reqHeader["Authorization"] = "-7Cp4u2qTTO5nwmbS4B7QgGx4UErg6Ko0bJCfSmGQ3qq5BVLEp5Mdr-u1taKcQ6Y5lumYF4bK-rMp-E8MS151g"
	client := &http.Client{}
	b, err := PostParam(ctx, "https://scrm.chengjiukehu.com", "/openapi/v2/message/batchusersendmsg", p, reqHeader, client)
	fmt.Println("res", b, err)
}

func PostParam(ctx context.Context, domain, path string, req interface{}, header map[string]string, client *http.Client) (ret []byte, err error) {
	header["Content-Type"] = "application/x-www-form-urlencoded"
	params := setQueryValues(req)
	fmt.Println("params", params)
	//ret, err = httputil.PostX(ctx, domain+path, params, header, client)
	//fmt.Println("ret", ret)

	return
}

func setQueryValues(ifc interface{}) url.Values {

	values := url.Values{}
	if mValues, ok := ifc.(url.Values); ok {
		for k := range mValues {
			values.Set(k, mValues.Get(k))
		}
	} else {
		elem := reflect.ValueOf(ifc)
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			kind := field.Kind()
			fmt.Println("f",field,kind)
			if (kind == reflect.Ptr || kind == reflect.Array || kind == reflect.Map || kind == reflect.Chan) && field.IsNil() {
				continue
			}
			tag := elem.Type().Field(i).Tag.Get("ArgName")
			if tag == "" {
				tag = elem.Type().Field(i).Name
			}
			arr := strings.Split(tag, ",")
			name := arr[0]
			empty_flag := false //如果值为空 且not_require ，则不放在valuse 里面
			auto_join := false  //如果[]string 且auto_join,则自动按照 ，join成字符串
			if len(arr) >= 2 {
				for i := 1; i < len(arr); i++ {
					if arr[i] == "not_require" {
						empty_flag = true
					}
					if arr[i] == "auto_join" {
						auto_join = true
					}
				}
			}
			switch kind {
			case reflect.Slice:
				is_empty := field.Len() == 0
				if is_empty == false || empty_flag == false {
					var str_arr []string
					for i := 0; i < field.Len(); i++ {
						value, _ := getQueryValue(field.Index(i))
						str_arr = append(str_arr, value)
					}
					if auto_join {
						values.Set(name, strings.Join(str_arr, ","))
					} else {
						for _, s := range str_arr {
							values.Add(fmt.Sprintf("%s[]", name), s)
						}
					}
				}
			default:
				value, is_empty := getQueryValue(field)
				if is_empty == false || empty_flag == false {
					values.Set(name, value)
				}
			}
		}
	}
	fmt.Println("values",values)
	return values
}
func getQueryValue(field reflect.Value) (string, bool) {
	kind := field.Kind()
	var value string
	var is_empty bool
	if (kind == reflect.Ptr || kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan) && field.IsNil() {
		return value, is_empty
	}
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := field.Int()
		value = strconv.FormatInt(i, 10)
		is_empty = i == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := field.Uint()
		value = strconv.FormatUint(u, 10)
		is_empty = u == 0
	case reflect.Float32:
		value = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		is_empty = field.Float() == 0
	case reflect.Float64:
		value = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		is_empty = field.Float() == 0
	case reflect.Bool:
		value = strconv.FormatBool(field.Bool())
		is_empty = field.Bool() == false
	case reflect.String:
		value = field.String()
		is_empty = value == ""
	}
	return value, is_empty
}
