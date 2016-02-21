package main

import (
  "fmt"
  "strings"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/ernesto-jimenez/scraperboard"
  "github.com/leebenson/conform"
)

func Name(r render.Render, params martini.Params) {

  titleUrl := fmt.Sprintf("http://www.imdb.com/name/%s", params["id"])

  scraper, _ := scraperboard.NewScraperFromFile("scraper_files/name.xml")

  var nameresponse NameResponse

  scraper.ExtractFromURL(titleUrl, &nameresponse)

  // IMDB sends back a lot of whitespaces, let's trim it
  conform.Strings(&nameresponse)

  // Get the full size of the poster
  image := nameresponse.Image
  image = strings.Replace(image, "_V1_UY317_CR7,0,214,317_AL_", "_V1_SX640_SY720_", -1)
  nameresponse.Image = image

  r.JSON(200, nameresponse)
}

type NameResponse struct {
  Name        string `conform:"trim"`
  Bio         string `conform:"trim"`
  Image       string
  Born        string
  KnownFor    []knownfor
}

type knownfor struct {
  Title        string `conform:"trim"`
  Year         string `conform:"trim"`
  Url          string
}
