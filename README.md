# LeakIXClient

This is a Go CLI & library making queries to LeakIX easier.

## Command line usage

```sh
$ leakix -h
Usage of leakix: 
  ./leakix -q '*' -l 200

  -l int
        Limit results output (default 100)
  -q string
        Specify search query (default "*")
  -s string
        Specify scope (default "leak")
  -t string
        Specify output template (default "{{ .Ip }}:{{ .Port }}")

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
```

## Library usage

```golang
package main
import (
	"fmt"
	"github.com/LeakIX/LeakIXClient"
)

func main(){
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

```
