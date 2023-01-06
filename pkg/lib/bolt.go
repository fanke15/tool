package lib

import (
	"github.com/boltdb/bolt"
	"github.com/fanke15/tool/pkg/basic"
	"sync"
	"time"
)

const (
	DefaultBucketName NameTypeOfBolt = "default"
	DefaultDBName     NameTypeOfBolt = "tool"
	DefaultDBDir                     = "./web/static/db/"
)

var (
	dbMap sync.Map
)

type (
	NameTypeOfBolt string
	DBer           struct {
		bucketName NameTypeOfBolt
		dbName     NameTypeOfBolt
	}
)

func NewBolt(name ...NameTypeOfBolt) {
	var (
		dbName = DefaultDBName
		db     *bolt.DB
		err    error
	)
	if len(name) > basic.Zero {
		dbName = name[basic.Zero]
	}
	// 关闭旧连接
	if dbTemp, _ := dbMap.Load(dbName); dbTemp != nil {
		_ = dbTemp.(*bolt.DB).Close()
	}
	dbMap.Delete(dbName)

	// 重新建立连接
	if err = basic.Retry(func() error {
		db, err = conn(dbName)
		return err
	}, basic.Five, basic.Three*time.Second); err != nil {
		panic(err)
	}
	dbMap.Store(dbName, db)
}

// InitBolt 加载存储桶
func InitBolt(name ...NameTypeOfBolt) *DBer {
	var (
		dn = DefaultDBName
		bn = DefaultBucketName
	)

	if len(name) > basic.Zero {
		dn = name[basic.Zero]
	}
	if len(name) > basic.One {
		bn = name[basic.One]
	}
	return &DBer{bn, dn}
}

// QueryAll 批量全部数据
func (b *DBer) QueryAll() map[string][]byte {
	var (
		db   = getBolt(b.dbName)
		data = make(map[string][]byte)
	)
	_ = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.bucketName))
		return bucket.ForEach(func(k, v []byte) error {
			data[string(k)] = v
			return nil
		})
	})
	return data
}

// Query 加载数据
func (b *DBer) Query(key string) []byte {
	var (
		db   = getBolt(b.dbName)
		data []byte
	)

	_ = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.bucketName))
		data = bucket.Get([]byte(key))
		return nil
	})
	return data
}

// SaveBatch 批量缓存数据
func (b *DBer) SaveBatch(data map[string][]byte) error {
	var (
		db = getBolt(b.dbName)
	)
	_ = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.bucketName))
		if err != nil {
			return err
		}
		for k, v := range data {
			err = bucket.Put([]byte(k), v)
		}
		return err
	})
	return nil
}

// Save 缓存数据
func (b *DBer) Save(key string, val []byte) error {
	var (
		db = getBolt(b.dbName)
	)
	_ = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.bucketName))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(key), val)
	})
	return nil
}

// 更新数据

// 删除数据

// ---------------------------内部私有方法---------------------------//
func conn(name NameTypeOfBolt) (*bolt.DB, error) {
	db, err := bolt.Open(basic.AnySliceToStr(basic.StrNull, DefaultDBDir, string(name), ".db"),
		0600, &bolt.Options{Timeout: basic.One * time.Second})
	return db, err
}

// 获取操作对象
func getBolt(name ...NameTypeOfBolt) *bolt.DB {
	var dbName = DefaultDBName
	if len(name) > basic.Zero {
		dbName = name[basic.Zero]
	}
	if dbTemp, ok := dbMap.Load(dbName); ok {
		return dbTemp.(*bolt.DB)
	}
	return nil
}
