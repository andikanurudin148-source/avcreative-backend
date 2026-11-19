
package main

import (
    "os"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Konfigurasi CORS
    router.Use(func(c *gin.Context) {
        origin := c.GetHeader("Origin")

        if origin == "http://localhost:3000" {
            c.Header("Access-Control-Allow-Origin", origin)
            c.Header("Vary", "Origin")
        }

        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

    // Endpoint API
    router.GET("/api/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Backend AV Creative berhasil terhubung!",
        })
    })

    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    router.Run(":" + port)
}
