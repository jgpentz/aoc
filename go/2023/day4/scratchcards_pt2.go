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
            sum += 1
        }
    }

    return sum
}

func process_cards(cards []string) int {
    var queue []int
    num_cards := 0

    // Iniitally store all cards in queue
    for icard, _ := range cards {
        queue = append(queue, icard)
    }

    for {
        if len(queue) < 1 {
            break
        }

        card_idx := queue[0]
        queue = queue[1:]
        num_cards += 1

        card_score := get_card_score(cards[card_idx])

        for i := 1; i < card_score + 1; i++ {
            if card_idx + i > len(cards) - 1 {
                break
            }

            queue = append(queue, card_idx + i)
        }
    }

    return num_cards
}

func main() {
    // ----- Verify that the test case works -----
    // Read in all cards
    _, cards := read_cards("input_test.txt")

    // Sum the score of each card
    sum := process_cards(cards)

    if sum != 30 {
        panic(fmt.Sprintf("Test case failed! Expected: 30, computed: %v", sum))
    }

    // ----- Run the real input -----
    // Read in all cards
    _, cards = read_cards("input.txt")

    // Sum the score of each card
    sum = process_cards(cards)

    fmt.Println(sum)
}
