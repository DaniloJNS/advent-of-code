package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    file, err := os.Open("input")

    if err != nil { log.Fatal(err) }

    scanner := bufio.NewScanner(file);

    var pairX0, pairX1, pairY0, pairY1 int 

    var pairIncluded int

    for scanner.Scan() {
	fmt.Sscanf(scanner.Text(),"%d-%d,%d-%d", &pairX0, &pairX1, &pairY0, &pairY1)

	deltaPairX := (pairX0 - pairY0)
	deltaPairY := (pairX1 - pairY1)

	if deltaPairX >= 0 && deltaPairY <= 0 {
	    pairIncluded++
	} else if deltaPairX <= 0 && deltaPairY >= 0 {
	    pairIncluded++
	}
    }

    fmt.Println("total =", pairIncluded)
}
