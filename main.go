package main

import (
  "github.com/gin-gonic/gin"
)

/**
 * The urls we're gonna use
 */
var urls = []string{"/top", "/toptv", "/title/:id"}
var DB = make(map[string]string)

func Index(c *gin.Context) {
  content := gin.H{"urls": urls}
  c.JSON(200, content)
}

func PageNotFound(c *gin.Context) {
  c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
}

func main() {
  app := gin.Default()

  v1 := app.Group("api/v1")
  {
    v1.GET("/", Index)
    v1.GET(urls[0], Top)
    v1.GET(urls[1], TopTV)
  }

  // Handle No Routes
  app.NoRoute(PageNotFound)

  // Let's go with the year IMDB was created
  app.Run(":1990")
}
