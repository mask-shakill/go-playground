package main

import "fmt"

type Solution struct{}

func (s *Solution) findTheDifference(s1, t1 string) byte {
    var sumT, sumS int
    for i := 0; i < len(t1); i++ {
        sumT += int(t1[i])
    }
    for i := 0; i < len(s1); i++ {
        sumS += int(s1[i])
    }
    return byte(sumT - sumS)
}

func main() {
    solution := &Solution{}
    s1, t1 := "abcd", "abcde"
    s2, t2 := "", "y"
    fmt.Println(string(solution.findTheDifference(s1, t1))) 
    fmt.Println(string(solution.findTheDifference(s2, t2))) 
}
