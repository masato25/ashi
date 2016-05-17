package consulAshi

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/hashicorp/consul/api"
)

func TestConsulAshi(t *testing.T) {
	Convey("Get consul client", t, func() {
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
    })
	})
}
