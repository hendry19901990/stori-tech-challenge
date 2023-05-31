package repository

import (
	"encoding/csv"
    "fmt"
	"os"
)



func ReadCsvFile(filePath string) (*[][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
		return nil, fmt.Errorf("Unable to read input file " + filePath, err)
        
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
		return nil, fmt.Errorf("Unable to parse file as CSV for " + filePath, err)
    }

    return &records, nil
}
 