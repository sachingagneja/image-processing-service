package store_master

import (
    "encoding/csv"
    "log"
    "os"
)

type Store struct {
    StoreID   string
    StoreName string
    AreaCode  string
}

var StoreData map[string]Store

func LoadStoreData(filepath string) error {
    file, err := os.Open(filepath)
    if err != nil {
        log.Printf("Failed to open file %s: %v", filepath, err)
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Printf("Failed to read CSV data: %v", err)
        return err
    }

    StoreData = make(map[string]Store)
    for i, record := range records {
        // Skip the header row
        if i == 0 {
            continue
        }

        if len(record) < 3 {
            log.Printf("Skipping invalid row %d: %v", i, record)
            continue // Skip invalid rows
        }

        store := Store{
            StoreID:   record[2], // StoreID is now the third column
            StoreName: record[1], // StoreName is now the second column
            AreaCode:  record[0], // AreaCode is now the first column
        }
        StoreData[store.StoreID] = store
    }

    log.Printf("Successfully loaded %d stores", len(StoreData))
    return nil
}
