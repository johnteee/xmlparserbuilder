package main

import (
  "fmt"
  )

func main() {
  parser := AdvertisementParser{}
  builder := AdvertisementBuilder{}

  var err error

  var ads Advertisement
  fmt.Println("\nBy etree library:")
  ads, err = parser.ParseByEtree("ads.xml")
  if (err != nil) {
    panic(err)
  } else {
    fmt.Println(ads)
  }

  fmt.Println("\nBy built-in xml.Unmarshal:")
  ads, err = parser.ParseByUnmarshal("ads.xml")
  if (err != nil) {
    panic(err)
  } else {
    fmt.Println(ads)
  }

  var xmlVersion string
  fmt.Println("\nBy built-in xml.MarshalIndent:")
  xmlVersion, err = builder.buildByMarshalIndent(ads)
  if (err != nil) {
    panic(err)
  } else {
    fmt.Println(xmlVersion)
  }
}
