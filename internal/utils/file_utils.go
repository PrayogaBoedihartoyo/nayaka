package utils

import (
    "encoding/json"
    "fmt"
    "os"
)

func SaveJSONToFile(data interface{}, filename string) error {
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return fmt.Errorf("error marshaling JSON: %v", err)
    }

    err = os.WriteFile(filename, jsonData, 0644)
    if err != nil {
        return fmt.Errorf("error writing to file: %v", err)
    }

    return nil
}

func ReadJSONFromFile(filename string) (map[string]interface{}, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var result map[string]interface{}
    err = json.Unmarshal(data, &result)
    if err != nil {
        return nil, err
    }

    return result, nil
}