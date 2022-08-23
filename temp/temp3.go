package main

import (
	"encoding/json"
	"fmt"
)

var msg = `{
    "work_code": "188192",
    "wx_user_id": "wangjiaqi14@188192",
    "messages": [
        {
            "stu_list": [
                {
                    "stu_name": "段楚凡",
                    "stu_id": 77121223,
                    "external_user_id": "wmHqcDDgAAnDN-mNXsugAoeHab_t86Tw",
                    "identity": 8,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "段楚凡",
                    "stu_id": 77121223,
                    "external_user_id": "wmHqcDDgAABVXRqCy-99-ehkrF6GFcww",
                    "identity": 2,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "然然",
                    "stu_id": 59062717,
                    "external_user_id": "wmHqcDDgAAr6Xr36yePAmRwAyJhhSLsA",
                    "identity": 2,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "姜奕辰",
                    "stu_id": 20134448,
                    "external_user_id": "wmHqcDDgAAp493AbXzvKHRHNST2rH54w",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "周庆誉",
                    "stu_id": 77410340,
                    "external_user_id": "wmHqcDDgAA5Cz6b4-zBBTTC-7rBvXEhw",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "郭采萱郭俊奎",
                    "stu_id": 43335418,
                    "external_user_id": "wmHqcDDgAAylDf8fNdtStt0ndJyYFb4A",
                    "identity": 8,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "戚雨旸",
                    "stu_id": 77137150,
                    "external_user_id": "wmHqcDDgAAffoZQ0u1yLXusI3Ob3R9MQ",
                    "identity": 8,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "朱熠辰",
                    "stu_id": 62491717,
                    "external_user_id": "wmHqcDDgAAV-tRyDnIKbTv4wCqmmX1zQ",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "钟梓宸",
                    "stu_id": 53726362,
                    "external_user_id": "wmHqcDDgAAzE8dRvN5BFbkFCMDYjO06Q",
                    "identity": 8,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "萌萌",
                    "stu_id": 46719874,
                    "external_user_id": "wmHqcDDgAAdDISHK_JyvN2iIHl0YohOA",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "雨泽",
                    "stu_id": 11002762,
                    "external_user_id": "wmHqcDDgAAMVQqqUKwYC3Dc3tvJpqyaQ",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "王天辰",
                    "stu_id": 8149779,
                    "external_user_id": "wmHqcDDgAAKWrKpzwCsc-THY68aRr7zw",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "马弋坤",
                    "stu_id": 20858355,
                    "external_user_id": "wmHqcDDgAA1Dy74jKY9kI_UkKeJMb11A",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "任海睿",
                    "stu_id": 42165589,
                    "external_user_id": "wmHqcDDgAAk7JmaXi3X7lsh4C7wukmXA",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "王三水",
                    "stu_id": 48482893,
                    "external_user_id": "wmHqcDDgAA_kTt7mei8uTO9emTruZRHg",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "程梓然",
                    "stu_id": 61000588,
                    "external_user_id": "wmHqcDDgAA2gIJWoIrmJb-2WbcOknbxA",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "陈腾逸",
                    "stu_id": 53980302,
                    "external_user_id": "wmHqcDDgAAeCdfygGinInmPLIDBGQPDg",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "谷婉琳",
                    "stu_id": 55898890,
                    "external_user_id": "wmHqcDDgAA4qsnxcXcjlCVmFwmpGGJqQ",
                    "identity": 2,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "庞祎恒",
                    "stu_id": 53956510,
                    "external_user_id": "wmHqcDDgAAmsgxycuf29-ztCMvehCJJw",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "李俊燊",
                    "stu_id": 13921156,
                    "external_user_id": "wmHqcDDgAANrQomESIJSHS8oIe5cAnog",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "黄歆桐",
                    "stu_id": 71903346,
                    "external_user_id": "wmHqcDDgAANOFRvxlzRRmEI_tm8v7N0w",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "嘉翊",
                    "stu_id": 65960142,
                    "external_user_id": "wmHqcDDgAAgr_xmVRhatvnsW__jmUR3A",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "杨岳霖",
                    "stu_id": 16570171,
                    "external_user_id": "wmHqcDDgAABHNLpLuXO56C-jI_CLq-XA",
                    "identity": 8,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "王开元",
                    "stu_id": 49943235,
                    "external_user_id": "wmHqcDDgAAKtS1ilIZ6NkTDuhmBTPRiQ",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "李新铎",
                    "stu_id": 47856930,
                    "external_user_id": "wmHqcDDgAAshz06SL25udL4n-PIb5lwQ",
                    "identity": 8,
                    "class_id": 10778300,
                    "plan_id": null
                },
                {
                    "stu_name": "夏子寒",
                    "stu_id": 75087292,
                    "external_user_id": "wmHqcDDgAA3P4LEHrlPN6Nzn9JqOvA3Q",
                    "identity": 3,
                    "class_id": 10778300,
                    "plan_id": null
                }
            ],
            "content_list": [
                {
                    "content_type": 1,
                    "content": "在2022年暑期编程前六讲课中，积极参与，勇于面对新学期挑战，在班级中做出了优秀表率，奖励小朋友奖状一份"
                },
                {
                    "content_type": 2,
                    "content": "https://static.xueersi.com/comfile/1660966965000testimonials.png"
                }
            ]
        }
    ],
    "plan_id": null,
    "course_id": 382474,
    "class_id": "10778299,10778300,10778301"
}`

type QiweiParam struct {
	WxUserId  string         `json:"wx_user_id"`
	WorkCode  string         `json:"work_code"`
	Messages  []QiweiMessage `json:"messages"`
	PlanId    int            `json:"plan_id"`
	ClassId   string         `json:"class_id"`
	SubjectId int            `json:"subject_id"`
	CourseId  int            `json:"course_id"`
}

type QiweiMessage struct {
	StuList     []QiweiStuList     `json:"stu_list"`
	ContentList []QiweiContentList `json:"content_list"`
}

type QiweiStuList struct {
	StuId          int    `json:"stu_id"`
	ExternalUserId string `json:"external_user_id"`
	Identity       int    `json:"identity"`
	StuName        string `json:"stu_name"`
	ClassId        int    `json:"class_id"`
	StuCouId       int    `json:"stu_cou_id"`
	PlanId         int    `json:"plan_id"`
	CourseId       int    `json:"course_id"`
}

type QiweiContentList struct {
	ContentType int    `json:"content_type"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	ImageUrl    string `json:"image_url"`
	Desc        string `json:"desc"`
	Title       string `json:"title"`
}

func main() {
	var p = &QiweiParam{}
	err := json.Unmarshal([]byte(msg), p)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(p.WorkCode)
}
