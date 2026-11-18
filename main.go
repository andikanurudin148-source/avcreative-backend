package main

import (
    "os"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

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