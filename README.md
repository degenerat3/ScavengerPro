# ScavengerPro
Read passwords from our collectors, store them in a cache, ship it to a webserver

## Deploy notes 
2 things need to be updated for each deployment.  
&nbsp;  
First is line 23 in `shipper.go`. That's the list of servers (for use of multiple IPs via [theArk](https://github.com/ritredteam/theark)), but if that's not being utilized you can just put one item in the list.  
&nbsp;  
Second thing to update is line 34 in `scavenger_pro.go`. This is the list of "dump" files that will be watched for credential entries. Any collector that is being deployed will have an output/dump file. List those here. For each file path/name, append ":def" to it since all collectors at this point utilize the default credential format. 
&nbsp; 
There  is a third, optional, variable that can be updated is "CacheSize" on line 13 of `scavenger_pro.go`. This is the number of collected credentials that are requried before exfiltration is executed. A higher cache size will result in less traffic, but obviously a delay in getting the credentials.
