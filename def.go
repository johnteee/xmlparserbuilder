package main

type AdvertisementVideo struct {
	Value string
	Duration string
	Type string
	Width string
	Height string
}

type AdvertisementTracking struct {
	Value string
}

type Advertisement struct {
	Title string
	Description string
	Icon string
	Videos []AdvertisementVideo
	Trackings []AdvertisementTracking
}
