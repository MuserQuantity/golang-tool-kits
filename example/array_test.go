package example

import (
	"errors"
	"github.com/MuserQuantity/golang-utils/utils"
	"testing"
)

var (
	ErrResult = errors.New("wrong result")
)

func TestItemsContainsItem(t *testing.T) {
	var arrayList1 []interface{}
	arrayList1 = append(arrayList1, "aaa")
	arrayList1 = append(arrayList1, "bbb")
	arrayList1 = append(arrayList1, "ccc")
	arrayStr := utils.InterfaceArray2StringArray(arrayList1)
	if !utils.ContainsString(arrayStr, "aaa") {
		t.Fatal(ErrResult)
	}
	if utils.ContainsString(arrayStr, "ddd") {
		t.Fatal(ErrResult)
	}

	var arrayList2 []interface{}
	arrayList2 = append(arrayList2, uint(111))
	arrayList2 = append(arrayList2, uint(222))
	arrayList2 = append(arrayList2, uint(333))
	arrayUint, ok := utils.InterfaceArray2UintArray(arrayList2)
	if !ok {
		t.Fatal(ErrResult)
	}
	if !utils.ContainsUint(arrayUint, 111) {
		t.Fatal(ErrResult)
	}
	if utils.ContainsUint(arrayUint, 444) {
		t.Fatal(ErrResult)
	}
}
