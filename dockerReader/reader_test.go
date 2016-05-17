package dockerReader

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDockerReader(t *testing.T) {
	Convey("Get dokcer info", t, func() {
		container, _ := ConatainerRead()
		Convey("Get number of docker", func(){
      So(len(container), ShouldBeGreaterThan, 0)
    })

		Convey("Get name of docker", func(){
			So(container[0].ID, ShouldNotBeEmpty)
		})

		Convey("Get ports", func(){
			cobj := GetPublicPort(container)
			So(len(cobj[0].Ports), ShouldBeGreaterThan, 0)
		})
	})
}
