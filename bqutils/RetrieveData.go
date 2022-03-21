package bqutils

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	projectId = "pokkt-dev-apps-p"
	datasetId = "vdo_pokkt_dev_apps_p"
	tableId   = "rahul_table"
	filename  = "writeEmp.csv"
)

func RetrieveData()  ([][]string) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		fmt.Printf("error %v :", err)
	}

	table := client.DatasetInProject(projectId, datasetId).Table(tableId)
	it := table.Read(ctx)

	rowLimit := 10
	var rowsRead int
	var data [][]string
	for {
		var row []bigquery.Value
		err := it.Next(&row)

		if err == iterator.Done || rowLimit <= rowsRead {
			break
		}
		if err != nil {
			log.Println("Error while iterating row: ", err)
		}
		rowsRead++
		segmentInt := row[1].(int64)
		segmentStr := strconv.Itoa(int(segmentInt))

		record := []string{row[0].(string), segmentStr, row[2].(string)}
		data = append(data, record)
		//fmt.Println(record)
	}

	return data
}
