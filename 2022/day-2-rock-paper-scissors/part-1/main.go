package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var scoreTable map[string]int = map[string]int{
    "X": 1,
    "Y": 2,
    "Z": 3,
}

var decryptionTable map[string]map[string]int = map[string]map[string]int {
    "X": {
        "A": 3,
        "B": 0,
        "C": 6,
    },
    "Y": { 
        "A": 6,
        "B": 3,
        "C": 0,
    },
    "Z": { 
        "A": 0,
        "B": 6,
        "C": 3,
    },
}

func main() {
    file, err := os.Open("input")

    if err != nil { log.Fatal(err) }

    scanner := bufio.NewScanner(file);

    var scoreTotal int = 0;

    for scanner.Scan() {
        choices := strings.SplitN(scanner.Text(), " ", 2)

        scoreMatch := scoreTable[choices[1]] + decryptionTable[choices[1]][choices[0]]
        scoreTotal += scoreMatch
    }

    fmt.Println("scoreTotal = ", scoreTotal)
}
