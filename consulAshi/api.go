package consulAshi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/masato25/ashi/g"
)

func tlsSetup(conf *g.Gconfig) *api.TLSConfig {
	var tlsconf api.TLSConfig
	tlsconf.Address = conf.IP
	certconf := conf.CertConf
	tlsconf.CAFile = certconf.CAFile
	tlsconf.CertFile = certconf.CertFile
	tlsconf.KeyFile = certconf.KeyFile
	tlsconf.InsecureSkipVerify = true
	return &tlsconf
}

func Client() *api.Client {
	conf := g.Config()
	apiconf := api.DefaultConfig()
	var port int
	apiconf.Token = conf.TOKEN
	if conf.CertConf.Enable {
		log.Println("cerconf is enabled")
		var transport http.Transport
		tlssetup := tlsSetup(conf)
		tlsconf, err := api.SetupTLSConfig(tlssetup)
		if err != nil {
			log.Println("cert: ", err.Error())
		}
		transport.TLSClientConfig = tlsconf
		apiconf.HttpClient.Transport = &transport
		apiconf.Scheme = "https"
		port = conf.CertConf.HTTPSPORT
	} else {
		port = conf.HTTPPORT
	}
	address := selectAvailableConsulServer(conf.ADDRESSES, port)
	apiconf.Address = fmt.Sprintf("%s:%d", address, port)
	apiconf.Datacenter = conf.DATACENTER
	client, _ := api.NewClient(apiconf)
	return client
}

func ParepareReg(ip string, name string, id string, port int) (reg *api.CatalogRegistration) {
	conf := g.Config()
	services := &api.AgentService{
		ID:      fmt.Sprintf("%s-%s", name, id),
		Service: name,
		Tags:    []string{name},
		Port:    port,
		Address: ip,
	}

	reg = &api.CatalogRegistration{
		Datacenter: conf.DATACENTER,
		Address:    selectAvailableConsulServer(conf.ADDRESSES, conf.HTTPPORT),
		Node:       conf.NODE,
		Service:    services,
	}
	return
}

func ParepareRegSer(ip string, name string, id string, port int) (reg *api.AgentServiceRegistration) {
	conf := g.Config()
	check := &api.AgentServiceCheck{
		TCP:      fmt.Sprintf("%s:%d", ip, port),
		Status:   api.HealthPassing,
		Interval: "20s",
		Timeout:  "3s",
	}
	reg = &api.AgentServiceRegistration{
		Address:           selectAvailableConsulServer(conf.ADDRESSES, conf.HTTPPORT),
		Check:             check,
		ID:                fmt.Sprintf("%s-%s", name, id),
		Name:              name,
		Tags:              []string{name},
		Port:              port,
		EnableTagOverride: false,
	}
	return
}

func ServiceRegister(ip string, name string, id string, port int, client *api.Client) (err error) {
	reg := ParepareRegSer(ip, name, id, port)
	// wmeta, _ := client.Catalog().Register(reg, nil)

	err = client.Agent().ServiceRegister(reg)
	return err
}

func QueryServies(servicesName string, client *api.Client) (services []*api.CatalogService) {
	catalog := client.Catalog()
	services, _, _ = catalog.Service("consul", "", nil)
	return
}

func selectAvailableConsulServer(consult []string, port int) string {
	address := "127.0.0.1"
	for _, addr := range consult {
		_, err := CheckConsulServer(addr, port)
		if err == nil || strings.Contains(err.Error(), "malformed HTTP response") {
			address = addr
			return address
		} else {
			log.Println("consul server check: ", err.Error())
		}
	}
	return address
}

func CheckConsulServer(host string, port int) (*http.Response, error) {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	url := fmt.Sprintf("http://%s:%d", host, port)
	resp, err := client.Get(url)
	return resp, err
}
