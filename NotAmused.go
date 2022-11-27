package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
)

type customerTime struct{
    //lastest in time
    inTime int
    overallTime int
}


func main() {
    m := make(map[string]*customerTime)
    scanner := bufio.NewScanner(os.Stdin)
    count := 0
    for scanner.Scan() {
        if scanner.Text() == "CLOSE" {
            count += 1
            fmt.Printf("Day %v\n", count)
            keys := make([]string, 0, len(m))
            for k := range m {
                keys = append(keys, k)
            }
            sort.Strings(keys)
            for _, k := range keys{
                fmt.Printf("%v $%.2f\n", k, float32(m[k].overallTime) * 0.10)
            }
            fmt.Println()
            m = make(map[string]*customerTime)
            continue
        }
        res := strings.Split(scanner.Text()," ")
        // record when people enter and exit, calculate the overall time
        if res[0] == "ENTER"{
            time, _ := strconv.Atoi(res[2])
            if customer, ok := m[res[1]]; ok {
                customer.inTime = time
            } else {
                m[res[1]] = &customerTime{inTime: time}
            }
        } else if res[0] == "EXIT"{
            time, _ := strconv.Atoi(res[2])
            if customer, ok := m[res[1]]; ok {
                customer.overallTime= time - customer.inTime + customer.overallTime
            } else {
                fmt.Println("error")
            }
        }
    }
}