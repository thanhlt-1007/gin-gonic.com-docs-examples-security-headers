package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func getPingHandler(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "pong",
        },
    )
}

func setHeaderMiddleware(context *gin.Context) {
    context.Header("X-Frame-Options", "DENY")
    context.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
    context.Header("X-XSS-Protection", "1; mode=block")
    context.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
    context.Header("Referrer-Policy", "strict-origin")
    context.Header("X-Content-Type-Options", "nosniff")
    context.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
    context.Next()
}

func main() {
    engine := gin.Default()
    engine.Use(setHeaderMiddleware)
    engine.GET("/ping", getPingHandler)
    engine.Run()
}
