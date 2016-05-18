package consulAshi

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/masato25/ashi/g"
)

func TestConsulAshi(t *testing.T) {
	g.Set("ashi_example", "../conf")
	conf := g.Config()
	Convey("Consul register testing", t, func() {
    clinet := Client()
		Convey("Gen Consul services register text", func(){
			apiw := ParepareReg(conf.IP, "nignig", "owl", 4567)
			So(apiw.Address, ShouldEqual, conf.IP)
		})
		Convey("Register Consul services", func(){
			err := ServiceRegister(conf.IP, "nignig", "owl", 4567, clinet)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Check Consul Server status", t, func(){
		Convey("get return code 200", func(){
			scode, err := CheckConsulServer(conf.IP, 8500)
			if err != nil{
				So(err, ShouldNotBeEmpty)
			}else{
				So(scode.StatusCode, ShouldEqual, 200)
			}
		})
		Convey("get return code nil", func(){
			_, err := CheckConsulServer("192.0.0.1", 8500)
			if err != nil{
				So(err.Error(), ShouldContainSubstring, "Timeout")
			}
		})
	})
}
