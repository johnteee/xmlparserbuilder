package main

import (
  "fmt"
  )

func main() {
  parser := AdvertisementParser{}
  ads := parser.Parse("ads.xml")

  fmt.Println(ads)
}
