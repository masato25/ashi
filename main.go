package main

import (
  "github.com/hashicorp/consul/api"
  "fmt"
)

func main(){
  client, err := api.NewClient(api.DefaultConfig())
  catalog := client.Catalog()
  if err != nil {
      panic(err)
  }
  services, meta, err := catalog.Service("consul", "", nil)
  fmt.Printf("%v %v %v", services[0].ServiceAddress, meta, err)
}
