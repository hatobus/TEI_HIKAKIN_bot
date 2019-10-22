package model

import "encoding/xml"

type Subsc struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Yt      string   `xml:"yt,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Link    []struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Entry   struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		VideoId   string `xml:"videoId"`
		ChannelId string `xml:"channelId"`
		Title     string `xml:"title"`
		Link      struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Author struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
			URI  string `xml:"uri"`
		} `xml:"author"`
		Published string `xml:"published"`
		Updated   string `xml:"updated"`
	} `xml:"entry"`
}
