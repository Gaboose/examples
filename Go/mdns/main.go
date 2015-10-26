package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Gaboose/mdns"
)

func publish(port int, tag string) (*mdns.Server, error) {
	// Setup our service export
	host, _ := os.Hostname()
	// Concatenating host with port allows to discover several instances on the same machine
	host = fmt.Sprintf("%s.%d", host, port)
	info := []string{"My awesome service"}
	service, err := mdns.NewMDNSService(host, tag, "", "", port, nil, info)
	if err != nil {
		return nil, err
	}
	fmt.Printf("just created service %v\n", service)

	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	return server, err
}

func lookup(tag string) {
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	// Start the lookup
	mdns.Lookup(tag, entriesCh)
	close(entriesCh)
}

func main() {
	pub := flag.Bool("pub", false, "")
	port := flag.Int("port", 8000, "")
	tag := flag.String("tag", "foobar", "")
	flag.Parse()
	if *pub {
		server, err := publish(*port, *tag)
		if err != nil {
			log.Fatal(err)
			return
		}
		select {}
		server.Shutdown()
	} else {
		lookup(*tag)
	}
}
