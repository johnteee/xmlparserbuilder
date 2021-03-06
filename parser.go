package main

import (
  "github.com/beevik/etree"
  "io/ioutil"
  "encoding/xml"
  // "fmt"
  )

type AdvertisementParser struct {
}

func (this AdvertisementParser) ParseByUnmarshal(filename string) (Advertisement, error) {
  var data []byte
  var err error
  data, err = ioutil.ReadFile(filename)
  if err != nil {
      return Advertisement{}, err
  }

  var ads Advertisement;
  err = xml.Unmarshal(data, &ads)
  return ads, nil
}

func (this AdvertisementParser) ParseByEtree(filename string) (Advertisement, error) {
  doc := etree.NewDocument()
  if err := doc.ReadFromFile(filename); err != nil {
      return Advertisement{}, err
  }

  ads := Advertisement{}
  root := doc.SelectElement("Advertisement")
  var element *etree.Element;

  element = root.SelectElement("Title")
  if (element != nil) {
    ads.Title = element.Text()
  }
  element = root.SelectElement("Description")
  if (element != nil) {
    ads.Description = element.Text()
  }
  element = root.SelectElement("Icon")
  if (element != nil) {
    ads.Icon = element.Text()
  }

  element = root.SelectElement("Videos")
  if (element != nil) {
    ads.Videos = []AdvertisementVideo{}
    for  _, item := range element.ChildElements() {
      // println(fmt.Sprintf("%d", i))

      var itemObject AdvertisementVideo
      itemObject.Value = item.Text()
      itemObject.Duration = item.SelectAttrValue("duration", "")
      itemObject.Type = item.SelectAttrValue("type", "")
      itemObject.Width = item.SelectAttrValue("width", "")
      itemObject.Height = item.SelectAttrValue("height", "")

      ads.Videos = append(ads.Videos, itemObject)
    }
  }

  element = root.SelectElement("Trackings")
  if (element != nil) {
    ads.Trackings = []AdvertisementTracking{}
    for  _, item := range element.ChildElements() {
      // println(fmt.Sprintf("%d", i))

      var itemObject AdvertisementTracking
      itemObject.Value = item.Text()
      itemObject.Event = item.SelectAttrValue("event", "")
      itemObject.Offset = item.SelectAttrValue("offset", "")

      ads.Trackings = append(ads.Trackings, itemObject)
    }
  }

  return ads, nil
}
