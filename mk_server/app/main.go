package main

import (
	"github.com/BurntSushi/toml"
	"mk_server/app/server"
)

func main() {
	path := "/Users/zhangzhao/goproj/goexample/mk_server/config/conf.toml"
	conf, err := ReadConfig(path)
	if err != nil {
		println("invalid config", err.Error())
		return
	}
	srv := server.NewZHttpServer(conf.Addr, 5)

	srv.SetRouter(server.NewHttpRouter())
	srv.StartServer()
}


type ServerConfig struct {
	Addr string `toml:"addr"`
}

func ReadConfig(path string) (*ServerConfig, error) {
	config := &ServerConfig{}
	_, err := toml.DecodeFile(path, config)
	return config, err
}





