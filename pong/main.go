// main.go
package main

import "os"

func main() {
	ipAddr := "*"
	if len(os.Args) >= 2 {
		ipAddr = os.Args[1]
	}
	start(ipAddr)
}
