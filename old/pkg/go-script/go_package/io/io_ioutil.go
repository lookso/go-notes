/*
@Time : 2019/3/24 10:00 AM
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	dir := os.Args[1]
	listAll(dir,0)
}

func listAll(path string, curHier int){
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil{fmt.Println(err); return}

	for _, info := range fileInfos{
		if info.IsDir(){
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(),"\\")
			listAll(path + "\\" + info.Name(),curHier + 1)
		}else{
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

