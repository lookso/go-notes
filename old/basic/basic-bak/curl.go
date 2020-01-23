package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var group sync.WaitGroup
	for i := 1; i < 1000; i++ {
		group.Add(1)
		go getCurl(&group)
	}
	group.Wait()
}

func getCurl(g *sync.WaitGroup) {

	url := "https://mobile.yangkeduo.com/proxy/api/api/george/resource_goods/query_list?page=10&size=200&subject_id=25309&mall_id=0&scene_name=17464&refer_page_sn=15002&list_id=p-zBoB74BAypromotion_op_k5bdCC&field_types=&resource_type=114&cart_id=60446&anti_content=0anAfan5NOcYq9ECsy7jo8hC4xg1U3Xatb5fosFk4H_psjyHvz5b_ex6Buf3aqkc4cKaGhoznqoGF5vU6IaMLRRy9X9o9hTV8OKymbEF_DCGyaFW8f5-hk03beyd-5Futf3z1FYzqcxGOy-dbAytAS1IVHMn37ieGKBwzaTK266MHm6BbzKpuxPfadI6Duvs9EVSrWK_KT1_5kZ_-lKFD4lvXuylJSo5gmaEpoErT2LWcVLrpeRyQHWR-_m6XnjoTgA6tGXMVmsoThTDKlFU1ENpQT9A9vy8CWNu8yorVcrCEv3pEFQWXJrAxPVMGQ8RymxqZCsKq2QE9i9OL0dZtkEeyCw3SYgxxoz15TGwS64CeTx_mgf7XOAeryAz6XtW3JVCl97nt1DyhO4kigx23QVvXZZXMzDWYp7FvgkjT_3Wd2qw62XhfZWSmc-YA1L_ZksmwMy6mIwmQfynIYbwLKEtWwsxvvNVHkbRvAbQMAoToXmlS_lYPQPateJkTQMHBc6rxeukMaMfvqXyLncQS3hb7yjs7bLqcDIeuVjjzAlUb-FOikELZSjzyK3_au8ZYh_7kycuIWQNLh6zwnXRCBu_tcp&pdduid=0"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "mobile.yangkeduo.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Cookie", "api_uid=CiFs3l1548ILbABABxnsAg==")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	content := UseNewEncoder(string(body), "gbk","utf8")

	fmt.Println(content)
	//fmt.Println(string(body))
	//fmt.Println(base64.StdEncoding.EncodeToString(body))
	fmt.Println("------------------------------------------------------")
	defer g.Done()
}
func UseNewEncoder(src string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}
