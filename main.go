
package main

import (
  "project/internal/db/cache"
  "project/internal/db/receipts"
  "github.com/gin-gonic/gin"
  "net/http"
)


func main() {
  router := gin.Default()

  router.POST("/receipts/process", func(c *gin.Context) {
    var receipt receipts.Receipt
    if err := c.ShouldBindJSON(&receipt); err != nil {
        c.JSON(400, gin.H{"Binding error": err.Error()})
    }

    id := cache.Set(receipt)

    cache.CachePoints[id] = cache.CalculatePoints(id)


    c.JSON(http.StatusOK, gin.H{"id": id})
  })

  router.GET("/", func(c *gin.Context) {
    // c.String(200, "Welcome.")
    c.Redirect(302, "/receipts/:id/points")
  })

  router.GET("/receipts/:id/points", func(c *gin.Context) {
    id := c.Param("id")
    
    c.JSON(http.StatusOK, gin.H{"points": cache.CachePoints[id]})
  })

  router.Run("localhost:8080")
}


