# LeakIXClient

This is a Go library making queries to LeakIX easier.


## Usage

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
