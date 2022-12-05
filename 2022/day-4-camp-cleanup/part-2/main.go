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
	fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &pairX0, &pairX1, &pairY0, &pairY1)

	deltaPairX := (pairY1 - pairX0)
	deltaPairY := (pairX1 - pairY0)

	if deltaPairX >= 0 && deltaPairY >= 0{
	    pairIncluded++
	} else if deltaPairY >= 0 && deltaPairX >= 0{
	    pairIncluded++
	}
    }

    fmt.Println("total =", pairIncluded)
}
