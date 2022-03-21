package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadDatafromCsv(filename string, chanStr chan<- []string) {
	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Csv file opened succesfully")
	//scanner := bufio.NewScanner(csvFile)
	csvData, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("error occured %v", err)
	}
	defer csvFile.Close()

	for _, val := range csvData {
		text := val
		chanStr <- text
	}
	close(chanStr)

}
