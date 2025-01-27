package main

import "fmt"

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    reachable := make([][]bool, numCourses)
    for i := range reachable {
        reachable[i] = make([]bool, numCourses)
    }

    for _, p := range prerequisites {
        reachable[p[0]][p[1]] = true
    }

    for k := 0; k < numCourses; k++ {
        for i := 0; i < numCourses; i++ {
            for j := 0; j < numCourses; j++ {
                if reachable[i][k] && reachable[k][j] {
                    reachable[i][j] = true
                }
            }
        }
    }

    result := make([]bool, len(queries))
    for i, q := range queries {
        result[i] = reachable[q[0]][q[1]]
    }

    return result
}

func main() {
    numCourses := 3
    prerequisites := [][]int{{1, 2}, {1, 0}, {2, 0}}
    queries := [][]int{{1, 0}, {1, 2}}

    result := checkIfPrerequisite(numCourses, prerequisites, queries)
    fmt.Println(result) 
}
