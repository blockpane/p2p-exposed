package main

import (
	"flag"
	"fmt"
	"github.com/blockpane/p2p-exposed/scan"
	"log"
	"net"
	"net/http"
	"net/netip"
	"strings"
	"sync"
	"time"
)

/*
 p2px-server starts an HTTP server. When a client connects to /cosmos it will scan for common cosmos ports, and
 return the results.
*/

const link = "https://github.com/blockpane/p2p-exposed"

func main() {
	var port int
	var xForwarded string
	var useXForwarded bool

	flag.IntVar(&port, "p", 80, "port to listen on")
	flag.StringVar(&xForwarded, "h", "X-Forwarded-For", "optional: trusted X-Forwarded-For Header")
	flag.BoolVar(&useXForwarded, "x", false, "Use the X-Forwarded-For header instead of remote address (reverse proxy)")
	flag.Parse()

	wait := make(map[string]*time.Time)
	mux := sync.Mutex{}

	// this occasionally cleans the rate-limiting list.
	go func() {
		for {
			time.Sleep(time.Minute)
			if len(wait) == 0 {
				continue
			}
			mux.Lock()
			for k, v := range wait {
				if v.Before(time.Now()) {
					delete(wait, k)
				}
			}
			mux.Unlock()
		}
	}()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		var chain string
		switch {
		case strings.HasPrefix(request.URL.Path, "/cosmos"):
			chain = "cosmos"
		case strings.HasPrefix(request.URL.Path, "/eth"):
			chain = "eth"
		case strings.HasPrefix(request.URL.Path, "/fantom"):
			chain = "fantom"
		default:
			writer.WriteHeader(404)
			writer.Write([]byte("not found\n"))
			return
		}

		var remoteIp string
		var e error

		if useXForwarded {
			remoteIp = request.Header.Get(xForwarded)
			_, e = netip.ParseAddr(remoteIp)

		} else {
			remoteIp, _, e = net.SplitHostPort(request.RemoteAddr)
		}

		if e != nil {
			log.Println(remoteIp, e)
			writer.WriteHeader(500)
			writer.Write([]byte("internal error\n"))
			return
		}

		writer.Header().Set("X-Powered-By", link)
		mux.Lock()
		locked := wait[remoteIp] != nil && wait[remoteIp].After(time.Now())
		mux.Unlock()
		if locked {
			writer.Write([]byte("Only one request per minute is allowed, please try again later\n"))
			log.Println(remoteIp, "blocked: too soon since last scan")
			return
		}
		next := time.Now().Add(time.Minute)
		mux.Lock()
		wait[remoteIp] = &next
		mux.Unlock()
		log.Println(remoteIp, "requested a scan for", chain)
		report := scan.PortScan(remoteIp, chain)
		writer.Write(report)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
