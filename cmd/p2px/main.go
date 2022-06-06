package main

/*
p2px scans an IP address for common cosmos ports, it is mostly intended for testing p2px-server
*/

import (
	"fmt"
	"github.com/blockpane/p2p-exposed/scan"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please provide an IP address or hostname as the only argument")
	}

	fmt.Println(string(scan.PortScan(os.Args[1])))
}
