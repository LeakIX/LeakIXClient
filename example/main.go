package main
import (
	"fmt"
	"github.com/LeakIX/LeakIXClient"
)

func main(){
	// Create a searcher
	LeakIXSearch := LeakIXClient.SearchResultsClient{
		Scope: "leak",
		Query: "+protocol:elasticsearch  +\"telecom_italia_data\"",
	}
	// Iterate, the lib will query further pages if needed
	for LeakIXSearch.Next() {
		// Use the result
		leak := LeakIXSearch.SearchResult()
		fmt.Printf("%s:%s - Country:%s\n", leak.Ip, leak.Port, leak.GeoIp.CountryName)
	}
}
