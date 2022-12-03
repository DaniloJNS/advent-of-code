package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)


func main() {
    file, err := os.Open("input")

    if err != nil {
        log.Fatal(err);
    }

    scanner := bufio.NewScanner(file);

    var sum int = 0
    greaters := []int{0, 0, 0};
    var unit int

    for scanner.Scan() {
        input := scanner.Text()

        if input == "" {
            if sum > greaters[0] {
                greaters[0] = sum
                sort.Ints(greaters);
            }

            sum = 0
            continue
        } 

        fmt.Sscanf(input, "%d\n", &unit)

        sum += unit
    }

    sum = 0;
    for greater := range greaters {
        sum += greater
    }

    fmt.Println("greater = ", sum)
}
