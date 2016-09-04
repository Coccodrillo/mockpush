package main

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"response": true})
	})
	r.POST("/apns", func(c *gin.Context) {
		_, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error:", err)
		}
		c.Writer.WriteHeader(200)
		c.Writer.Write([]byte(""))
	})
	r.POST("/gcm", func(c *gin.Context) {
		_, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error:", err)
		}
		c.Writer.WriteHeader(200)
		c.Writer.Write([]byte(""))
	})
	r.Run(":8888")
}
