package example

import (
	"github.com/MuserQuantity/golang-utils/utils"
	"log"
	"testing"
)

func TestMD5(t *testing.T) {
	password := []byte("123456")
	hash := utils.MD5(password)
	log.Println(hash)
}
