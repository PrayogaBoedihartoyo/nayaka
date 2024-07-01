package main

import (
    "fmt"
    "strings"
)

func checkCity(city string, cities []string) (bool, string) {
    city = strings.ToLower(city)
    for _, c := range cities {
        if strings.ToLower(c) == city {
            return true, ""
        }
    }

    var suggestions []string
    firstLetter := string(city[0])
    lastLetter := string(city[len(city)-1])

    for _, c := range cities {
        c = strings.ToLower(c)
        if string(c[0]) == firstLetter {
            suggestions = append(suggestions, strings.Title(c))
        }
        if string(c[len(c)-1]) == lastLetter && !contains(suggestions, strings.Title(c)) {
            suggestions = append(suggestions, strings.Title(c))
        }
    }

    if len(suggestions) > 0 {
        return false, "Saran Kota = " + strings.Join(suggestions, " , ")
    }

    return false, ""
}

func contains(slice []string, item string) bool {
    for _, a := range slice {
        if a == item {
            return true
        }
    }
    return false
}

func main() {
    cities := []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}
    
    testCases := []string{"Bandung", "Bogor", "Jakarta", "Ambon"}
    
    for _, city := range testCases {
        found, suggestion := checkCity(city, cities)
        if found {
            fmt.Printf("%s: true\n", city)
        } else {
            if suggestion != "" {
                fmt.Printf("%s: false, %s\n", city, suggestion)
            } else {
                fmt.Printf("%s: false\n", city)
            }
        }
    }
}