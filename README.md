# p2p-exposed

This is a simple web service for checking what blockchain ports are exposed to the internet. It is useful for two reasons, not having P2P ports open can make peering more difficult, and **accidently exposing RPC** ports can create security risks.

## Supported chains

* Cosmos/Tendermint: use the '/cosmos' path
* Ethereum: use the '/eth' path
* Fantom: use the '/fantom' path

## How does it work?

From the host that will to be scanned, simply send an HTTP request to the service and it will print out the results. The scan uses a pre-defined list of ports and usually finishes in only a couple of seconds. There is a publicly available endpoint running at https://p2p.exposed

Example:

```
$ curl p2p.exposed/cosmos
a.b.c.d:80    - generic http
a.b.c.d:443   - tls http
a.b.c.d:1317  - cosmos rest
a.b.c.d:9090  - tendermint grpc
a.b.c.d:9091  - tendermint grpc web
a.b.c.d:26656 - tendermint p2p
a.b.c.d:26657 - tendermint rpc

done: 7 open ports.
```

## Contributions Welcome!

Currently this only supports 3 chains, but it is easy to add more. If you want to contribute, please submit a PR with the ports and a description and I'll add the necessary code. See [sigs/cosmos-ports.yml](sigs/cosmos-ports.yaml) for an example.