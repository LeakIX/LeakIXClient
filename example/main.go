package main
import (
	"fmt"
	"gitlab.nobody.run/tbi/LeakIXClient"
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
