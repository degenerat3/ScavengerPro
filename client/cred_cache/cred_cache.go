//Package cred_cache This file holds the "object" information/functions needed for the credential cache
//@author: degenerat3
package cred_cache

import (
	"fmt"
)

//CredCache is the struct will be the "object" which holds cred info
type CredCache struct {
	Hostname    string
	Credentials []string
}

//Get_entries fetches all stored credentials
func (c *CredCache) Get_entries() []string {
	fmt.Printf("%s\n", c.Credentials)
	return c.Credentials
}

//Add_entry adds a credential set to the object
func (c *CredCache) Add_entry(cred string) {
	creds := c.Credentials
	creds = append(creds, cred)
	c.Credentials = creds
}

//Count_entries returns number of stored credentials
func (c *CredCache) Count_entries() int {
	return len(c.Credentials)
}

//Get_hostname returns the hostname associated with the cache
func (c *CredCache) Get_hostname() string {
	return (c.Hostname)
}

//Encrypt_entries uses AES to encrypt passwords pre-exfil
func (c *CredCache) Encrypt_entries() {
	// iterate through entries, encrypt each one
}
