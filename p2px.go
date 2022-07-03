package p2px

import (
	_ "embed"
)

/*
p2px is a web service that will perform a limited port-scan of the host connecting to it and report back on any commonly
used Cosmos services that are available.
*/

// //go:embed sigs/http.yaml
// var Http []byte

//go:embed sigs/cosmos-ports.yaml
var CosmosPorts []byte

//go:embed sigs/fantom-ports.yaml
var FantomPorts []byte

//go:embed sigs/eth-ports.yaml
var EthPorts []byte
