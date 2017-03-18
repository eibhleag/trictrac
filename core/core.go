package core

import (
	"encoding/binary"
	"log"
	"path"

	"os"

	"fmt"

	"container/list"

	"bytes"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
)

type Collection struct {
	db *bolt.DB
}

type Pair struct {
	Key   string
	Value uint64
}

// Everything here needs to handle possible failure - it currently doesn't

func OpenCollection() *Collection {
	base, pathError := homedir.Dir()
	if pathError != nil {
		base, _ = os.Getwd()
	}
	dbFolder := path.Join(base, ".trictrac")
	os.MkdirAll(dbFolder, 0700)
	dbPath := path.Join(dbFolder, "default.tt")
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("main"))
		if err != nil {
			return fmt.Errorf("couldn't create main bucket: %s", err)
		}
		return nil
	})

	return &Collection{db}
}

func (c *Collection) Close() {
	c.db.Close()
}

func requireKey(b *bolt.Bucket, key []byte) []byte {
	value := b.Get(key)
	if value == nil {
		fmt.Printf("'%s': not found\n", key)
		os.Exit(-1)
	}
	return value
}

func (c *Collection) List() *list.List {
	l := list.New()
	c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			l.PushBack(Pair{string(k), btoi(v)})
		}

		return nil
	})
	return l
}

func (c *Collection) SumPrefix(prefix string) uint64 {
	keyPrefix := []byte(prefix + ".")
	var total uint64
	c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		cursor := b.Cursor()

		for k, v := cursor.Seek(keyPrefix); k != nil && bytes.HasPrefix(k, keyPrefix); k, v = cursor.Next() {
			total += btoi(v)
		}

		return nil
	})
	return total
}

func (c *Collection) ListPrefix(prefix string) *list.List {
	// avoid partially matching e.g. 'abc' with 'a'
	// vs intended behaviour of matching 'a.b.c'
	keyPrefix := []byte(prefix + ".")
	l := list.New()
	c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		cursor := b.Cursor()

		for k, v := cursor.Seek(keyPrefix); k != nil && bytes.HasPrefix(k, keyPrefix); k, v = cursor.Next() {
			l.PushBack(Pair{string(k), btoi(v)})
		}

		return nil
	})
	return l
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func (c *Collection) New(key string, value uint64) uint64 {
	c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		byteKey := []byte(key)
		exists := b.Get(byteKey)
		if exists != nil {
			fmt.Printf("'%s': already exists\n", key)
			os.Exit(-1)
		}
		return b.Put(byteKey, itob(value))
	})
	return value
}

func (c *Collection) Del(key string) {
	c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		byteKey := []byte(key)
		requireKey(b, byteKey)
		return b.Delete(byteKey)
	})
}

func (c *Collection) Set(key string, value uint64) uint64 {
	c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		byteKey := []byte(key)
		requireKey(b, byteKey)
		return b.Put(byteKey, itob(value))
	})
	return value
}

func (c *Collection) Get(key string) uint64 {
	result := make([]byte, 8)
	c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("main"))
		value := requireKey(b, []byte(key))
		copy(result, value)
		return nil
	})
	return btoi(result)
}

func (c *Collection) Inc(key string, value uint64) uint64 {
	old := c.Get(key)
	return c.Set(key, old+value)
}

func (c *Collection) Dec(key string, value uint64) uint64 {
	old := c.Get(key)
	if old < value {
		return c.Set(key, 0)
	}
	return c.Set(key, old-value)
}
