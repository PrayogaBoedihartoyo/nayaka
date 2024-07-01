package main

import (
    "fmt"
    "math/rand"
    "sort"
    "strings"
    "time"
)

// Fungsi untuk generate string random
func generateRandomString() string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    const digits = "0123456789"
    
    rand.Seed(time.Now().UnixNano())
    
    result := make([]byte, 100)
    for i := 0; i < 50; i++ {
        result[i] = letters[rand.Intn(len(letters))]
        result[i+50] = digits[rand.Intn(len(digits))]
    }
    
    rand.Shuffle(len(result), func(i, j int) {
        result[i], result[j] = result[j], result[i]
    })
    
    return string(result)
}

// a. Fungsi untuk menghitung statistik
func countStatistics(s string) (int, int, int, int) {
    totalLetters := 0
    totalVowels := 0
    totalDigits := 0
    totalEvenDigits := 0
    vowels := "aeiouAEIOU"
    
    for _, char := range s {
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            totalLetters++
            if strings.ContainsRune(vowels, char) {
                totalVowels++
            }
        } else if char >= '0' && char <= '9' {
            totalDigits++
            if (char-'0')%2 == 0 {
                totalEvenDigits++
            }
        }
    }
    
    return totalLetters, totalVowels, totalDigits, totalEvenDigits
}

// b. Fungsi untuk mengurutkan dan menghapus duplikat
func sortAndRemoveDuplicates(s string) string {
    letters := make(map[rune]bool)
    digits := make(map[rune]bool)
    
    for _, char := range s {
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            letters[char] = true
        } else if char >= '0' && char <= '9' {
            digits[char] = true
        }
    }
    
    sortedLetters := []rune{}
    for letter := range letters {
        sortedLetters = append(sortedLetters, letter)
    }
    sort.Slice(sortedLetters, func(i, j int) bool {
        return sortedLetters[i] < sortedLetters[j]
    })
    
    sortedDigits := []rune{}
    for digit := range digits {
        sortedDigits = append(sortedDigits, digit)
    }
    sort.Slice(sortedDigits, func(i, j int) bool {
        return sortedDigits[i] > sortedDigits[j]
    })
    
    return string(sortedDigits) + string(sortedLetters)
}

// c. Fungsi untuk mengurutkan dengan kombinasi
func sortCombined(s string) string {
    type charInfo struct {
        char rune
        isDigit bool
    }
    
    var chars []charInfo
    for _, char := range s {
        isDigit := char >= '0' && char <= '9'
        chars = append(chars, charInfo{char, isDigit})
    }
    
    sort.Slice(chars, func(i, j int) bool {
        if chars[i].isDigit != chars[j].isDigit {
            return chars[i].isDigit
        }
        if chars[i].isDigit {
            return chars[i].char > chars[j].char
        }
        return chars[i].char < chars[j].char
    })
    
    result := ""
    for _, ci := range chars {
        result += string(ci.char)
    }
    
    return result
}

func main() {
    randomString := generateRandomString()
    fmt.Println("Random string:", randomString)
    
    // a. Menghitung statistik
    totalLetters, totalVowels, totalDigits, totalEvenDigits := countStatistics(randomString)
    fmt.Printf("Total huruf: %d\n", totalLetters)
    fmt.Printf("Total huruf vokal: %d\n", totalVowels)
    fmt.Printf("Total angka: %d\n", totalDigits)
    fmt.Printf("Total angka genap: %d\n", totalEvenDigits)
    
    // b. Mengurutkan dan menghapus duplikat
    sortedUnique := sortAndRemoveDuplicates(randomString)
    fmt.Println("Sorted and unique:", sortedUnique)
    
    // c. Mengurutkan kombinasi
    sortedCombined := sortCombined(randomString)
    fmt.Println("Sorted combined:", sortedCombined)
}