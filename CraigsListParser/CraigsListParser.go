package CraigsListParser

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func Search(city string, category string, searchTerm string) (*SearchResult, error) {
	resp, err := http.Get(fmt.Sprintf("https://%s.craigslist.org/search/%s?format=rss&query=%s", city, category, searchTerm))
	if err != nil {
		return nil, err
	}
	searchResult := &SearchResult{}
	err = xml.NewDecoder(resp.Body).Decode(searchResult)
	if err != nil {
		return nil, err
	}
	return searchResult, nil
}

type SearchResult struct {
	XMLName xml.Name `xml:"RDF"`
	Text    string   `xml:",chardata"`
	Rdf     string   `xml:"rdf,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Enc     string   `xml:"enc,attr"`
	Ev      string   `xml:"ev,attr"`
	Content string   `xml:"content,attr"`
	Dcterms string   `xml:"dcterms,attr"`
	Syn     string   `xml:"syn,attr"`
	Dc      string   `xml:"dc,attr"`
	Taxo    string   `xml:"taxo,attr"`
	Admin   string   `xml:"admin,attr"`
	Channel struct {
		Text            string `xml:",chardata"`
		About           string `xml:"about,attr"`
		Title           string `xml:"title"`
		Link            string `xml:"link"`
		Description     string `xml:"description"`
		Language        string `xml:"language"`
		Rights          string `xml:"rights"`
		Publisher       string `xml:"publisher"`
		Creator         string `xml:"creator"`
		Source          string `xml:"source"`
		Type            string `xml:"type"`
		UpdateBase      string `xml:"updateBase"`
		UpdateFrequency string `xml:"updateFrequency"`
		UpdatePeriod    string `xml:"updatePeriod"`
		Items           struct {
			Text string `xml:",chardata"`
			Seq  struct {
				Text string `xml:",chardata"`
				Li   []struct {
					Text     string `xml:",chardata"`
					Resource string `xml:"resource,attr"`
				} `xml:"li"`
			} `xml:"Seq"`
		} `xml:"items"`
	} `xml:"channel"`
	Item []struct {
		Text        string `xml:",chardata"`
		About       string `xml:"about,attr"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Date        time.Time `xml:"date"`
		Language    string `xml:"language"`
		Rights      string `xml:"rights"`
		Source      string `xml:"source"`
		Type        string `xml:"type"`
		Enclosure   struct {
			Text     string `xml:",chardata"`
			Resource string `xml:"resource,attr"`
			Type     string `xml:"type,attr"`
		} `xml:"enclosure"`
		Issued string `xml:"issued"`
	} `xml:"item"`
}

