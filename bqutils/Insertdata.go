package bqutils

import (
	"context"
	"fmt"
	"sync"
	"cloud.google.com/go/bigquery"
)

type Person struct {
	name string
	age  string
	city string
}

func (i *Person) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"name": i.name,
		"age":  i.age,
		"city": i.city,
	}, bigquery.NoDedupeID, nil
}

func InsertCsvData(projectID, datasetID, tableID string, ch <-chan [][]string, wg *sync.WaitGroup) {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID)

	if err != nil {
		fmt.Printf("bigquery.NewClient: %v", err)
	}

	defer client.Close()
	for {
		strCsv, more := <-ch
		if !more {
			break
		} else {
			// Convert to array
			person := []*Person{}
			for _, csv := range strCsv {

				person = append(person, &Person{csv[0], csv[1], csv[2]})
			}

			inserter := client.Dataset(datasetID).Table(tableID).Inserter()

			if err := inserter.Put(ctx, person); err != nil {
				fmt.Printf("error occured %v", err)
			}

		}
	}
	client.Close()
	wg.Done()

}
