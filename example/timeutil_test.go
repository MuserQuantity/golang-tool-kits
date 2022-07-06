package example

import (
	"fmt"
	"github.com/MuserQuantity/golang-utils/utils/timeutil"
	"testing"
)

func TestTimeUtil(t *testing.T) {
	timestamp, err := timeutil.ParseTimeStrInLocation(timeutil.PureNumLayout, "20220706142622", timeutil.ShanghaiTimeZone)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(timestamp.Unix())
}
