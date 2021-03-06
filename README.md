# LeakIXClient

This is a Go CLI & library making queries to LeakIX easier.

## leakix - Command line usage

```sh
$ leakix -h
Usage of leakix: 

  -j    JSON mode, (excludes -t)
  -l int
        Limit results output (default 100)
  -q string
        Search mode, specify search query (default "*")
  -r    Realtime mode, (excludes -q)
  -s string
        Specify scope (default "leak")
  -t string
        Specify output template (default "{{ .Ip }}:{{ .Port }}")

$ # Example query on the index
$ leakix -l 2 -q "protocol:web AND plugin:GitConfigPlugin" -t "{{ .Ip }}:{{ .Port }} : {{ .Data }}"
178.62.217.44:80 : Found git deployment configuration
[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = https://gitlab.com/lyranalytics/lyra-website.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "abdulrahman"]
	remote = origin
	merge = refs/heads/abdulrahman

2604:a880:800:a1::10f:1001:80 : Found git deployment configuration
[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = https://github.com/mautic/mautic.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "staging"]
	remote = origin
	merge = refs/heads/staging

$ # Stream results in realtime from the engine, no filtering
$ ./leakix -r -s services -l 0
14.167.7.149:81
54.249.38.136:9200
23.65.39.190:80
[2a01:4f8:10a:1b5a::2]:80
23.225.38.43:3306
210.16.68.51:80
...keeps streaming...
```

## Library usage

```golang
package main
import (
	"fmt"
	"github.com/LeakIX/LeakIXClient"
)

func DoSearch(){
	// Create a searcher
	LeakIXSearch := LeakIXClient.SearchResultsClient{
		Scope: "leak",
		Query: "protocol:kafka AND \"telecom_italia_data\"",
	}
	// Iterate, the lib will query further pages if needed
	for LeakIXSearch.Next() {
		// Use the result
		leak := LeakIXSearch.SearchResult()
		fmt.Printf("%s:%s - Country:%s\n", leak.Ip, leak.Port, leak.GeoLocation.CountryName)
	}
}


func LiveStream() {
	// Get a channel from the websocket
	serviceChannel, err := LeakIXClient.GetChannel("services")
	if err != nil {
		log.Println("Websocket connection error:")
		log.Fatal(err)
	}
	for {
		// Print everything received on the channel
		service := <- serviceChannel
		log.Println(service.Ip)
	}
}
```
