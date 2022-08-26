package boltdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"go-notes/sourcecode/boltdb/bolt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

const (
	DbName          = "myFile.db"
	FatherBlockName = "fatherBlocks"
	ChildBlockName  = "childBlocks"
)

func TestBolt(t *testing.T) {
	pwd, err := os.Getwd()
	db, err := bolt.Open(pwd+"/"+DbName, 0600, &bolt.Options{
		ReadOnly: false,
		Timeout:  1 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	newDb, err := bolt.Open(pwd+"/test.db", 0600, bolt.DefaultOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer newDb.Close()

	// 显式读写事务
	tx, err := newDb.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	m, err := tx.CreateBucketIfNotExists([]byte("my_test"))
	if err != nil {
		log.Fatal(err)
	}
	m.Put([]byte("a"), []byte("aaa"))
	m.Put([]byte("b"), []byte("bbb"))

	fmt.Println("b", string(m.Get([]byte("b"))))
	pageInfo, _ := tx.Page(0)
	fmt.Printf("pageInfo:%+v\n", pageInfo)

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	if err = db.Update(func(tx *bolt.Tx) error {
		f := tx.Bucket([]byte(FatherBlockName))
		f.Put([]byte("f2"), []byte("this is f2"))
		fmt.Println("f2:", string(f.Get([]byte("f2"))))
		return nil
	}); err != nil {

	}
	// 置隐式事务Update, View, Batch
	if err = db.Update(func(tx *bolt.Tx) error {
		// 判断要创建的表是否存在
		f := tx.Bucket([]byte(FatherBlockName))
		if f == nil {
			// new bucket
			f, err = tx.CreateBucket([]byte(FatherBlockName))
			if err != nil {
				return err
			}
			f.Put([]byte("f1"), []byte("this is f1"))
			return nil
		}
		f.Put([]byte("f2"), []byte("this is f2"))

		c, err := tx.CreateBucketIfNotExists([]byte(ChildBlockName))
		if err != nil {
			return err
		}
		c.Put([]byte("c1"), []byte("this is c1"))
		for i := 0; i < 5; i++ {
			c.Put([]byte(cast.ToString(i)), []byte("this is"+cast.ToString(i)))
		}
		err = c.Delete([]byte("c1"))
		if err != nil {
			return err
		}
		fmt.Printf("c1:%+v\n", string(c.Get([]byte("c1"))))

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// 只读数据库
	if err = db.View(func(tx *bolt.Tx) error {
		f := tx.Bucket([]byte(FatherBlockName))
		if f != nil {
			f1 := f.Get([]byte("f1"))
			fmt.Printf("f1:%s\n", string(f1))

			f2 := f.Get([]byte("f2"))
			fmt.Printf("f2:%s\n", string(f2))
		}

		fmt.Println("------区间扫描------")
		// Assume our events bucket exists and has RFC3339 encoded time keys.
		c := tx.Bucket([]byte(ChildBlockName))
		if c != nil {
			cur3 := c.Cursor()
			// Our time range spans the 90's decade.
			min := []byte("1")
			max := []byte("4")

			// Iterate over the 90's.
			for k, v := cur3.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = cur3.Next() {
				fmt.Printf("%s: %s\n", k, v)
			}
		}

		fmt.Println("------ 前缀扫描------")
		// Assume bucket exists and has keys
		if c2 := tx.Bucket([]byte(ChildBlockName)); c2 != nil {
			cur := c2.Cursor()

			prefix := []byte("c")
			for k, v := cur.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cur.Next() {
				fmt.Printf("key=%s, value=%s\n", k, v)
			}
		}
		fmt.Println("------根据key遍历-----")
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(FatherBlockName))
		if b != nil {
			cur2 := b.Cursor()
			for k, v := cur2.First(); k != nil; k, v = cur2.Next() {
				fmt.Printf("key=%s, value=%s\n", k, v)
			}
		}
		fmt.Println("--------------")
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// 批量读写数据库：
	if err = db.Batch(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(ChildBlockName))
		if c == nil {
			return errors.New("bucket is nil")
		}
		for i := 0; i < 8; i++ {
			key := strconv.Itoa(i)
			err = c.Put([]byte("c"+key), []byte("this is c"+key))
			if err != nil {
				return err
			}
			if i > 5 {
				c.Delete([]byte("c" + key))
			}
		}

		if err = c.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		}); err != nil {
			log.Fatal(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	go func() {
		// Grab the initial stats.
		prev := db.Stats()
		fmt.Printf("prev status:%+v\n", prev)
		for {
			// Wait for 10s.
			time.Sleep(10 * time.Second)

			// Grab the current stats and diff them.
			stats := db.Stats()
			diff := stats.Sub(&prev)

			// Encode stats to JSON and print to STDERR.
			json.NewEncoder(os.Stderr).Encode(diff)

			// Save stats for the next loop.
			prev = stats
			fmt.Printf("diff status:%+v\n", diff)
		}
	}()
}
