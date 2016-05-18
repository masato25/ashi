package dockerReader

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/masato25/ashi/g"
)

func TestDockerReader(t *testing.T) {
	g.Set("ashi", "../conf")
	conf := g.Config()
	Convey("Get dokcer info", t, func() {
		container, _ := GetContainerList(conf.DOCKERSOC)
		Convey("Get ip addresss of container", func(){
			ip := getDokcerIpv4(container[0].Networks.Networks)
			So(ip, ShouldNotBeEmpty)
		})
		Convey("Get number of docker", func(){
      So(len(container), ShouldBeGreaterThan, 0)
    })

		Convey("Get name of docker", func(){
			So(container[0].ID, ShouldNotBeEmpty)
		})

		Convey("Get ports", func(){
			cobj := GetContainers(conf.DOCKERSOC)
			So(len(cobj[0].Ports), ShouldBeGreaterThan, 0)
		})
	})
}
