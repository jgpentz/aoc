package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func read_cards(fname string) (int, []string) {
    var cards []string

    // Open the file and close it when this function finishes
    file, err := os.Open(fname)
    if err != nil {
        fmt.Println("Error occured when opening file")
        return 1, cards
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        cards = append(cards, scanner.Text())
    }

    return 0, cards
}

func get_card_score(card string) int {
    sum := 0

    // Split the string into 3 parts: Card #, winning numbers, our numbers
    card = strings.Split(card, ": ")[1]
	winning_nums := strings.Fields(strings.Split(card, " | ")[0])
	our_nums := strings.Fields(strings.Split(card, " | ")[1])

    // Store the winning numbers in a hash set
    winning_set := make(map[string]bool)
    for _, number := range winning_nums {
        winning_set[number] = true
    }
    

    // Check our numbers for existence in hash set, doubling score each time a
    // number is found
    for _, number := range our_nums {
        if winning_set[number] {
            if sum == 0 {
                sum = 1
            } else {
                sum = sum * 2
            }
        }
    }

    return sum
}

func main() {
    // ----- Verify that the test case works -----
    // Read in all cards
    _, cards := read_cards("input_test.txt")

    // Sum the score of each card
    sum := 0
    for _, card := range cards {
        sum += get_card_score(card)
    }

    if sum != 13 {
        panic(fmt.Sprintf("Test case failed! Expected: 13, computed: %v", sum))
    }

    // ----- Run the real input -----
    // Read in all cards
    _, cards = read_cards("input.txt")

    // Sum the score of each card
    sum = 0
    for _, card := range cards {
        sum += get_card_score(card)
    }

    fmt.Println(sum)
}
