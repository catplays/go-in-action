package matcher

import (
	"catwang.com/go-in-action/rss/search"
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"regexp"
)

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rss implements Matcher interface
type RssMatcher struct {

}
func (rss RssMatcher) Search(feed *search.Feed, searchItem string) ([] *search.Result, error) {
	var results [] *search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)
	document, err := rss.retrive(feed)
	if err != nil {
		return nil, err
	}

	for _, item := range document.Channel.Item{
		matched, err := regexp.MatchString(searchItem, item.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content: item.Title,
			})
		}
		// Check the description for the search term.
		matched, err = regexp.MatchString(searchItem, item.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: item.Description,
			})
		}
	}
	return results, err
}

func (rss RssMatcher) retrive(feed *search.Feed) (*rssDocument, error) {
	if  feed.URI == "" {
		return nil, errors.New("no rss feed URI provided.")
	}
	defer func() {
		if err := recover(); err!= nil {
			log.Println("http error:%+v")
		}
	}()
	// get resp by feed url
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Http response error by url:%s, %+v", feed.URI, err)
		return nil, err
	}
	var document  rssDocument
	err = json.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

func init()  {
	var rssMatcher RssMatcher
	search.Register(rssMatcher, "rss")
}
