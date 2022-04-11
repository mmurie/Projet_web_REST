package persistence

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type BoltDb struct {
	Db *bolt.DB
}

func (bdb *BoltDb) DbOpen(fileName string) {
	var err error
	bdb.Db, err = bolt.Open(fileName, 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (bdb *BoltDb) DbClose() {
	bdb.Db.Close()
}

func (bdb *BoltDb) DbPath() string {
	return bdb.Db.Path()
}

func (bdb *BoltDb) DbCreateBucket(bucketName string) {
	err := bdb.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (bdb *BoltDb) DbPut(bucketName string, key string, value string) {
	err := bdb.Db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte(bucketName)).Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (bdb *BoltDb) DbGet(bucketName string, key string) string {

	err := bdb.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket \"%q\" not found", bucketName)
		}
		val := bucket.Get([]byte(key))
		fmt.Println(val)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return ""
}

func (bdb *BoltDb) DbGetAll(bucketName string) []string {
	var data []string

	err := bdb.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket \"%q\" not found", bucketName)
		}

		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			data = append(data, string(v))
			fmt.Println(v)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (bdb *BoltDb) DbDelete(bucketName string, key string) {
	err := bdb.Db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket \"%q\" not found", bucketName)
		}
		err := bucket.Delete([]byte(key))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
