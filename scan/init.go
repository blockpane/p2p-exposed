package scan

import (
	"fmt"
	p2px "github.com/blockpane/p2p-exposed"
	"gopkg.in/yaml.v2"
	"log"
)

var (
	Ports      map[uint32]string
	Signatures []Signature
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// YAML files are embedded at build time using go:embed
	var err error

	// TODO: advanced fingerprinting, sec warnings etc.
	//err = yaml.Unmarshal(p2px.Http, &Signatures)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	err = yaml.Unmarshal(p2px.Ports, &Ports)
	if err != nil {
		fmt.Println(string(p2px.Ports))
		log.Fatal(err)
	}
}
