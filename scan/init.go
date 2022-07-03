package scan

import (
	"fmt"
	p2px "github.com/blockpane/p2p-exposed"
	"gopkg.in/yaml.v2"
	"log"
	"sync"
)

var (
	CosmosPortsMap map[uint32]string
	FantomPortsMap map[uint32]string
	EthPortsMap    map[uint32]string

	cosmosMu sync.RWMutex
	fantomMu sync.RWMutex
	ethMu    sync.RWMutex

	// Signatures []Signature // TODO
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error

	// TODO: advanced fingerprinting, sec warnings etc.
	// err = yaml.Unmarshal(p2px.Http, &Signatures)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	err = yaml.Unmarshal(p2px.CosmosPorts, &CosmosPortsMap)
	if err != nil {
		fmt.Println(string(p2px.CosmosPorts))
		log.Fatal(err)
	}

	err = yaml.Unmarshal(p2px.EthPorts, &EthPortsMap)
	if err != nil {
		fmt.Println(string(p2px.EthPorts))
		log.Fatal(err)
	}

	err = yaml.Unmarshal(p2px.FantomPorts, &FantomPortsMap)
	if err != nil {
		fmt.Println(string(p2px.FantomPorts))
		log.Fatal(err)
	}
}
