package main

import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "net/http"
    //"os"
    "strings"
)

func main(){
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        fmt.Println("Welcome!")
        c.String(http.StatusOK, "Welcome!")
    })
    router.GET("/health", func(c *gin.Context) {
        fmt.Println("OK")
        c.String(http.StatusOK, "200")
    })
    router.POST("/events", GithubHMACValidator(), AMQPEventHandler())
    router.Run(":8080")
}

func GithubHMACValidator() gin.HandlerFunc {
    return func(c *gin.Context) {
        body, _ := ioutil.ReadAll(c.Request.Body)
        values, _ := c.Request.Header["X-Hub-Signature"]

        enc := strings.Replace(values[0], "sha1=", "", 1)
        sig, _ := hex.DecodeString(enc)

        key := []byte("PaSsWoRdSaReGr8")
        digest := hmac.New(sha1.New, key)
        digest.Write(body)
        computed := digest.Sum(nil)

        if ok := hmac.Equal(computed, sig); ok {
            c.Next()
        } else {
            c.String(http.StatusBadRequest, "Error: invalid signature")
        }
    }
}

func AMQPEventHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        var eventType string
        if values, _ := c.Request.Header["X-Github-Event"]; len(values) > 0 {
            eventType = values[0]
        }
        log.Println("Recieved Event:", eventType)
        log.Println("TODO: Implement an AMQP Hander to queue these events")
    }
}
