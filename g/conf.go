package g

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"regexp"
	"sync"
)

type Gconfig struct {
	ADDRESSES  []string
	IP         string
	DATACENTER string
	NODE       string
	HTTPPORT   int
	DOCKERSOC  string
}

var (
	gconfig    *Gconfig
	configLock = new(sync.RWMutex)
)

func Config() *Gconfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return gconfig
}

func Set(f string, confpath string) {
	viper.SetConfigType("json")
	viper.AddConfigPath(confpath)
	viper.SetConfigName(f)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	var c Gconfig
	c.ADDRESSES = regexp.MustCompile("\\s*,\\s*").Split(viper.Get("consul_server_addresses").(string), -1)
	c.IP = viper.Get("ip").(string)
	if c.IP == "" {
		com := exec.Command("hostname -I | awk -F \" \" '{ print $1 }'")
		output, err := com.Output()
		if err != nil {
			log.Println("can not get ip \"from hostname -I\"")
		}
		c.IP = string(output)
	}
	c.DATACENTER = viper.Get("datacenter").(string)
	c.NODE = viper.Get("node").(string)
	c.HTTPPORT = int(viper.Get("http_port").(float64))
	c.DOCKERSOC = viper.Get("dockersoc").(string)
	gconfig = &c
}
