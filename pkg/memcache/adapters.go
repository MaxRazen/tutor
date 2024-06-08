package memcache

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"cloud.google.com/go/storage"
)

type GCPStorage struct {
	bucket *storage.BucketHandle
	object string
}

func (s *GCPStorage) NewStorableItem(key, val string, expIn int64) StorableItem {
	return &storableItem{
		key:       key,
		value:     val,
		expiresIn: expIn,
	}
}

func (s *GCPStorage) LoadMemcache() ([]StorableItem, error) {
	ctx := context.TODO()
	items := make([]StorableItem, 0)
	or, err := s.bucket.Object(s.object).NewReader(ctx)

	if err != nil {
		return items, err
	}

	csvr := csv.NewReader(or)

	for {
		row, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("memcache/load", err)
			break
		}

		expiresIn, err := strconv.ParseInt(row[2], 10, 64)

		if err != nil {
			log.Println("memcache/load:", err)
			continue
		}

		items = append(items, storableItem{
			key:       row[0],
			value:     row[1],
			expiresIn: expiresIn,
		})
	}

	return items, nil
}

func (s *GCPStorage) DumpMemcache(items []StorableItem) error {
	ctx := context.TODO()
	ow := s.bucket.Object(s.object).NewWriter(ctx)
	ow.ObjectAttrs.ContentType = "text/csv"
	csvw := csv.NewWriter(ow)

	for _, item := range items {
		err := csvw.Write([]string{
			item.Key(),
			item.Value(),
			strconv.FormatInt(item.ExpiresIn(), 10),
		})

		if err != nil {
			log.Println("memcache/dump:", err)
			continue
		}
	}

	csvw.Flush()
	return ow.Close()
}

func NewGCPStorage(bucket *storage.BucketHandle, object string) *GCPStorage {
	return &GCPStorage{
		bucket: bucket,
		object: object,
	}
}
