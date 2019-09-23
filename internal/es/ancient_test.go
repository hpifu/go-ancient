package es

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestES_SearchAncient(t *testing.T) {
	e, err := NewES("http://127.0.0.1:9200")
	Convey("test ancient", t, func() {
		So(err, ShouldBeNil)
		as, err := e.SearchAncient("李白 行路难", 0, 10)
		So(err, ShouldBeNil)
		for _, ancient := range as {
			fmt.Println(ancient)
		}
	})
}

func TestSplit(t *testing.T) {
	Convey("test split", t, func() {
		So(split("你好，世界。你好,golang"), ShouldResemble, []string{
			"你好", "世界", "你好", "golang",
		})
	})
}
