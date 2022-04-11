package persistence

import (
	"log"
	"bolt"
)

type BoltDb struct {
	Db bolt.DB
}

func (bdb *BoltDb) DbOpen(fileName string) {
	bdb.Db, err := bolt.Open(fileName, 0600, nil)

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
	bucket, err := tx.CreateBucketIfNotExists(bucketName)
	if err != nil {
		log.Fatal(err)
	}
}

func (bdb *BoltDb) DbPut(bucketName string, key string, value string) {
	err = bdb.Db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.Bucket([]byte(bucketName))
		if err != nil {
			return err
		}
		bucket.Put([]byte(key), []byte(value))
		})
	if err != nil {
		log.Fatal(err)
	}
}

func (bdb *BoltDb) DbGet(bucketName string, key string) string {
	err = bdb.Db.View(func(tx *bolt.Tx) error {
		data, err := tx.Bucket([]byte(bucketName)).Get([]byte(key))
		if err != nil {
			log.Fatal(err)
		}	
		return data
	})
}

func (bdb *BoltDb) DbGetAll(bucketName string) []string {
	data := []

	err = bdb.Db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.Bucket([]byte(bucketName));
		if err != nil {
			log.Fatal(err)
		}	

		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			data = append(data, v)
		}
		return data
	})
}

func (bdb *BoltDb) DbDelete(bucketName string, key string) {
	err = bdb.Db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.Bucket([]byte(bucketName))
		if err != nil {
			return err
		}
		bucket.Delete([]byte(key))
		})
	if err != nil {
		log.Fatal(err)
	}
}