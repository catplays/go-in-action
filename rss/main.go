package main

import (
	_ "catwang.com/go-in-action/rss/matcher"
	"catwang.com/go-in-action/rss/search"
	"log"
	"os"
)

func init()  {
	// change log to std
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
