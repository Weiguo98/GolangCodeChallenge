package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)


func main() {
    m := make(map[string]int)
    scanner := bufio.NewScanner(os.Stdin)
    count := 0
    for scanner.Scan() {
        wood := scanner.Text()
        count++
        if _, ok := m[wood]; ok {
            m[wood] = m[wood]+1
        } else {
            m[wood] = 1
        }
    }
    
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys{
        fmt.Printf("%v %.6f\n", k, float32(m[k])*100.0 / float32(count))
    }
}