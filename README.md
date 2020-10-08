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

## leakix-ns - Command line usage

```
$ ./bin/leakix-ns-linux-64 
Usage of leakix-dns: 
  ./leakix-ns -d <domain> -l 200

  -d string
        Specify domain or IP
  -j    JSON mode, (excludes -t)
  -l int
        Limit results output (default 100)
$ ./bin/leakix-ns-linux-64 -d microsoft.com -l 3000
PTR records :
[22-09-2020 03:50] sophosent.microsoft.com <- 157.58.217.12
[04-10-2020 21:08] tide131.microsoft.com <- 213.199.128.156
[01-10-2020 12:42] 52-114-124-186.relay.teams.microsoft.com <- 52.114.124.186
[27-09-2020 06:11] dbecbip01.one.microsoft.com <- 213.199.129.203
[04-10-2020 13:15] duschy644p52v2-portal.msftvpn.ras.microsoft.com <- 157.58.214.101
[02-10-2020 16:10] timestamping4096.microsoft.com <- 207.46.153.123
[02-10-2020 13:07] invoicepresentmentjournal.co1.cp.microsoft.com <- 65.55.120.33
[28-09-2020 05:28] 52-114-124-198.relay.teams.microsoft.com <- 52.114.124.198
[27-09-2020 23:11] dm2.usmevpn.microsoft.com <- 131.253.121.62
[25-09-2020 12:56] blugro7gms.groove.microsoft.com <- 134.170.3.177
[27-09-2020 20:35] 52-114-125-70.relay.teams.microsoft.com <- 52.114.125.70
[22-09-2020 01:46] bn2-dspcdn.tlu.dl.delivery.mp.microsoft.com <- 40.77.228.30
[03-10-2020 10:39] accounts.bn2.cp.microsoft.com <- 40.77.228.40
[01-10-2020 19:53] duschy644p52v2.msftvpn-alt.ras.microsoft.com <- 157.58.214.99
[23-09-2020 11:27] primary.exchange.microsoft.com <- 157.58.197.9
[03-10-2020 09:33] invoicepresentment.dm2.cp.microsoft.com <- 65.55.145.65
[01-10-2020 16:43] origin-internal-dev.support.services.microsoft.com <- 40.77.232.18
[01-10-2020 04:14] invoicepresentment.co1.cp.microsoft.com <- 65.55.120.34
[28-09-2020 01:50] co1.gmevpn.microsoft.com <- 65.55.5.44
[06-10-2020 00:53] 52-114-124-181.relay.teams.microsoft.com <- 52.114.124.181
[03-10-2020 01:04] pages.bing-email.microsoft.com <- 13.111.35.40
[01-10-2020 01:31] accounts.bkgprocessing.cy2.cp.microsoft.com <- 40.77.232.79
[01-10-2020 00:46] 52-114-125-17.relay.teams.microsoft.com <- 52.114.125.17
[27-09-2020 21:52] 52-114-124-88.relay.teams.microsoft.com <- 52.114.124.88
[30-09-2020 12:22] 52-114-124-202.relay.teams.microsoft.com <- 52.114.124.202
[22-09-2020 22:54] staff.microsoft.com <- 209.141.49.113
[22-09-2020 03:50] staff.microsoft.com <- 209.141.61.131
[22-09-2020 10:44] click.bing-email.microsoft.com <- 13.111.36.33
Forward records :
[06-10-2020 02:26] office2018microsoft.com -> 13.77.161.179
[06-10-2020 02:26] office2018microsoft.com -> 104.215.148.63
[06-10-2020 02:26] office2018microsoft.com -> 40.113.200.201
[06-10-2020 02:26] office2018microsoft.com -> 104.215.148.63
[06-10-2020 02:26] office2018microsoft.com -> 13.77.161.179
[06-10-2020 02:26] office2018microsoft.com -> 40.113.200.201
[06-10-2020 02:26] office2018microsoft.com -> 40.112.72.205
[06-10-2020 02:26] office2018microsoft.com -> 40.112.72.205
[06-10-2020 02:26] office2018microsoft.com -> 40.76.4.15
[06-10-2020 02:26] office2018microsoft.com -> 40.76.4.15
[05-10-2020 02:50] www.microsoft.com -> 2600:140a:a000:299::356e
[05-10-2020 02:50] www.microsoft.com -> 2600:140a:a000:299::356e
[05-10-2020 02:50] www.microsoft.com -> 2600:140a:a000:293::356e
[05-10-2020 02:50] www.microsoft.com -> 104.77.222.2
[05-10-2020 02:50] www.microsoft.com -> 2600:140a:a000:293::356e
[05-10-2020 02:50] www.microsoft.com -> 104.77.222.2
[05-10-2020 02:50] es.microsoft.com -> 51.143.57.13
[05-10-2020 02:50] es.microsoft.com -> 51.143.57.13
[05-10-2020 02:50] mail.microsoft.com -> 157.58.197.10
[05-10-2020 02:50] mail.microsoft.com -> 167.220.71.19
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
