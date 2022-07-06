package example

import (
	"github.com/MuserQuantity/golang-utils/utils"
	"log"
	"testing"
)

func TestGenerator(t *testing.T) {
	log.Println(utils.NewRandStringRunes(10))
	log.Println(utils.NewUUID())
}
