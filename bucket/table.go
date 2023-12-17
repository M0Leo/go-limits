package bucket 

import (
	"fmt"
)

type Table struct {
	Buckets map[string]*Bucket
}

func NewTable() *Table {
	return &Table{
		Buckets: make(map[string]*Bucket),
	}
}

func (t *Table) GetBucket(ip string) *Bucket {
	if bucket, ok := t.Buckets[ip]; ok {
		return bucket
	}
	newBucket := NewBucket()
	t.Buckets[ip] = newBucket
	return newBucket
}


func (t *Table) HandleRequest(ip string) bool {
	bucket := t.GetBucket(ip)

	if bucket.HasToken() {
		fmt.Printf("Request from IP %s handled. Token: %s\n", ip, bucket.Tokens[len(bucket.Tokens)-1])
		bucket.Pop()
		return true
	}

	fmt.Printf("Request from IP %s declined. Bucket is empty.\n", ip)
	return false
}
