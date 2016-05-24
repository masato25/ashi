package main

import (
	"log"

	"github.com/masato25/ashi/consulAshi"
	"github.com/masato25/ashi/dockerReader"
	"github.com/masato25/ashi/g"
)

func main() {
	//set conf paht
	g.Set("ashi", "./conf")
	conf := g.Config()
	// //get docker list
	container := dockerReader.GetContainers(conf.DOCKERSOC)
	client := consulAshi.Client()
	for _, c := range container {
		log.Println("will register", c.ID, c.Name, c.OnePort, c.IP)
		err := consulAshi.ServiceRegister(conf.IP, c.Name, c.ID, c.OnePort, client)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
