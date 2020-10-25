package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/LeakIX/LeakIXClient"
	"os"
	"strings"
)

func main() {
	//Config our app
	app := App{}
	flag.StringVar(&app.Domain, "d", "", "Specify domain")
	flag.BoolVar(&app.OutputJson, "j", false, "JSON mode, (excludes -t)")
	flag.IntVar(&app.Limit, "l", 100, "Limit results output")
	flag.StringVar(&app.ApiKey, "k", "", "API Key")
	flag.Usage = func() {
		fmt.Printf("Usage of leakix-ns: \n")
		fmt.Printf("  ./leakix-ns -d <domain> -l 200\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(app.Domain) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	app.Run()
}

type App struct {
	Domain     string
	OutputJson bool
	Limit      int
	Searcher   *LeakIXClient.SearchResultsClient
	Reverse    map[string][]LeakIXClient.SearchResult
	Forward    map[string][]LeakIXClient.SearchResult
	ApiKey     string
}

func (app *App) Run() {
	app.Searcher = &LeakIXClient.SearchResultsClient{
		Scope:    "service",
		Query:    fmt.Sprintf("hostname:\"%s\" OR reverse:\"%s\" OR ip:\"%s\"", app.Domain, app.Domain, app.Domain),
		ApiKey:   app.ApiKey,
		Endpoint: "https://leakix.net",
	}
	app.Reverse = make(map[string][]LeakIXClient.SearchResult)
	app.Forward = make(map[string][]LeakIXClient.SearchResult)
	count := 0
	for app.Searcher.Next() {
		if !strings.Contains(app.Searcher.SearchResult().Reverse, app.Domain) &&
			!strings.Contains(app.Searcher.SearchResult().Hostname, app.Domain) &&
			!strings.Contains(app.Searcher.SearchResult().Ip, app.Domain) {
			continue
		}
		count++
		if count > app.Limit {
			break
		}
		if app.OutputJson {
			jsonLine, _ := json.Marshal(app.Searcher.SearchResult())
			fmt.Println(string(jsonLine))
			continue
		}
		reverse := strings.TrimRight(app.Searcher.SearchResult().Reverse, ".")
		ip := app.Searcher.SearchResult().Ip
		hostname := app.Searcher.SearchResult().Hostname

		if hostname != ip && len(hostname) > 2 && (strings.Contains(hostname, app.Domain) || app.Domain == ip) {
			app.Forward[hostname] = append(app.Forward[hostname], app.Searcher.SearchResult())
		}
		if len(reverse) > 1 && (strings.Contains(reverse, app.Domain) || app.Domain == ip) {
			app.Reverse[reverse] = append(app.Reverse[reverse], app.Searcher.SearchResult())
		}
	}
	if app.OutputJson {
		os.Exit(0)
	}
	fmt.Println("PTR records :")
	for reverseName, results := range app.Reverse {
		for _, result := range results {
			fmt.Printf("[%s] %s <- %s", result.Time.Format("02-01-2006 15:04"), reverseName, result.Ip)
			if len(result.Hostname) > 1 && result.Hostname != result.Ip {
				fmt.Printf(" -> %s", result.Hostname)
			}
			fmt.Println()
		}
	}
	fmt.Println("Forward records :")

	for forwardName, results := range app.Forward {
		for _, result := range results {
			fmt.Printf("[%s] %s -> %s", result.Time.Format("02-01-2006 15:04"), forwardName, result.Ip)
			if len(result.Reverse) > 1 {
				fmt.Printf(" <- %s", result.Reverse)
			}
			fmt.Println()
		}
	}
}
