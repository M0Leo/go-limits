package bucket

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"
)

type Bucket struct {
	Capacity int
	Tokens   []string
	Mu       sync.Mutex
}

func NewBucket() *Bucket {
	b := &Bucket{
		Tokens:   []string{},
		Capacity: 10,
	}

	for i := 0; i < 5; i++ {
		b.Push()
	}

	go b.startTokenAddition()
	return b
}

func (b *Bucket) Pop() {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	if len(b.Tokens) > 0 {
		b.Tokens = b.Tokens[:len(b.Tokens)-1]
	}
}

func (b *Bucket) startTokenAddition() {
    ticker := time.NewTicker(5 * time.Second)

    for range ticker.C {
        if len(b.Tokens) < b.Capacity {
            b.Push()
        }
    }
}

func (b *Bucket) Push() {
	token, err := generateToken()
	if err != nil {
		logError(err.Error())
	}
	b.Tokens = append(b.Tokens, token)
}

func (b *Bucket) HasToken() bool {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	return len(b.Tokens) > 0
}

func generateToken() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(randomBytes)
	return token, nil
}

func logError(msg string) error {
	log.Println(msg)
	return fmt.Errorf(msg)
}
