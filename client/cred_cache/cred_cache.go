// This file holds the "object" information/functions needed for the credential cache
//@author: degenerat3
package cred_cache

import(
	"fmt"
)

type CredCache struct {
	Hostname	string
	Credentials	[]string
}

func (c *CredCache) Get_entries() []string{
	fmt.Printf("%s\n",c.Credentials)
	return c.Credentials
}

func (c *CredCache) Add_entry(cred string){
	creds := c.Credentials
	creds = append(creds, cred)
	c.Credentials = creds
}

func (c *CredCache) Count_entries() int{
	return len(c.Credentials)
}

func (c *CredCache) Get_hostname() string{
	return(c.Hostname)
}
