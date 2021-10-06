package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/LeakIX/LeakIXClient"
	"github.com/LeakIX/l9format"
	"log"
	"os"
	"text/template"
)

// It's all sequential, Couldn't care sorry :
var outputJson bool
var scope string
var query string
var liveStream bool
var outputTemplate string
var apiKey string
var endPoint string
var limit int
var tmpl *template.Template
var err error

func main() {
	//Config our app
	err = nil
	flag.StringVar(&scope, "s", "leak", "Specify scope")
	flag.StringVar(&query, "q", "*", "Search mode, specify search query")
	flag.BoolVar(&liveStream, "r", false, "Realtime mode, (excludes -q)")
	flag.BoolVar(&outputJson, "j", false, "JSON mode, (excludes -t)")
	flag.StringVar(&outputTemplate, "t", "{{ .Ip }}:{{ .Port }}", "Specify output template")
	flag.StringVar(&apiKey, "k", "", "Specify API key")
	flag.StringVar(&endPoint, "e", "https://leakix.net", "Leakix endpoint to use")

	flag.IntVar(&limit, "l", 100, "Limit results output")
	flag.Usage = func() {
		fmt.Printf("Usage of leakix: \n")
		fmt.Printf("  ./leakix -q '*' -s leaks -l 200\n\n")
		fmt.Printf("  ./leakix -r -s services -l 0\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	tmpl, err = template.New("output").Parse(outputTemplate + "\n")
	if err != nil {
		log.Println("Template error :")
		log.Fatal(err)
	}

	// Run the right thing
	if liveStream {
		LiveStream()
	} else {
		Search()
	}

}

func Search() {
	searcher := LeakIXClient.SearchResultsClient{
		Scope:    scope,
		Query:    query,
		ApiKey:   apiKey,
		Endpoint: endPoint,
	}
	count := 0
	for searcher.Next() {
		count++
		OutputSearchResult(searcher.SearchResult())
		if count >= limit || count >= 10000 {
			os.Exit(0)
		}
	}
	if searcher.LastError != nil {
		log.Println("finished with errors: " + searcher.LastError.Error())
	}
}

func LiveStream() {
	count := 0
	searcher := LeakIXClient.SearchResultsClient{
		ApiKey:   apiKey,
		Endpoint: endPoint,
	}
	serviceChannel, err := searcher.GetChannel(scope)
	if err != nil {
		log.Println("Websocket connection error:")
		log.Fatal(err)
	}
	for {
		service := <-serviceChannel
		count++
		OutputSearchResult(service)
		if count >= limit && limit != 0 {
			os.Exit(0)
		}
	}
}

func OutputSearchResult(searchResult l9format.L9Event) {
	if outputJson {
		jsonBody, err := json.Marshal(searchResult)
		if err != nil {
			log.Println("JSON error :")
			log.Fatal(err)
		}
		fmt.Println(string(jsonBody))
	} else {
		err := tmpl.Execute(os.Stdout, searchResult)
		if err != nil {
			log.Println("Template error :")
			log.Fatal(err)
		}
	}
}
