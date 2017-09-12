package main

type AdvertisementVideo struct {
	Value string `xml:",chardata"`
	Duration string `xml:"duration,attr"`
	Type string `xml:"type,attr"`
	Width string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}

type AdvertisementTracking struct {
	Value string `xml:",chardata"`
	Event string `xml:"event,attr"`
	Offset string `xml:"offset,attr"`
}

type Advertisement struct {
	Title string `xml:"Title"`
	Description string `xml:"Description"`
	Icon string `xml:"Icon"`
	Videos []AdvertisementVideo `xml:"Videos>Video"`
	Trackings []AdvertisementTracking `xml:"Trackings>Tracking"`
}
