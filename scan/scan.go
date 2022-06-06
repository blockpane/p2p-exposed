package scan

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

type result struct {
	Open bool
	Port uint32
}

// PortScan accepts an IP/hostname and returns scan results.
func PortScan(address string) []byte {
	doneChan := make(chan interface{}, 1)
	foundChan := make(chan result)
	wg := &sync.WaitGroup{}
	wg.Add(len(Ports))

	finalResults := make([]result, 0)

	go func() {
		counter := len(Ports)
		for f := range foundChan {
			counter -= 1
			if f.Open {
				finalResults = append(finalResults, f)
			}
			if counter == 0 {
				close(doneChan)
			}

		}
	}()

	for k := range Ports {
		go openTcp(address, k, wg, foundChan)
		time.Sleep(20 * time.Millisecond)
	}

	// wait for results
	wg.Wait()
	<-doneChan

	sort.Slice(finalResults, func(i, j int) bool {
		return finalResults[i].Port < finalResults[j].Port
	})

	buf := bytes.NewBuffer(nil)
	for _, f := range finalResults {
		buf.WriteString(fmt.Sprintf("%s:%-5d - %s\n", address, f.Port, Ports[f.Port]))
	}
	buf.WriteString(fmt.Sprintf("\ndone: %d open ports.\n", len(finalResults)))

	return buf.Bytes()
}

// openTcp opens a tcp connection, and closes it. If it connected it returns true.
func openTcp(ipAddress string, port uint32, wg *sync.WaitGroup, results chan result) {
	defer wg.Done()
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", ipAddress, port))
	if err != nil {
		results <- result{false, port}
		return
	}
	_ = conn.Close()
	results <- result{true, port}
}
