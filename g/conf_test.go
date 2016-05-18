package g

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConfReader(t *testing.T) {
	Convey("Get dokcer info", t, func() {
    Set("ashi_example", "../conf")
		conf := Config()
		So(len(conf.ADDRESSES), ShouldBeGreaterThan, 0)
  })
}
