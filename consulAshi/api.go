package consulAshi

import (
  "github.com/hashicorp/consul/api"
  "github.com/masato25/ashi/g"
  "fmt"
  "net/http"
  "time"
)

func Client() (*api.Client) {
  conf := g.Config()
  apiconf := api.DefaultConfig()
  address := selectAvailableConsulServer(conf.ADDRESSES, conf.HTTPPORT)
  apiconf.Address = fmt.Sprintf("%s:%d", address, conf.HTTPPORT)
  apiconf.Datacenter = conf.DATACENTER
  client, _ := api.NewClient(apiconf)
  return client
}

func ParepareReg(ip string, name string, id string, port int) (reg *api.CatalogRegistration){
  conf := g.Config()
  services := &api.AgentService{
    ID: fmt.Sprintf("%s-%s", name, id),
    Service: name,
    Tags: []string{name},
    Port:port,
    Address: ip,
  }
  // check := &api.AgentServiceCheck{
  //     TCP: fmt.Sprintf("%s:%d", ip, port),
	// }
  reg = &api.CatalogRegistration{
    Datacenter: conf.DATACENTER,
    Address:    selectAvailableConsulServer(conf.ADDRESSES, conf.HTTPPORT),
    Node:       conf.NODE,
    Service:    services,
    // Check:      check,
  }
  return
}

// func ParepareReg2()(reg *api.CatalogRegistration){
//   service := &api.AgentService{
// 		ID:      "redis1",
// 		Service: "redis",
// 		Tags:    []string{"master", "v1"},
// 		Port:    8000,
// 	}
//
// 	check := &api.AgentCheck{
// 		Node:      "foobar",
// 		CheckID:   "service:redis1",
// 		Name:      "Redis health check",
// 		Notes:     "Script based health check",
// 		ServiceID: "redis1",
// 	}
//
// 	reg = &api.CatalogRegistration{
// 		Datacenter: "owl",
// 		Node:       "myarch",
// 		Address:    "10.0.0.165",
// 		Service:    service,
// 		Check:      check,
// 	}
//   return
// }


func ParepareRegSer(ip string, name string, id string, port int) (reg *api.AgentServiceRegistration){
  conf := g.Config()
  check := &api.AgentServiceCheck{
      TCP: fmt.Sprintf("%s:%d", ip, port),
      Status: api.HealthPassing,
      Interval: "20s",
      Timeout: "3s",
	}
  reg = &api.AgentServiceRegistration{
    Address:  selectAvailableConsulServer(conf.ADDRESSES, conf.HTTPPORT),
    Check: check,
    ID:fmt.Sprintf("%s-%s", name, id),
    Name: name,
    Tags: []string{name},
    Port: port,
    EnableTagOverride: false,
  }
  return
}

func ServiceRegister(ip string, name string, id string, port int, client *api.Client) (err error){
  reg := ParepareRegSer(ip, name, id, port)
  // wmeta, _ := client.Catalog().Register(reg, nil)

  err = client.Agent().ServiceRegister(reg)
  return err
}

func QueryServies(servicesName string, client *api.Client) (services []*api.CatalogService){
  catalog := client.Catalog()
  services, _, _ = catalog.Service("consul", "", nil)
  return
}


func selectAvailableConsulServer(consult []string, port int) (string){
  address := "127.0.0.1"
  for _, addr := range consult{
    _, err := CheckConsulServer(addr, port)
    if err == nil{
      address = addr
      return address
    }
  }
  return address
}

func CheckConsulServer(host string, port int) (*http.Response, error){
  timeout := time.Duration(2 * time.Second)
  client := http.Client{
    Timeout: timeout,
  }
  url := fmt.Sprintf("http://%s:%d", host, port)
  resp, err := client.Get(url)
  return resp, err
}
