package consulAshi

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/hashicorp/consul/api"
	"github.com/masato25/ashi/g"
)

func TestConsulAshi(t *testing.T) {
	g.Set("ashi_example", "../conf")
	Convey("Get consul client kv", t, func() {
    clinet := Client()
		Convey("Put KV", func(){
      // Get a handle to the KV API
      kv := clinet.KV()
      // PUT a new KV pair
      p := &api.KVPair{Key: "foo", Value: []byte("test")}
      _, err := kv.Put(p, nil)
      if err != nil {
          panic(err)
      }
      // Lookup the pair
      pair, _, err := kv.Get("foo", nil)
      if err != nil {
          panic(err)
      }
      So(string(pair.Value), ShouldEqual, "test")
			kv.Delete("foo", nil)
    })
		Convey("Gen Consul services register text", func(){
			apiw := ParepareReg("10.0.0.165", "nignig", "owl", 4567)
			So(apiw.Address, ShouldEqual, "10.0.0.165")
		})
		Convey("Register Consul services", func(){
			err := ServiceRegister("10.0.0.165", "nignig", "owl", 4567, clinet)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Check Consul Server status", t, func(){
		Convey("get return code 200", func(){
			scode, err := CheckConsulServer("owl.fastweb.com.cn", 80)
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
