package main

import (
  "fmt"
  )

func main() {
  parser := AdvertisementParser{}

  var ads Advertisement
  fmt.Println("\nBy etree library:")
  ads = parser.Parse("ads.xml")
  fmt.Println(ads)

  fmt.Println("\nBy built-in xml.Unmarshal:")
  ads = parser.ParseByUnmarshal("ads.xml")
  fmt.Println(ads)
}
