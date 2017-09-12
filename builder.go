package main

import (
  "encoding/xml"
  // "fmt"
  )

type AdvertisementBuilder struct {
}

func (this AdvertisementBuilder) buildByMarshalIndent(ads Advertisement) (string, error) {
  var xmlVersion []byte
  var err error
  xmlVersion, err = xml.MarshalIndent(ads, "", "\t")
  if err != nil {
      return "", err
  }

  return string(xmlVersion), nil
}
