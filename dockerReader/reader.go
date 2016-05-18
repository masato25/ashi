package dockerReader

import (
  "github.com/fsouza/go-dockerclient"
  "regexp"
  "log"
)


func GetContainers(dockersoc string) ([]ServiesObj){
  container, err := GetContainerList(dockersoc)
  if(err != nil){
    log.Println(err.Error())
  }
  return getContainerObj(container)
}

func GetContainerList(dockersoc string) (contains []docker.APIContainers, err error) {
  var client *docker.Client
  client, err = docker.NewClient(dockersoc)
  if err != nil{
    return
  }
  contains, err = client.ListContainers(docker.ListContainersOptions{All: false})
  return
}

type ServiesObj struct{
  Name string
  ID string
  Ports []int
  OnePort int
  IP string
}

func getContainerObj(containers []docker.APIContainers) (serobj []ServiesObj){
  if len(containers) == 0{
    return
  }else{
    for _, c := range containers{
      if c.State == "running"{
        var ser ServiesObj
        ser.Name = getName(c.Names[0])
        ser.ID = c.ID[0:4]
        ser.IP = getDokcerIpv4(c.Networks.Networks)
        ser.OnePort = int(c.Ports[0].PublicPort)
        var ports []int
        for _, p := range c.Ports{
          ports = append(ports, int(p.PublicPort))
        }
        ser.Ports = ports
        serobj = append(serobj, ser)
      }
    }
    return
  }
}

func getName(name string)(string){
  return regexp.MustCompile("^\\W+").ReplaceAllString(name, "")
}

func getDokcerIpv4(dockernet map[string]docker.ContainerNetwork) (string){
  return dockernet["bridge"].IPAddress
}
