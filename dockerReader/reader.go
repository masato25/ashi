package dockerReader

import (
  "github.com/fsouza/go-dockerclient"
)

var dockersoc string = "unix:///var/run/docker.sock"
func ConatainerRead() (contains []docker.APIContainers, err error){
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
}

func GetPublicPort(contains []docker.APIContainers) (serobj []ServiesObj){
  if len(contains) == 0{
    return
  }else{
    for _, c := range contains{
      if c.State == "running"{
        var ser ServiesObj
        ser.Name = c.Names[0]
        ser.ID = c.ID[0:4]
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
