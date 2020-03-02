package main

import (
	"flag"
	"log"
	"network/skydive-visualizer-go/server"
	skydiveClient "network/skydive-visualizer-go/skydive"
	"network/skydive-visualizer-go/source"
	"network/skydive-visualizer-go/source/chef"
	"network/skydive-visualizer-go/source/dns"
	"network/skydive-visualizer-go/source/ipam"
	"network/skydive-visualizer-go/source/ports"
	skydiveSource "network/skydive-visualizer-go/source/skydive"
	"time"
)

func main() {
	configPath := flag.String("config", "config.yml", "")
	flag.Parse()

	cfg, err := GetConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	skydiveClient := skydiveClient.New(cfg.Skydive.URL)

	var src source.Source
	src = skydiveSource.NewSkydive(skydiveClient)
	src = dns.New(src)
	src = ports.New(src)
	src, err = ipam.New(src, cfg.IPAM)
	if err != nil {
		log.Fatal(err)
	}
	src, err = chef.New(src, cfg.Chef)
	if err != nil {
		log.Fatal(err)
	}

	src = source.NewPeriodic(src, 10*time.Minute)

	server := server.New(cfg.Server.Listen, src)
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
