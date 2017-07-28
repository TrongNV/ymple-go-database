package package_elast

import (
	"gopkg.in/olivere/elastic.v3"
	"fmt"
	"reflect"
	"encoding/json"
	//"os"
)

type Tweet struct {
	User    string
	Message string
}

func Read() {

	// Create a client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	const s = "notest"

	if s == "test" {
		_, err = client.CreateIndex("twitter2").Do()
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	// Add a document to the index [START]
	tweet := Tweet{User: "olivere", Message: "Take Five 2"}

	_, err = client.Index().
		Index("twitter2").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}




	//termQuery := elastic.NewMatchQuery("User", "olive")

	termQuery := elastic.NewRegexpQuery("User", "olive.*")

	searchResult, err := client.Search().
		Index("twitter2").// search in index "twitter"
		Query(termQuery).// specify the query
		Sort("User", true).// sort by "user" field, ascending
		From(0).Size(10).// take documents 0-9
		Pretty(true).// pretty print request and response JSON
		Do()                // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.

	fmt.Printf("\n\r")

	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	fmt.Printf("\n\r")

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits != nil {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	// Delete the index again

	if s == "test" {

		_, err = client.DeleteIndex("twitter").Do()
		if err != nil {
			// Handle error
			panic(err)
		}
	}
}