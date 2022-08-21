package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    a := []int{}
    f, err := os.Open("reversed.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        lineStr := scanner.Text()
        num, _ := strconv.Atoi(lineStr)

        a = append(a, num)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("sorted: ", mergeSort(a))

}

func mergeSort(items []int) []int {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[0:len(items)/2])
    second := mergeSort(items[len(items)/2:len(items)])
    return merge(first, second)
}

func merge(left []int, right []int) []int {
    merged := []int{}
    leftIndex := 0
    rightIndex := 0
    for leftIndex < len(left) && rightIndex < len(right) {
        if left[leftIndex] < right[rightIndex] {
            merged = append(merged, left[leftIndex])
            leftIndex++
        } else {
            merged = append(merged, right[rightIndex])
            rightIndex++
        }
    }
    for ; leftIndex < len(left); leftIndex++ {
        merged = append(merged, left[leftIndex])
    }
    for ; rightIndex < len(right); rightIndex++ {
        merged = append(merged, right[rightIndex])
    }
    return merged
}
