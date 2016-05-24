# Ashi
![](./asset/img/icon.jpg)

## Conpect
![](./asset/img/conpect.png)

Every Docker host has own consul (server/client), but no one can help docker to update new docker container to consul services list. On the low level of Consul, It's using [ectd](https://github.com/coreos/etcd) to store shared information.

![](./asset/img/dockerps.png)

Docker-cli can know the status of each container, so ashi using Docker API to get the container info and update it into consul world.


## How set up conf
`cd conf && cp ashi_example.json ashi.json`

* conf:
  * consul_server_addresses -> consul server list
  * ip -> ip address of this docker host machine
  * http_port -> web port of consul http api [default is 8500]
  * datacenter -> datacenter of consul server
  * node -> node name of consul server
  * dockersoc ->  Docker daemon to listen on.
  * token -> [ToDo]
  * cert -> TLS support

* A sample server conf for TLS
```
{
    "bootstrap": true,
    "server": true,
    "datacenter": "tw",
    "data_dir": "./data",
    "ui_dir": "./ui",
    "log_level": "INFO",
    "enable_syslog": true,
    "bind_addr": "192.168.2.105",
    "ports": {
      "https": 8888
    },
    "addresses" : {
      "https": "192.168.2.105"
    },
    "ca_file": "./ssl/ca/ca.crt",
    "cert_file": "./ssl/server/server.crt",
    "key_file": "./ssl/server/server.key"
}
```
