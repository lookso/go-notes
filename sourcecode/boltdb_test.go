package sourcecode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-notes/sourcecode/boltdb/bolt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

const (
	DbName          = "myFile.db"
	FatherBlockName = "FatherBlocks"
	ChildBlockName  = "ChildBlocks"
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

	if err := db.Update(func(tx *bolt.Tx) error {
		// 判断要创建的表是否存在
		f := tx.Bucket([]byte(FatherBlockName))
		if f == nil {
			//创建叫"MyBucket"的表,建表 new bucket
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
		f1 := f.Get([]byte("f1"))
		fmt.Printf("f1:%s\n", string(f1))

		f2 := f.Get([]byte("f2"))
		fmt.Printf("f2:%s\n", string(f2))


		// 区间扫描
		// Assume our events bucket exists and has RFC3339 encoded time keys.
		cur3 := tx.Bucket([]byte("Events")).Cursor()

		// Our time range spans the 90's decade.
		min := []byte("1990-01-01T00:00:00Z")
		max := []byte("2000-01-01T00:00:00Z")

		// Iterate over the 90's.
		for k, v := cur3.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = cur3.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}

		// 前缀扫描
		// Assume bucket exists and has keys
		cur := tx.Bucket([]byte("MyBucket")).Cursor()

		prefix := []byte("1234")
		for k, v := cur.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cur.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		//根据key遍历
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("MyBucket"))

		cur2 := b.Cursor()
		for k, v := cur2.First(); k != nil; k, v = cur2.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// 批量读写数据库：
	if err = db.Batch(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(ChildBlockName))
		for i := 0; i < 100; i++ {
			key := strconv.Itoa(i)
			err := c.Put([]byte("c"+key), []byte("this is c"+key))
			if err != nil {
				return err
			}
			if i > 30 {
				c.DeleteBucket([]byte("c" + key))
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
		}
	}()

}
