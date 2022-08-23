package main

import (
	"go-notes/sourcecode/boltdb/bolt"
	"log"
)

const (
	BlockBucket  = "MyBlocks"
	BlockChainDb = "blockChain.db"
)

func main() {
	db, err := bolt.Open("../go-notes/sourcecode/blotdb/"+BlockChainDb, 0600, bolt.DefaultOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//2.创建表
	err = db.Update(func(tx *bolt.Tx) error {
		//判断要创建的表是否存在
		b := tx.Bucket([]byte(BlockBucket))
		if b == nil {
			//创建叫"MyBucket"的表
			_, err := tx.CreateBucket([]byte(BlockBucket))
			if err != nil {
				//也可以在这里对表做插入操作
				log.Fatal(err)
			}
		}
		//一定要返回nil
		return nil
	})

	////更新数据库失败
	//if err != nil {
	//	log.Fatal(err)
	//}
	////3.更新表数据
	//err = db.Update(func(tx *bolt.Tx) error {
	//
	//	//取出叫"MyBucket"的表
	//	b := tx.Bucket([]byte(BlockBucket))
	//
	//	//往表里面存储数据
	//	if b != nil {
	//		//插入的键值对数据类型必须是字节数组
	//		err := b.Put([]byte("l"), []byte("0x0000"))
	//		err = b.Put([]byte("ll"), []byte("0x0001"))
	//		err = b.Put([]byte("lll"), []byte("0x0002"))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//
	//	//一定要返回nil
	//	return nil
	//})
	//
	////更新数据库失败
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////4.查看表数据
	//err = db.View(func(tx *bolt.Tx) error {
	//
	//	//取出叫"MyBucket"的表
	//	b := tx.Bucket([]byte(BlockBucket))
	//
	//	//往表里面存储数据
	//	if b != nil {
	//
	//		data := b.Get([]byte("l"))
	//		fmt.Printf("%s\n", string(data))
	//		data = b.Get([]byte("l"))
	//		fmt.Printf("%s\n", string(data))
	//	}
	//
	//	//一定要返回nil
	//	return nil
	//})
	//
	////查询数据库失败
	//if err != nil {
	//	log.Fatal(err)
	//}
}
