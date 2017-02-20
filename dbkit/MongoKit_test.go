package dbkit

import (
	. "github.com/smartystreets/goconvey/convey"
	"git.gumpcome.com/gokit/logkit"
	"testing"
)

func TestMongoDB(t *testing.T) {
	logkit.InitLog()
	m := (&MongoDB{BeeLogger: logkit.GetLog()}).New("yp", "gumpcometest", "42.159.235.113:28010", "admin", "test").GetMongo()
	err := m.C("demo").Insert(struct {
		Id   int
		Name string
	}{
		1,
		"yp",
	})
	Convey("TestMongoDB", t, func() {
		So(err, ShouldBeNil)
	})
}
