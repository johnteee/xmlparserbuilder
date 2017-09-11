package main

type AdvertisementVideo struct {
	Value string
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
