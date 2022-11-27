package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var boxNum, boxVolume, a, b, c  int
    if scanner.Scan(){
        s := strings.Split(scanner.Text(), " ")
        boxNum, _ = strconv.Atoi(s[0])
        boxVolume, _= strconv.Atoi(s[1])
    }
    var maxVolume int = 0
    for i := 0; i < boxNum; i++ {
        if scanner.Scan(){
            s := strings.Split(scanner.Text(), " ")
            a, _ = strconv.Atoi(s[0])
            b, _= strconv.Atoi(s[1])
            c, _ = strconv.Atoi(s[2])
        }
        currentVolume := a*b*c
        if currentVolume > maxVolume {
            maxVolume = currentVolume
        }
    }
    fmt.Print(maxVolume-boxVolume)
}