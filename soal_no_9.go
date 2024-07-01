package main

import (
    "fmt"
    "sort"
    "strings"
)

// a. Mengurutkan dan menghapus duplikat
func sortAndRemoveDuplicates(arr []int) []int {
    uniqueMap := make(map[int]bool)
    var result []int
    for _, num := range arr {
        if !uniqueMap[num] {
            uniqueMap[num] = true
            result = append(result, num)
        }
    }
    sort.Ints(result)
    return result
}

// b. Menampilkan total duplikat
func countDuplicates(arr []int) string {
    countMap := make(map[int]int)
    for _, num := range arr {
        countMap[num]++
    }
    
    var result []string
    for _, num := range sortAndRemoveDuplicates(arr) {
        result = append(result, fmt.Sprintf("%d[%d]", num, countMap[num]))
    }
    return strings.Join(result, ",")
}

// c. Menghapus value sesuai input
func removeValues(arr []int, toRemove []int) []int {
    removeMap := make(map[int]bool)
    for _, num := range toRemove {
        removeMap[num] = true
    }
    
    var result []int
    for _, num := range arr {
        if !removeMap[num] {
            result = append(result, num)
        }
    }
    return result
}

// d. Menjumlahkan sesuai input dengan maksimal 10
func addToMaxTen(arr []int, toAdd int) []int {
    result := make([]int, len(arr))
    copy(result, arr)
    
    for i := range result {
        if toAdd > 0 {
            space := 10 - result[i]
            if space > 0 {
                if toAdd >= space {
                    result[i] = 10
                    toAdd -= space
                } else {
                    result[i] += toAdd
                    toAdd = 0
                }
            }
        } else {
            break
        }
    }
    return result
}

func main() {
    arr := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}

    // a. Mengurutkan dan menghapus duplikat
    fmt.Println("a. Sorted and unique:", sortAndRemoveDuplicates(arr))

    // b. Menampilkan total duplikat
    fmt.Println("b. Duplicates count:", countDuplicates(arr))

    // c. Menghapus value sesuai input
    toRemove := []int{9, 1, 6, 4}
    fmt.Println("c. After removing", toRemove, ":", removeValues(arr, toRemove))

    // d. Menjumlahkan sesuai input dengan maksimal 10
    toAdd := 15
    fmt.Println("d. After adding", toAdd, ":", addToMaxTen(arr, toAdd))
}