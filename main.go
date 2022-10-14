package main

import (
  	"fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type msg struct {
    Message  string  `json:"message"`
}

var messages = msg{Message: "API Testing"}

func main() {
    router := gin.Default()
    router.GET("/", getMessages)
    router.GET("/v1/messages", getMessages)
    router.Run("localhost:8080")
  
    fmt.Println("Test console..")
}

func getMessages(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, messages)
}
