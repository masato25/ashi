package consulAshi

import (
  "github.com/hashicorp/consul/api"
)

var client *api.Client
func Client() (*api.Client){
  client, _ = api.NewClient(api.DefaultConfig())
  return client
}


func QueryServies(servicesName string, client *api.Client ) (services []*api.CatalogService){
  catalog := client.Catalog()
  services, _, _ = catalog.Service("consul", "", nil)
  return
}
