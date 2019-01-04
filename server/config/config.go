package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	Executor Executor
}

type Executor struct {
	Host string
	Port string
}

type conf struct {
	Executor Executor
}

var Config *conf

func Init() {
	var server Server
	if err := envconfig.Process("server", &server); err != nil {
		log.Fatal(err)
		panic(err)
	}

	Config = &conf{
		Executor: server.Executor,
	}

}
