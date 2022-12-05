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

    var total int = 0;

    for scanner.Scan() {
        allItens := scanner.Text()
  
        bagOneItens, bagTwoItens := allItens[:len(allItens) / 2], allItens[len(allItens) / 2:]
        sharedItens := "" 
  
        for _, item := range bagOneItens { 
             if strings.Contains(bagTwoItens, string(item)) && !strings.Contains(sharedItens, string(item)) {
                 sharedItens += string(item)
                 asciiChar := int(item)

                 if asciiChar > deltaLowerLetterPriority {
                    total += asciiChar - deltaLowerLetterPriority
                } else {
                    total += asciiChar - deltaUpperLetterPriority
                }
            }
        }
    }

    fmt.Println("total = ", total)
}
