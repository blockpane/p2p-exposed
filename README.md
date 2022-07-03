# p2p-exposed

This is a simple web service for checking what blockchain ports are exposed to the internet. It is useful for two reasons, not having P2P ports open can make peering more difficult, and **accidently exposing RPC** ports can create security risks.

## Supported chains

* Cosmos/Tendermint: use the '/cosmos' path
* Ethereum: use the '/eth' path
* Fantom: use the '/fantom' path

## How does it work?

From the host that will to be scanned, simply send an HTTP request to the service and it will print out the results. The scan uses a pre-defined list of ports and usually finishes in only a couple of seconds. There is a publicly available endpoint running at https://p2p.exposed


## Contributions Welcome!

Currently this only supports 3 chains, but it is easy to add more. If you want to contribute, please submit a PR with the ports and a description and I'll add the necessary code. See [sigs/cosmos-ports.yml](sigs/cosmos-ports.yaml) for an example.

## Example:

```
$ curl p2p.exposed/cosmos
Scanning the following ports:
80: generic http
1080: generic http
3000: generic http
8080: generic http
8088: generic http
8887: generic http
8888: generic http
8889: generic http

9100: generic prometheus node exporter
22: generic ssh
2222: generic alt ssh

443: generic tls http
8443: generic tls http

16656: tendermint p2p
26656: tendermint p2p
36656: tendermint p2p
46656: tendermint p2p
56656: tendermint p2p

16657: tendermint rpc
26657: tendermint rpc
36657: tendermint rpc
46657: tendermint rpc
56657: tendermint rpc

1317: cosmos rest
2317: cosmos rest
3317: cosmos rest
4317: cosmos rest
5317: cosmos rest
10317: cosmos rest

16660: cosmos prometheus
26660: cosmos prometheus
36660: cosmos prometheus
46660: cosmos prometheus
56660: cosmos prometheus

6060: tendermint pprof

9090: tendermint grpc
19090: tendermint grpc
29090: tendermint grpc
39090: tendermint grpc
49090: tendermint grpc
59090: tendermint grpc

9091: tendermint grpc web
19091: tendermint grpc web
29091: tendermint grpc web
39091: tendermint grpc web
49091: tendermint grpc web
59091: tendermint grpc web
-----------------------------------------

a.b.c.d:80    - generic http
a.b.c.d:443   - generic tls http
a.b.c.d:1317  - cosmos rest
a.b.c.d:9090  - tendermint grpc
a.b.c.d:9091  - tendermint grpc web
a.b.c.d:26656 - tendermint p2p
a.b.c.d:26657 - tendermint rpc

done: 7 open ports.
```
