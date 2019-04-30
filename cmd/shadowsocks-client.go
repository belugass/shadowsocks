package main

import (
	"flag"
	"log"
	"os"
	"path"
	"shadowsocks/shadowsocks"
)

// shadowsocks client implement
func main() {

	var configFile string
	var cmdConfig *shadowsocks.Config

	flag.StringVar(&configFile, "c", "config.json", "config file ")
	flag.StringVar(&cmdConfig.Local, "b", "", "local address,listen only to this address if set")
	flag.IntVar(&cmdConfig.LocalPort, "l", 0, "local socks5 proxy port")
	flag.StringVar(&cmdConfig.Server, "s", "", "server socks5 address")
	flag.IntVar(&cmdConfig.ServerPort, "p", 0, "server socks5 proxy port")
	flag.IntVar(&cmdConfig.Timeout, "t", 100, "timeout in second")
	flag.StringVar(&cmdConfig.Method, "m", "", "encryption method, default: aes-256-cfb")

	flag.Parse()

	exists, err := shadowsocks.IsExist(configFile)

	workpath := path.Dir(os.Args[0])

	if (err != nil || !exists) && workpath != "" && workpath != "." {
		oldConfig := configFile
		configFile = path.Join(workpath, "config.json")
		log.Printf("%s is not found, try config file %s\n", oldConfig, configFile)
	}

	config ,err := shadowsocks.ParseConfig(configFile)
	if err != nil {
		config = cmdConfig
		if os.IsNotExist(err) {
			log.Fatalf("config file path: %s is not found",configFile)
		}
	}else{
		shadowsocks.UpdateConfig(config,cmdConfig)
	}

	// init default config params
	if config.Method == "" {
		config.Method = "aes-256-cfb"
	}
	if config.Local == "" {
		config.Local = "127.0.0.1"
	}
	if config.LocalPort == 0 {
		config.LocalPort = 1080
	}

	checkOptions(config)


}

func checkOptions(config *shadowsocks.Config) {
	if config.Server == "" || config.ServerPort == 0 || config.Password == ""{
		log.Fatalf("server address,server port and password is must")
	}
}
