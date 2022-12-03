package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var scoreChoice map[string]int = map[string]int{
    "A": 1,
    "B": 2,
    "C": 3,
}

var scoreTableMatch map[string]int = map[string]int{
    "X": 0,
    "Y": 3,
    "Z": 6,
}

var decryptionTable map[string]map[string]string = map[string]map[string]string {
    "A": {
        "X": "C",
        "Y": "A",
        "Z": "B",
    },
    "B": { 
        "X": "A",
        "Y": "B",
        "Z": "C",
    },
    "C": { 
        "X": "B",
        "Y": "C",
        "Z": "A",
    },
}

func main() {
    file, err := os.Open("input")

    if err != nil { log.Fatal(err) }

    scanner := bufio.NewScanner(file);

    var totalScore int = 0;

    for scanner.Scan() {
        choices := strings.SplitN(scanner.Text(), " ", 2)

        opponentChoice, myChoice := choices[0], choices[1]

        decryptedChoice := decryptionTable[opponentChoice][myChoice]

        matchScore := scoreTableMatch[myChoice] + scoreChoice[decryptedChoice]

        totalScore += matchScore
    }

    fmt.Println("score = ", totalScore)
}
