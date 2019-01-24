package main
import "ScavengerPro/client/cred_cache"

func main(){
	c := cred_cache.CredCache {
		Hostname: "Test",
		Credentials: []string{"p1","p2","p3"},
	}
	c.Get_entries()
	c.Add_entry("p4")
	c.Get_entries()
	c.Count_entries()
}
