package main

import (
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("metadata.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	if err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		if err = b.Put([]byte("dean"), []byte("dut")); err != nil {
			return fmt.Errorf("create data: %v", err)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		school := b.Get([]byte("dean"))
		fmt.Println(string(school))
		return nil
	}); err != nil {
		panic(err)
	}
}
