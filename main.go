
package main

import (
  // "bytes"
  "project/internal/db/cache"
  "project/internal/db/receipts"
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
	// "reflect"
	// "encoding/json"
	// "string"
	// "errors"
	// "github.com/google/uuid"
  // "io"
  // "os"
  // "encoding/json"
)
    // body, _ := io.ReadAll(c.Request.Body)
    // fmt.Println("Raw body:", string(body))
    // // Reset body so ShouldBindJSON still works
    // c.Request.Body = io.NopCloser(bytes.NewBuffer(body))


func main() {
  router := gin.Default()

  router.POST("/receipts/process", func(c *gin.Context) {

    var receipt receipts.Receipt
    if err := c.ShouldBindJSON(&receipt); err != nil {
        c.JSON(400, gin.H{"Binding error": err.Error()})
    }

    id := cache.Set(receipt)
    fmt.Printf("Parsed Receipt: %+v\n", cache.CacheMap[id])

    cache.CachePoints[id] = cache.CalculatePoints(id)


    c.JSON(http.StatusOK, gin.H{"id": id})
  })

  router.GET("/receipts/:id/points", func(c *gin.Context) {
    id := c.Param("id")


    c.JSON(http.StatusOK, gin.H{"points": cache.CachePoints[id]})
  })

  router.Run("localhost:8080")
}

    // content := gin.H{"points": cache.CachePoints.id}

    // c.JSON(http.StatusOK, content)


