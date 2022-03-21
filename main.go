package main

import (
	"github.com/gin-gonic/gin"
	"handlers"
)

func main() {
	router := gin.Default()
	router.GET("/bigquery/insert",handlers.InsertDatatoBqTable)
	router.GET("/bigquery/retrieve", handlers.RetrieveDataFromBqTable)
	router.Run() 
}
