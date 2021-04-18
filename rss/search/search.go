package search

import (
	"log"
	"sync"
)

// a map of registered for searching
var matchers = make(map[string] Matcher)

func Run(searchItem string) error  {
	feeds, err := ReceiveFeed()
	if err != nil {
		log.Fatal(err)
		return err
	}

	results := make(chan *Result)

	// setup a waitGroup so we can process feed async
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds{
		matcher, exist := matchers[feed.Type]
		if !exist {
			matcher = matchers["default"]
		}

		go func(matcher2 Matcher, feed * Feed) {
			Match(matcher2, feed, searchItem, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		// close chan so that we can exit from display funtion
		close(results)
	}()

	display(results)
	return nil
}

func display(results chan *Result) {
	// the loop runs util results chan is closed
	for result := range results{
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

func Register(matcher Matcher, feedType string)  {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}