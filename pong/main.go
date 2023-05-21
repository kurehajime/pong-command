// main.go
package main

import (
    "os"
    "strconv"
)

func main() {
    ipAddr := "*"
    target := 0
    if len(os.Args) >= 2 {
        ipAddr = os.Args[1]
    }
    if len(os.Args) >= 3 {
        if i, err := strconv.Atoi(os.Args[2]); err == nil {
            target = i
        }
    }
    start(ipAddr, target)
}
