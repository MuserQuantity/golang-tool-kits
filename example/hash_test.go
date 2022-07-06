package example

import (
	"github.com/MuserQuantity/golang-utils/utils"
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	password := "123456"
	passHash := utils.BcryptHash(password)
	log.Println(passHash)
	ok := utils.BcryptCheck(password, passHash)
	if !ok {
		t.Fatal("校验失败")
	}
}
