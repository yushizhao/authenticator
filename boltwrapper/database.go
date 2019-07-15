package boltwrapper

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type AUTHDB struct {
	boltptr *bolt.DB
}

var UserDB AUTHDB

// multiple processes cannot open the same database at the same time
func InitDB() {
	// It will be created if it doesn't exist.
	db, err := bolt.Open("user.bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("USER"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	UserDB.boltptr = db
}

// If the key exists then it will return its byte slice value. If it doesn't exist then it will return nil.
func (this *AUTHDB) GetUser(name string) (v []byte) {
	this.boltptr.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("USER"))
		v = b.Get([]byte(name))
		return nil
	})
	return v
}

func (this *AUTHDB) SetUser(name string, v []byte) error {
	return this.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("USER"))
		err := b.Put([]byte(name), v)
		return err
	})
}

func (this *AUTHDB) ListUser() map[string][]byte {
	tmp := make(map[string][]byte)
	this.boltptr.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("USER"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tmp[string(k)] = v
		}
		return nil
	})
	return tmp
}

func (this *AUTHDB) DelUser(name string, v []byte) error {
	return this.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("USER"))
		err := b.Delete([]byte(name))
		return err
	})
}
