package g

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sync"

	"github.com/spf13/viper"
)

type Gconfig struct {
	ADDRESSES  []string
	IP         string
	DATACENTER string
	NODE       string
	HTTPPORT   int
	DOCKERSOC  string
	TOKEN      string
	CertConf   *CertConfig
}

type CertConfig struct {
	Enable    bool
	CAFile    string
	CertFile  string
	KeyFile   string
	HTTPSPORT int
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
	c.DATACENTER = viper.GetString("datacenter")
	c.NODE = viper.GetString("node")
	c.HTTPPORT = int(viper.GetFloat64("http_port"))
	c.DOCKERSOC = viper.GetString("dockersoc")
	c.TOKEN = viper.GetString("token")
	var certconf CertConfig
	certconf.CAFile = viper.GetString("cert.cafile")
	certconf.CertFile = viper.GetString("cert.certfile")
	certconf.KeyFile = viper.GetString("cert.keyfile")
	certconf.Enable = viper.GetBool("cert.enable")
	certconf.HTTPSPORT = viper.GetInt("cert.https_port")
	c.CertConf = &certconf
	gconfig = &c
}
