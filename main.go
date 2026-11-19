
package main

import (
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()

    // Konfigurasi CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost:3000",
        },
        AllowMethods: []string{
            "GET",
            "POST",
            "PUT",
            "DELETE",
            "OPTIONS",
        },
        AllowHeaders: []string{
            "Origin",
            "Content-Type",
            "Accept",
            "Authorization",
        },
        MaxAge: 12 * time.Hour,
    }))

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
