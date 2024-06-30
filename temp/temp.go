package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

func main()  {
	standardAnswer := "好的##答案##123"
	replacer := strings.NewReplacer("＃＃答案＃＃", "", "##答案##", "", "#答案#", "", "# # 答案 ##", "", "##答案", "")
	newStandardAnswer := replacer.Replace(standardAnswer)
	fmt.Println(newStandardAnswer)
	videoStr:=`{"aigc_cluster_videos":{"check_list":[{"video_url":"https://chat.xuepaipai.com/#/handwriting?question_id=20000154"}],"cover_txt":"","cover_url":"","math_qt":2,"type":1,"video_url":"https://chat.xuepaipai.com/#/handwriting?question_id=20000154"},"source":"","task_id":0}`
	s,err:=filterQuestion(videoStr)
	fmt.Println(s,err)

}


func filterQuestion(videoStr string) (isFilter bool, err error) {

	var video map[string]interface{}
	err = json.Unmarshal([]byte(videoStr), &video)
	if err != nil {
		return false, fmt.Errorf("videoData-data unmarshal err: %v", err.Error())
	}
	if _, ok := video["aigc_cluster_videos"]; ok {
		return true, nil
	}
	videos, vok := video["videos"]
	source, sok := video["source"]
	if sok && vok && len(cast.ToSlice(videos)) > 0 && (source == "zb" || source == "mc") {
		isFilter = true
	}
	//10000001


	return
}
