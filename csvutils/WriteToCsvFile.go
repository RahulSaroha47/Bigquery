package csvutils

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteToCsv(data [][]string) {

	csvfile, err := os.Create("write.csv")

	if err != nil {
		log.Println(err)
	}

	writeToCsv := csv.NewWriter(csvfile)

	for _, line := range data {
		writeToCsv.Write(line)
	}

	writeToCsv.Flush()
	csvfile.Close()
}
