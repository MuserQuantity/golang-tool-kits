package utils

import (
	"github.com/lithammer/shortuuid/v4"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// NewRandStringRunes 从表里字符生成随机串
func NewRandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// NewUUID 从表里字符生成UUID
func NewUUID() string {
	alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxy-"
	uuid := shortuuid.NewWithAlphabet(alphabet)
	if uuid[0] == '-' || uuid[len(uuid)-1] == '-' {
		return NewUUID()
	}
	return uuid
}
