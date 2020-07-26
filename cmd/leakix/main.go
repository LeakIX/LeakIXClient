package main

import (
	"flag"
	"fmt"
	"github.com/LeakIX/LeakIXClient"
	"text/template"
	"log"
	"os"
)

func main() {
	var scope string
	flag.StringVar(&scope, "s", "leak", "Specify scope, default to leak")
	var query string
	flag.StringVar(&query, "q", "*", "Specify search query, default to *")
	var outputTemplate string
	flag.StringVar(&outputTemplate, "t", "{{ .Ip }}:{{ .Port }}", "Specify output template, default to \"{{ .Ip }}:{{ .Port }}\"")
	var limit int
	flag.IntVar(&limit, "l", 100, "Limit results output, default to 100")

	flag.Usage = func() {
		fmt.Printf("Usage of leakix: \n")
		fmt.Printf("  ./leakix -q '*' -l 200\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	searcher := LeakIXClient.SearchResultsClient{
		Scope:         scope,
		Query:         query,
	}

	tmpl, err := template.New("output").Parse(outputTemplate + "\n")
	if err != nil {
		log.Println("Template error :")
		log.Fatal(err)
	}
	count := 0
	for searcher.Next() {
		count++
		err = tmpl.Execute(os.Stdout, searcher.SearchResult())
		if err != nil {
			log.Println("Template error :")
			log.Fatal(err)
		}
		if count >= limit {
			os.Exit(0)
		}
	}
}
