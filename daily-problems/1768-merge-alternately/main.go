package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
    i, j := 0, 0
    merged := make([]byte, 0, len(word1)+len(word2))

    for i < len(word1) || j < len(word2) {
        if i < len(word1) {
            merged = append(merged, word1[i])
            i++
        }
        if j < len(word2) {
            merged = append(merged, word2[j])
            j++
        }
    }

    return string(merged)
}

func main() {
    word1 := "abc"
    word2 := "pqr"
    fmt.Println(mergeAlternately(word1, word2)) 
}
