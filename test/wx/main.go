package main


import (
	"net/url"
)

func ParseXml(xml []byte) url.Values {
	xml = xml[5 : len(xml)-6] // 剔除 <xml></xml>
	params := url.Values{}
	for _, a := range reg.FindAllSubmatch(xml, -1) {
		// 去掉 <![CDATA[JSAPI]]> 的头尾
		if string(a[2][:9]) == "<![CDATA[" {
			a[2] = a[2][9 : len(a[2])-3]
		}
		params.Set(string(a[1]), string(a[2]))
	}
	return params
}
