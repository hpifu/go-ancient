package mysql

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMysqlDB_SelectAncient(t *testing.T) {
	m, err := NewMysql("hatlonely:keaiduo1@tcp(test-mysql:3306)/hads?charset=utf8&parseTime=True&loc=Local")
	Convey("test ancient", t, func() {
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)

		ancient := &Ancient{
			Title:   "test 静夜思",
			Author:  "李白",
			Dynasty: "唐",
			Content: "床前明月光,疑是地上霜。举头望明月,低头思故乡",
		}
		for i := 0; i < 20; i++ {
			So(m.db.Delete(&Ancient{ID: i + 1}).Error, ShouldBeNil)
			So(m.db.Create(&Ancient{
				ID:      i + 1,
				Title:   fmt.Sprintf("%s-%v", ancient.Title, i+1),
				Author:  ancient.Author,
				Dynasty: ancient.Dynasty,
				Content: ancient.Content,
			}).Error, ShouldBeNil)
		}

		Convey("select ancients", func() {
			{
				as, err := m.SelectAncients(0, 10)
				So(err, ShouldBeNil)
				So(len(as), ShouldEqual, 10)
				for i := 0; i < 10; i++ {
					So(as[i].ID, ShouldEqual, i+1)
					So(as[i].Title, ShouldEqual, fmt.Sprintf("%s-%v", ancient.Title, i+1))
					So(as[i].Author, ShouldEqual, ancient.Author)
					So(as[i].Dynasty, ShouldEqual, ancient.Dynasty)
					So(as[i].Content, ShouldEqual, "")
				}
			}
			{
				as, err := m.SelectAncients(10, 20)
				So(err, ShouldBeNil)
				So(len(as), ShouldEqual, 10)
				for i := 0; i < 10; i++ {
					So(as[i].ID, ShouldEqual, i+11)
					So(as[i].Title, ShouldEqual, fmt.Sprintf("%s-%v", ancient.Title, i+11))
					So(as[i].Author, ShouldEqual, ancient.Author)
					So(as[i].Dynasty, ShouldEqual, ancient.Dynasty)
					So(as[i].Content, ShouldEqual, "")
				}
			}
			{
				as, err := m.SelectAncients(20, 10)
				So(err, ShouldBeNil)
				So(len(as), ShouldEqual, 0)
			}
		})

		Convey("select ancient by id", func() {
			for i := 0; i < 20; i++ {
				a, err := m.SelectAncientByID(i + 1)
				So(err, ShouldBeNil)
				So(a.ID, ShouldEqual, i+1)
				So(a.Title, ShouldEqual, fmt.Sprintf("%s-%v", ancient.Title, i+1))
				So(a.Author, ShouldEqual, ancient.Author)
				So(a.Dynasty, ShouldEqual, ancient.Dynasty)
				So(a.Content, ShouldEqual, ancient.Content)
			}

			a, err := m.SelectAncientByID(21)
			So(err, ShouldBeNil)
			So(a, ShouldBeNil)
		})

		Convey("select authors", func() {
			{
				authors, err := m.SelectAuthors(0, 10)
				So(err, ShouldBeNil)
				So(authors, ShouldResemble, []string{"李白"})
			}
			{
				authors, err := m.SelectAuthors(10, 10)
				So(err, ShouldBeNil)
				So(authors, ShouldBeNil)
			}
		})

		Convey("select dynastys", func() {
			{
				dynastys, err := m.SelectDynastys(0, 10)
				So(err, ShouldBeNil)
				So(dynastys, ShouldResemble, []string{"唐"})
			}
			{
				dynastys, err := m.SelectDynastys(10, 10)
				So(err, ShouldBeNil)
				So(dynastys, ShouldBeNil)
			}
		})

		Convey("select author", func() {
			{
				as, err := m.SelectAncientByAuthor("李白", 0, 10)
				So(err, ShouldBeNil)
				for i := 0; i < 10; i++ {
					So(as[i].ID, ShouldEqual, i+1)
					So(as[i].Title, ShouldEqual, fmt.Sprintf("%s-%v", ancient.Title, i+1))
					So(as[i].Author, ShouldEqual, ancient.Author)
					So(as[i].Dynasty, ShouldEqual, ancient.Dynasty)
					So(as[i].Content, ShouldEqual, "")
				}
			}
			{
				as, err := m.SelectAncientByAuthor("李白", 20, 10)
				So(err, ShouldBeNil)
				So(as, ShouldBeNil)
			}
		})

		Convey("select dynasty", func() {
			{
				ds, err := m.SelectAncientByDynasty("唐", 0, 10)
				So(err, ShouldBeNil)
				for i := 0; i < 10; i++ {
					So(ds[i].ID, ShouldEqual, i+1)
					So(ds[i].Title, ShouldEqual, fmt.Sprintf("%s-%v", ancient.Title, i+1))
					So(ds[i].Author, ShouldEqual, ancient.Author)
					So(ds[i].Dynasty, ShouldEqual, ancient.Dynasty)
					So(ds[i].Content, ShouldEqual, "")
				}
			}
			{
				ds, err := m.SelectAncientByDynasty("唐", 20, 10)
				So(err, ShouldBeNil)
				So(ds, ShouldBeNil)
			}
		})
	})
}
