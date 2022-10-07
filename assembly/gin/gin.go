package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func init() {
	fmt.Println(312321321)
	//time.Sleep(5 * time.Second)
}

type Param struct {
	Name string `json:"name"`
	Skip []int  `json:"skip"`
}
type response struct {
	Code    int         `json:"code"`
	Stat    int         `json:"stat"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Success(v interface{}) interface{} {
	ret := response{Stat: 1, Code: 0, Message: "ok", Data: v}
	return ret
}

func CheckSign(params url.Values) error {
	sign := cast.ToString(params["sign"])
	if sign == "" || sign != CreateMd5Sign(params) {
		return errors.New("签名错误")
	}
	return nil
}

// 创建签名
func CreateMd5Sign(params url.Values) string {
	// 自定义 MD5 组合
	return MD5("1231232132" + createEncryptStr(params))
}

func createEncryptStr(params url.Values) string {
	// 将请求参数的key提取出来，并排好序
	newKeys := make([]string, 0, 8)
	for k, _ := range params {
		newKeys = append(newKeys, k)
	}
	sort.Strings(newKeys)
	fmt.Println("newKeys", newKeys)
	var originStr = url.Values{}
	// 将输入进行标准化的处理
	for _, v := range newKeys {
		if v == "sign" {
			continue
		}
		originStr.Add(v, params.Get(v))
	}
	return originStr.Encode()
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

type User struct {
	IDs  []int  `json:"id"`
	Name string `json:"name"`
}

type TaskParam struct {
	u *User
}

func main() {
	a := 1
	if a == 1 {
		a = 2
	}
	fmt.Println(a)

	var param = Param{
		Skip: make([]int, 0),
	}
	gin.SetMode(gin.DebugMode) // 设置在gin.New()前面才有效

	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/post/data", func(c *gin.Context) {
		var values url.Values
		reqQuery := ""
		reqUri := c.Request.RequestURI
		idx := strings.Index(reqUri, "?")
		if idx > 0 {
			reqQuery = reqUri[idx+1:]
		}
		values, _ = url.ParseQuery(reqQuery)
		fmt.Println("v1", values)
		c.Request.ParseForm()
		values = c.Request.Form
		fmt.Println("v2", values)
		name,_:=c.GetQuery("name")
		fmt.Println("name",name)

		//var vv url.Values
		//var data =make(map[string]interface{})
		//bd, _ := ioutil.ReadAll(c.Request.Body)
		//// 再重新写回请求体body中，ioutil.ReadAll会清空c.Request.Body中的数据
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bd))
		//err:=json.Unmarshal(bd, &data)
		//fmt.Println("err",err)
		//for k, v := range data {
		//	vv[k] = v[0][cast.ToString()]
		//}
		//fmt.Println("v4",vv)

		c.JSON(http.StatusOK, Success(nil))
	})

	router.POST("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(http.StatusOK, "ok")
	})
	//router.GET("/post", LoggerMiddleWare(), func(c *gin.Context) {
	//	// 获取原始字节
	//	d, err := c.GetRawData()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	c.String(http.StatusOK, string(d)+":热编译123456")
	//})
	router.GET("/get/json", func(c *gin.Context) {

		type Resopnse struct {
			Users []User `json:"users"`
		}
		//var r = Resopnse{} // { "ids": null }
		var r = Resopnse{
			Users: make([]User, 0), // { "ids": []] }
		}
		name, ok := c.GetQuery("name")
		if !ok {
			return
		}
		u := new(User)
		u.IDs = []int{1, 2, 3}
		u.Name = name

		defer ants.Release()
		p, err := ants.NewPoolWithFunc(10, ActionDataBackUpFun, ants.WithPanicHandler(panicHandler))
		if err != nil {
			return
		}
		var taskParam = &TaskParam{
			u: u,
		}
		if err = p.Invoke(taskParam); err != nil {
			c.JSON(http.StatusBadGateway, Success(r))
			return
		}

		c.JSON(http.StatusOK, Success(r))
	})

	router.POST("/post/json", func(c *gin.Context) {
		// 获取原始字节
		err := c.ShouldBind(&param)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(param.Name)
		c.JSON(http.StatusOK, param)
	})
	router.HEAD("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(http.StatusOK, string(d))
	})
	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	router.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// 访问 http://localhost:8000/jsonp?callback=call
		// 将会输出:   call({foo:"bar"})
		c.JSONP(http.StatusOK, data)
	})

	router.Run(":8000")
}

func panicHandler(err interface{}) {
	fmt.Println(err)
}

func ActionDataBackUpFun(data interface{}) {
	if data == nil {
		return
	}
	taskParam := data.(*TaskParam)
	fmt.Println("taskParam", taskParam)
	u := taskParam.u
	fmt.Println(u)
}
