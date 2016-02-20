package main

import (
  "github.com/gin-gonic/gin"
)

/**
 * The urls we're gonna use
 */
var urls = []string{"/top", "/toptv", "/title/:id"}

func index(c *gin.Context) {
  content := gin.H{"urls": urls}
  c.JSON(200, content)
}

func main() {
  app := gin.Default()
  app.GET("/", index)
  app.GET(urls[0], Top)
  app.GET(urls[1], TopTV)
  app.Run(":1990") // Let's go with the year IMDB was created
}
