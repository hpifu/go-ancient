package mysql

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMysqlDB_SelectAncient(t *testing.T) {
	m, err := NewMysqlDB("hatlonely:keaiduo1@tcp(test-mysql:3306)/hads?charset=utf8&parseTime=True&loc=Local")
	Convey("test ancient", t, func() {
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		ancient := &Ancient{
			Title:   "test 静夜思",
			Author:  "李白",
			Dynasty: "唐",
			Content: "床前明月光,疑是地上霜。举头望明月,低头思故乡",
		}
		err = m.db.Delete(&Ancient{
			Author:  ancient.Author,
			Dynasty: ancient.Dynasty,
		}).Error
		So(err, ShouldBeNil)
		for i := 0; i < 20; i++ {
			a := &Ancient{
				Title:   fmt.Sprintf("%s-%v", ancient.Title, i),
				Author:  ancient.Author,
				Dynasty: ancient.Dynasty,
				Content: ancient.Content,
			}
			err = m.db.Create(a).Error
			So(err, ShouldBeNil)
		}

		Convey("select ancient", func() {
			as, err := m.SelectAncientByTitleAndAuthor(&Ancient{
				Author:  ancient.Author,
				Dynasty: ancient.Dynasty,
			}, 0, 10)
			So(err, ShouldBeNil)
			So(len(as), ShouldEqual, 10)
			for i := 0; i < 10; i++ {
				So(as[i].Title, ShouldEqual, fmt.Sprintf("%s-%v", ancient.Title, i))
				So(as[i].Author, ShouldEqual, ancient.Author)
				So(as[i].Dynasty, ShouldEqual, ancient.Dynasty)
				So(as[i].Content, ShouldEqual, ancient.Content)
			}
		})
	})
}
