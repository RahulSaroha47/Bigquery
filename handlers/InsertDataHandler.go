package handlers

import (
	"fmt"
	"log"
	"sync"
	"time"

	"bqutils"
	"csvutils"

	"github.com/gin-gonic/gin"
)

func InsertDatatoBqTable(c *gin.Context) {

	c1 := make(chan []string)
	c2 := make(chan [][]string)

	start := time.Now()

	go csvutils.ReadDatafromCsv("writeEmp.csv", c1)

	go func(chInput <-chan []string, chOutput chan<- [][]string, n int) {
		rows := [][]string{}
		counter := 0
		for {
			row, more := <-chInput
			if !more {
				break
			}
			rows = append(rows, row)
			counter++
			if counter == n {
				chOutput <- rows
				rows = [][]string{}
				counter = 0
			}
		}
		if len(rows) != 0 {
			chOutput <- rows
		}
		close(chOutput)
	}(c1, c2, 500)

	var wg sync.WaitGroup
	//
	for idx := 0; idx < 300; idx++ {
		wg.Add(1)
		go bqutils.InsertCsvData(projectId, datasetId, tableId, c2, &wg)
	}

	wg.Wait()

	log.Println("Csv Data Inserted")
	fmt.Printf("done in %v seconds", time.Since(start).Seconds())
}
