/*
@Time : 2018/12/18 9:13 AM
@Author : Tenlu
@File : actiondb
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

/*
CREATE TABLE `mid_autumn_winners` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` int(11) NOT NULL,
	`name` varchar(32) NOT NULL DEFAULT '',
	`phone` varchar(32) NOT NULL DEFAULT '',
	`province` varchar(32) NOT NULL DEFAULT '',
	`city` varchar(32) NOT NULL DEFAULT '',
	`area` varchar(32) NOT NULL DEFAULT '',
	`addinfo` text NOT NULL,
	`prize_level` int(11) NOT NULL COMMENT '0:谢谢参与 1:一等奖 2:二等奖 3:三等奖',
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	PRIMARY KEY (`id`),
	KEY `user_id_idx` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
*/

type Winners struct {
	ID         uint `gorm:"primary_key"`
	UserId     int
	Name       string
	Phone      int
	Province   string
	City       string
	Area       string
	AddInfo    string `gorm:"column:addinfo"` // 重新设置列名为`addinfo`,不打tag会根据大小写被解析为add_info
	PrizeLevel int
	CreatedAt  time.Time
}

// 设置User的表名为`profiles`
func (Winners) TableName() string {
	return "mid_autumn_winners"
}
func main() {
	db := connect()
	//创建数据
	//createWinner := Winners{UserId: 100861, Name: "jack", Phone: 15010549088, Province: "河南省", City: "南阳市", Area: "龙阳县", AddInfo: "龙阳县聚龙潭街108单元100861室", PrizeLevel: 1}
	createWinner := Winners{UserId: 100861, Name: "jack", Phone: 15010549088, Province: "河南省", City: "南阳市", Area: "龙阳县", AddInfo: "龙阳县聚龙潭街108单元100861室", PrizeLevel: 1}

	db.NewRecord(createWinner) // => 返回 `true` ，因为主键为空

	db.Create(&createWinner)

	db.NewRecord(createWinner) // => 在 `user` 之后创建返回 `false`

	// 查询
	var winner Winners
	db.First(&winner)
	fmt.Printf("id:%d,userId:%d,name:%s,phone:%d\n", winner.ID, winner.UserId, winner.Name, winner.Phone)

	var myWin Winners
	db.First(&myWin, "id = ?", 16)
	fmt.Printf("id:%d,userId:%d,name:%s,phone:%d\n", myWin.ID, myWin.UserId, myWin.Name, myWin.Phone)

	var findWinner Winners
	db.Where("name = ?", "jack").Find(&findWinner)
	// SELECT * FROM `mid_autumn_winners`  WHERE (name = 'jack')
	fmt.Printf("id:%d,userId:%d,name:%s,phone:%d\n", findWinner.ID, findWinner.UserId, findWinner.Name, findWinner.Phone)

	var rawWinner Winners
	db.Raw("SELECT id,user_id,name,phone FROM mid_autumn_winners WHERE id = ?", 10).Scan(&rawWinner)
	fmt.Printf("id:%d,userId:%d,name:%s,phone:%d\n", rawWinner.ID, rawWinner.UserId, rawWinner.Name, rawWinner.Phone)

	var updateWinner Winners
	//更新记录
	updateErr := db.Model(&updateWinner).Where("id=?", 24).Update("name1", "toms").Error
	if updateErr != nil {
		panic(updateErr)
	}

	// 删除记录
	var deleteWinner Winners
	deleteErr := db.Where("id=?", 25).Delete(&deleteWinner).Error
	if deleteErr != nil {
		panic(deleteErr)
	}

	defer db.Close()
}

func connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:lookfor@Dny8$@/activity?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Printf("mysql Gorm Connect Success\n")
	return db
}

//相关文档教程:
//1.http://gorm.book.jasperxu.com/crud.html#c
//2.http://gorm.io/zh_CN/docs/method_chaining.html
