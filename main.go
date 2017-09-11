package main

import (
  )

func main() {
  parser := AdvertisementParser{}
  ads := parser.Parse("ads.xml")

  println(ads.Title)
}
