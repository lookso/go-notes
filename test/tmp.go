package main

import (
	"encoding/json"
	"fmt"
)

type GroupOperationCreateGroupResponse struct {
	AddTime                   int          `json:"add_time"`
	ChatroomDataFlag          int          `json:"chatroom_data_flag"`
	ChatroomLocalVersion      int          `json:"chatroom_local_version"`
	ChatroomNickname          string       `json:"chatroom_nickname"`
	ChatroomNoticePublishTime int          `json:"chatroom_notice_publish_time"`
	ChatroomVersion           int          `json:"chatroom_version"`
	DisplayName               string       `json:"display_name"`
	IsShowname                int          `json:"is_showname"`
	JoinChatroomCheck         bool         `json:"join_chatroom_check"`
	JoinChatroomCheckSwitch   bool         `json:"join_chatroom_check_switch"`
	MemberList                []MemberList `json:"member_list"`
	ModifyTime                int          `json:"modify_time"`
	RoomFlag                  int          `json:"room_flag"`
	Style                     int          `json:"style"`
	WebID                     string       `json:"web_id"`
}

type MemberList struct {
	Alias      string `json:"alias"`
	Nickname   string `json:"nickname"`
	RemarkName string `json:"remark_name"`
	Username   string `json:"username"`
}

var str = `{
    "add_time": 0,
    "chatroom_data_flag": 0,
    "chatroom_local_version": 0,
    "chatroom_nickname": "矩阵微信测试abc",
    "chatroom_notice_publish_time": 0,
    "chatroom_version": 0,
    "display_name": "",
    "is_showname": 0,
    "join_chatroom_check": false,
    "join_chatroom_check_switch": false,
    "member_list": [
        {
            "alias": "",
            "nickname": "",
            "remark_name": "",
            "username": "wxid_wtz7u4rgjow822"
        },
        {
            "alias": "",
            "nickname": "",
            "remark_name": "",
            "username": "wxid_kscqeayousrv22"
        },
        {
            "alias": "",
            "nickname": "",
            "remark_name": "",
            "username": "wxid_0a63hhyrw7g321"
        }
    ],
    "modify_time": 0,
    "room_flag": 0,
    "style": 0,
    "web_id": "1621568896548494139"
}`

func main() {

	var group = GroupOperationCreateGroupResponse{}
	err := json.Unmarshal([]byte(str), &group)
	fmt.Println("group", group)
	fmt.Println("err", err)
}
