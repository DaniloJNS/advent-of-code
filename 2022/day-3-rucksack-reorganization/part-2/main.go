package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const deltaLowerLetterPriority int = int('a') - 1
const deltaUpperLetterPriority int = int('A') - 27

func main() {
    file, err := os.Open("input")

    if err != nil { log.Fatal(err) }

    scanner := bufio.NewScanner(file);

    var total, count int = 0, 1

    var bagItensGroup []string

    for scanner.Scan() {
        bagItensGroup = append(bagItensGroup, scanner.Text())

        if count % 3 == 0 {
            groupOneItens, groupTwoItens, groupTreeItens := bagItensGroup[0], bagItensGroup[1], bagItensGroup[2]
            sharedItens := "" 

            for _, item := range groupOneItens {
                 if strings.Contains(groupTwoItens, string(item)) && strings.Contains(groupTreeItens, string(item)) && 
                 !strings.Contains(sharedItens, string(item)) {
                     sharedItens += string(item)
                     asciiChar := int(item)

                     if asciiChar > deltaLowerLetterPriority {
                        total += asciiChar - deltaLowerLetterPriority
                    } else {
                        total += asciiChar - deltaUpperLetterPriority
                    }
                }
            }

            bagItensGroup = nil
        }

        count++
    }

    fmt.Println("total = ", total)
}
