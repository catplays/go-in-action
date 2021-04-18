package search

import "log"

//结果
type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchItem string) ([] *Result, error)
}

func Match(matcher Matcher, feed *Feed, searchItem string, results chan <- *Result )  {
	searchResults, err := matcher.Search(feed, searchItem)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, result := range searchResults{
		// put search result into chan
		results <- result
	}

}
